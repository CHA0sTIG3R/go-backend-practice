package workerpool

import (
	"context"
	"sync"
	"testing"
)

const WORKERS = 2
const JOBSCOUNT = 4

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

	jobs := make(chan int, JOBSCOUNT)
	results := make(chan Result, JOBSCOUNT)
	started := make(chan int, WORKERS)
	release := make(chan struct{})

	for i := range JOBSCOUNT {
		jobs <- i
	}

	close(jobs)

	for range WORKERS {
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

	acquired := make(map[int]bool, WORKERS)

	for range WORKERS {
		job := <-started

		if job < 0 || job >= JOBSCOUNT {
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

func TestWorkersCompleteWhenReleased(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	jobs := make(chan int, JOBSCOUNT)
	results := make(chan Result, JOBSCOUNT)
	started := make(chan int, WORKERS)
	release := make(chan struct{})

	for i := range JOBSCOUNT {
		jobs <- i
	}
	close(jobs)

	for range WORKERS {
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

	acquired := make(map[int]bool, WORKERS)

	for range WORKERS {
		job := <-started

		if job < 0 || job >= JOBSCOUNT {
			t.Fatalf("worker acquired invalid job %d", job)
		}
		if acquired[job] {
			t.Fatalf("job %d was acquired more than once", job)
		}

		acquired[job] = true
	}

	if len(acquired) != WORKERS {
		t.Fatalf("number of acquired jobs = %d, want %d", len(acquired), WORKERS)
	}

	close(release)

	go func() {
		wg.Wait()
		close(results)
	}()

	finalResultsMap := make(map[int]int)

	for result := range results {
		t.Logf("input: %d\n output: %d", result.Input, result.Output)

		if !acquired[result.Input] {
			t.Fatalf("worker returned result for unacquired job %d", result.Input)
		}
		if _, ok := finalResultsMap[result.Input]; ok {
			t.Fatalf("worker returned result for already completed job %d", result.Input)
		}
		if result.Output != result.Input*result.Input {
			t.Fatalf("worker returned invalid result %v, want %d", result, result.Input*result.Input)
		}
		finalResultsMap[result.Input] = result.Output
	}

	if len(finalResultsMap) != WORKERS {
		t.Fatalf("number of completed jobs = %d, want %d", len(finalResultsMap), WORKERS)
	}

	for job := range finalResultsMap {
		if !acquired[job] {
			t.Fatalf("job %d was completed but not acquired", job)
		}
		if finalResultsMap[job] != job*job {
			t.Fatalf("job %d was completed with invalid result %d, want %d", job, finalResultsMap[job], job*job)
		}
	}
}
