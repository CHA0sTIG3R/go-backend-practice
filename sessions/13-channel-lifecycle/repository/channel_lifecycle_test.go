package repository

import (
	"fmt"
	"sync"
	"testing"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/10-interfaces/repository"
)

func TestConcurrentAddsUntilChannelClosed(t *testing.T) {
	repo := repository.NewMemoryTaskRepository()

	workers := 10

	var wg sync.WaitGroup

	resultch := make(chan error)

	for i := range workers {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			task := task.Task{
				Name:     "Task " + fmt.Sprintf("%d", i),
				Priority: i % 3,
			}
			resultch <- repo.AddTask(task)
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultch)
	}()

	i := 1
	for err := range resultch {
		err = <-resultch
		if err != nil {
			t.Errorf("go routine %d: Errors from chennels: %v \n", i, err)
		}
		t.Logf("go routine %d: No Errors from chennels: %v \n", i, err)
		i++
	}

	addedTasks, err := repo.GetTasks()
	if err != nil {
		t.Fatalf("Error retrieving tasks: %v", err)
	}

	if len(addedTasks) != workers {
		t.Errorf("Expected %d tasks, but got %d instead", workers, len(addedTasks))
	} else {
		t.Logf("Success! Expected %d tasks, and got %d", workers, len(addedTasks))
	}

}
