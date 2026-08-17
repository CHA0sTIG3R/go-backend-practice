package repository

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/CHA0sTIG3R/go-backend-practice/sessions/04-json-tasks/task"
	"github.com/CHA0sTIG3R/go-backend-practice/sessions/10-interfaces/repository"
)

func testUniqueTasks(t *testing.T, repo repository.TaskRepository, workers int, wg *sync.WaitGroup) {
	fmt.Printf("started %d unique AddTask operations concurrently\n", workers)
	for i := range workers {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			task := task.Task{
				Name:     "Task " + fmt.Sprintf("%d", i),
				Priority: i % 3,
			}
			err := repo.AddTask(task)
			if err != nil {
				t.Errorf("Error adding task: %v", err)
			}
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
}

func testConcurrentReads(t *testing.T, repo repository.TaskRepository, workers int, wg *sync.WaitGroup) {
	for i := range workers {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			retrievedTasks, err := repo.GetTasks()
			if err != nil {
				t.Errorf("Error retrieving tasks: %v", err)
			}

			if len(retrievedTasks) != workers {
				t.Errorf("Expected %d tasks, but got %d", workers, len(retrievedTasks))
			}
		}(i)
	}
	wg.Wait()
}

func testDuplicateTasks(t *testing.T, repo repository.TaskRepository, workers int, wg *sync.WaitGroup) {
	fmt.Printf("started %d duplicate AddTask operations concurrently\n", workers)
	var start sync.WaitGroup
	start.Add(1)
	errs := make([]error, workers)

	for i := range workers {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			start.Wait()
			errs[i] = repo.AddTask(task.Task{
				Name:     "Same Task",
				Priority: i % 3,
			})
		}(i)
	}
	start.Done()
	wg.Wait()

	validateDuplicateErrors(t, errs, workers)
	validateFinalTaskCount(t, repo)
}

func validateDuplicateErrors(t *testing.T, errs []error, workers int) {
	var success, duplicate int
	for i, err := range errs {
		switch {
		case err == nil:
			success++
		case errors.As(err, new(*repository.DuplicateTaskError)):
			duplicate++
		default:
			t.Errorf("worker %d: unexpected error type %T: %v", i, err, err)
		}
	}

	if success != 1 {
		t.Errorf("expected exactly 1 success, got %d", success)
	}
	if duplicate != workers-1 {
		t.Errorf("expected %d duplicate errors, got %d", workers-1, duplicate)
	}
}

func validateFinalTaskCount(t *testing.T, repo repository.TaskRepository) {
	got, _ := repo.GetTasks()
	if len(got) != 1 {
		t.Errorf("expected 1 stored task, got %d", len(got))
	}
}

func startWriters(wg *sync.WaitGroup, repo repository.TaskRepository, mu *sync.Mutex, writeCount *int, numWriters, tasksPerWriter int) {
	for w := range numWriters {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			for t := range tasksPerWriter {
				task := task.Task{
					Name:     fmt.Sprintf("Task-W%d-T%d", w, t),
					Priority: (w + t) % 3,
				}
				err := repo.AddTask(task)
				if err == nil {
					mu.Lock()
					*writeCount++
					mu.Unlock()
				}
			}
		}(w)
	}
}

func startReaders(t *testing.T, wg *sync.WaitGroup, repo repository.TaskRepository, mu *sync.Mutex, readCount *int, numReaders int) {
	for r := range numReaders {
		wg.Add(1)
		go func(r int) {
			defer wg.Done()
			for range 100 {
				tasks, err := repo.GetTasks()
				if err != nil {
					t.Errorf("reader %d: error reading tasks: %v", r, err)
					continue
				}
				mu.Lock()
				*readCount++
				mu.Unlock()
				_ = tasks
			}
		}(r)
	}
}

func validateConcurrentResults(t *testing.T, repo repository.TaskRepository, writeCount, readCount, numWriters, tasksPerWriter int) {
	finalTasks, err := repo.GetTasks()
	if err != nil {
		t.Fatalf("Error retrieving final tasks: %v", err)
	}

	expectedTaskCount := numWriters * tasksPerWriter
	if len(finalTasks) != expectedTaskCount {
		t.Errorf("Expected %d tasks in repository, got %d", expectedTaskCount, len(finalTasks))
	}

	if writeCount != expectedTaskCount {
		t.Errorf("Expected %d successful writes, got %d", expectedTaskCount, writeCount)
	}

	if readCount == 0 {
		t.Errorf("Expected readers to perform reads, got %d", readCount)
	}

	fmt.Printf("Completed concurrent test: %d writes, %d reads\n", writeCount, readCount)
}

func TestMemoryRepositoryConcurrentAddsDuringWrites(t *testing.T) {
	repo := repository.NewMemoryTaskRepository()
	var wg sync.WaitGroup
	var mu sync.Mutex
	var writeCount, readCount int

	numWriters := 50
	numReaders := 50
	tasksPerWriter := 20

	startWriters(&wg, repo, &mu, &writeCount, numWriters, tasksPerWriter)
	startReaders(t, &wg, repo, &mu, &readCount, numReaders)

	wg.Wait()

	validateConcurrentResults(t, repo, writeCount, readCount, numWriters, tasksPerWriter)
}
