package contextdemo

import (
	"context"
	"testing"
	"time"
)

func TestWorkerStopsWhenContextCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})
	started := make(chan struct{})
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	go func() {
		defer close(done)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				select {
				case <-started:
					// The first tick was already reported
				default:
					close(started)
				}
			}
		}
	}()

	select {
	case <-started:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("worker did not start")
	}

	cancel()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("worker did not stop after context cancellation")
	}

	if ctx.Err() != context.Canceled {
		t.Fatalf("context error = %v, want %v", ctx.Err(), context.Canceled)
	}

}
