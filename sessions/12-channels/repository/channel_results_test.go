package repository

import (
	"fmt"
	"testing"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/10-interfaces/repository"
)

func TestConcurrentAddsWithResultChannel(t *testing.T) {
	repo := repository.NewMemoryTaskRepository()

	workers := 10

	resultch := make(chan error)

	for i := range workers {

		go func(i int) {
			task := task.Task{
				Name:     "Task " + fmt.Sprintf("%d", i),
				Priority: i % 3,
			}
			resultch <- repo.AddTask(task)
		}(i)
	}

	for i := range workers {
		err := <-resultch
		fmt.Printf("go routine %d: Errors from chennels: %v \n", i, err)
	}

	addedTasks, err := repo.GetTasks()
	if err != nil {
		t.Fatalf("Error retrieving tasks: %v", err)
	}
	fmt.Printf("Expected %d tasks, got %d instead", workers, len(addedTasks))
}
