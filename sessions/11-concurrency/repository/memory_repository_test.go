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

	// Launch roughly 1000 goroutines Each should create a different task
	workers := 1000

	for i := range workers {
		wg.Add(1)
		fmt.Printf("started %d unique AddTask operations concurrently", workers)
		go func(i int) {
			task := task.Task{
				Name:     "Task " + fmt.Sprintf("%d", i), // All tasks have unique names
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

	retrievedTasks, err := repo.GetTasks()
	if err != nil {
		t.Errorf("Error retrieving tasks: %v", err)
	}
	if len(retrievedTasks) != workers {
		t.Errorf("Expected %d tasks, but got %d", workers, len(retrievedTasks))
	}


	for i := range workers {
		wg.Add(1)
		go func(i int) {
			retrievedTasks, err := repo.GetTasks()
			if err != nil {
				t.Errorf("Error retrieving tasks: %v", err)
			}

			// concurrent read should not affect the number of tasks
			if len(retrievedTasks) != workers {
				t.Errorf("Expected %d tasks, but got %d", workers, len(retrievedTasks))
			}

			wg.Done()
		}(i)
	}
	wg.Wait()

}
