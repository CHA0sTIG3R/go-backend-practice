package repository

import (
	"fmt"
	"sync"
	"testing"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/10-interfaces/repository"
)

func TestMemoryRepositoryConcurrentAdds(t *testing.T) {
	repo := repository.NewMemoryTaskRepository()

	var wg sync.WaitGroup

	// Launch roughly 100 goroutines Each should create a different task
	workers := 100

	for i := range workers {
		wg.Add(1)
		fmt.Printf("started %d unique AddTask operations concurrently", workers)
		go func(i int) {
			task := task.Task{
				Name:     "same task", // All tasks have the same name
				Priority: i % 3,
			}
			err := repo.AddTask(task)
			// Check for error, but we expect some of them to fail due to duplicate names
			// We can log the error but not fail the test because we expect duplicates
			if err != nil {
				t.Logf("Error adding task: %v", err)
			}
			
			wg.Done()
		}(i)
	}

	wg.Wait()

	expectedTasks := 1 // Since all tasks have the same name, only one should be added
	retrievedTasks, err := repo.GetTasks()
	if err != nil {
		t.Errorf("Error retrieving tasks: %v", err)
	}
	if len(retrievedTasks) != expectedTasks {
		t.Errorf("Expected %d tasks, but got %d", expectedTasks, len(retrievedTasks))
	}

	for i := range workers {
		wg.Add(1)
		go func(i int) {
			task := task.Task{
				Name:     "diff tasks" + fmt.Sprintf("%d", i), // Ensure unique names for each task
				Priority: i % 3,
			}
			err := repo.AddTask(task)
			if err != nil {
				t.Errorf("Error adding task: %v", err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	retrievedTasks, err = repo.GetTasks()
	if err != nil {
		t.Errorf("Error retrieving tasks: %v", err)
	}
	if len(retrievedTasks) != expectedTasks+workers {
		t.Errorf("Expected %d tasks, but got %d", expectedTasks+workers, len(retrievedTasks))
	}


	for i := range workers {
		wg.Add(1)
		go func(i int) {
			retrievedTasks, err := repo.GetTasks()
			if err != nil {
				t.Errorf("Error retrieving tasks: %v", err)
			}

			// concurrent read should not affect the number of tasks
			if len(retrievedTasks) != expectedTasks+workers {
				t.Errorf("Expected %d tasks, but got %d", expectedTasks+workers, len(retrievedTasks))
			}

			wg.Done()
		}(i)
	}
	wg.Wait()

}
