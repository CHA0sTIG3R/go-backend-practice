package workerpool

import (
	"context"
	"sync"
	"testing"
)

type Result struct {
	Input  int
	Output int
}

func worker(
	ctx context.Context,
	jobs <-chan int,
	results chan<- Result,
	started chan<- int,
	release <-chan struct{},
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	started <- 0

	select {
	case <-ctx.Done():
		return
	case <-release:
		for job := range jobs {
			res := Result{
				Input:  job,
				Output: job * job,
			}

			results <- res
		}
	}
}

func TestWorkerStopsWhenContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	const workers = 2
	const jobsCount = 4

	jobs := make(chan int, jobsCount)
	results := make(chan Result, jobsCount)
	started := make(chan int, workers)
	release := make(chan struct{})

	for i := range jobsCount {
		jobs <- i
	}

	close(jobs)

	for range workers {
		wg.Add(1)
		go worker(
			ctx,
			jobs,
			results,
			started,
			release,
			&wg,
		)
	}

	for range workers {
		select {
		case <-started:
		case <-ctx.Done():
			t.Fatal("worker did not start")
		}
	}

	cancel()

	select {
	case <-ctx.Done():
	case <-release:
		t.Fatal("worker did not stop after context cancellation")
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		t.Fatalf("Result has values: %v, This test is a Failure, we want no values", result)
	}
}
