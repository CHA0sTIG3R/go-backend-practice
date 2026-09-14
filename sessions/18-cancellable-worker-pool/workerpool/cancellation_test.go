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

	var job int

	select {
	case <-ctx.Done():
		return
	case receivedJob, ok := <-jobs:
		if !ok {
			return
		}
		job = receivedJob
		started <- job
	}

	select {
	case <-ctx.Done():
		return
	case <-release:
		var res Result
		res.Input = job
		res.Output = res.Input * res.Input
		results <- res
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

	acquired := make(map[int]bool, workers)

	for range workers {
		job := <-started

		if job < 0 || job >= jobsCount {
			t.Fatalf("worker acquired invalid job %d", job)
		}
		if acquired[job] {
			t.Fatalf("job %d was acquired more than once", job)
		}

		acquired[job] = true
	}

	cancel()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		t.Fatalf("Result has values: %v, This test is a Failure, we want no values", result)
	}

	if ctx.Err() != context.Canceled {
		t.Fatalf("context error = %v, want %v", ctx.Err(), context.Canceled)
	}
}
