package workerpool

import (
	"sync"
	"testing"
)

type Result struct {
	Input  int
	Output int
}

func worker(
	jobs <-chan int,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	var res Result
	for job := range jobs {
		res.Input = job
		res.Output = job * job

		results <- res
	}

}

func TestWorkerPool(t *testing.T) {
	const workers = 2
	const jobsCount = 4

	var wg sync.WaitGroup

	jobs := make(chan int, 4)

	results := make(chan Result, 4)

	for i := range jobsCount {
		jobs <- i
	}

	for range workers {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	finalResultsMap := make(map[int]int)

	for result := range results {
		if _, ok := finalResultsMap[result.Input]; ok {
			t.Fatalf("Duplicate input detected! '%d' is already recorded in the map", result.Input)
		}
		finalResultsMap[result.Input] = result.Output
	}

	if len(finalResultsMap) != jobsCount {
		t.Fatalf("Expected %d results got %d instead", jobsCount, len(finalResultsMap))
	}

	for i := range jobsCount {
		val, ok := finalResultsMap[i]
		if !ok {
			t.Fatalf("key %d doesn't exist", i)
		}
		expectedSquare := i * i
		if val != expectedSquare {
			t.Fatalf("Input %d doesn't have correct output, output: %d, want %d", i, val, expectedSquare)
		}
	}
}
