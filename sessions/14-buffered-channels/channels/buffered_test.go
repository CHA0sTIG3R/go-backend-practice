package channels

import (
	"testing"
)

func TestBufferedChannelCapacity(t *testing.T) {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	if got := len(ch); got != 3 {
		t.Fatalf("buffered values = %d, want 3", got)
	}
	if got := cap(ch); got != 3 {
		t.Fatalf("channel capacity = %d, want 3", got)
	}

	value := <-ch
	if value != 1 {
		t.Fatalf("first value = %d, want 1", value)
	}
	if got := len(ch); got != 2 {
		t.Fatalf("buffered values after receive = %d, want 2", got)
	}

	ch2 := make(chan int, 2)

	ch2 <- 1
	ch2 <- 2

	if got := len(ch2); got != 2 {
		t.Fatalf("buffered values = %d, want 2", got)
	}
	if got := cap(ch2); got != 2 {
		t.Fatalf("channel capacity = %d, want 2", got)
	}
	sent := make(chan struct{})

	go func() {
		ch2 <- 3
		sent <- struct{}{}
	}()

	select {
	case <-sent:
		t.Fatal("third send completed while channel buffer was full")
	default:
		// The full buffer applies backpressure to the third send.
	}

	val := <-ch2
	if val != 1 {
		t.Fatalf("first value after backpressure = %d, want 1", val)
	}

	<-sent

	if got := len(ch2); got != 2 {
		t.Fatalf("buffered values after third send = %d, want 2", got)
	}
	if value := <-ch2; value != 2 {
		t.Fatalf("second value = %d, want 2", value)
	}
	if value := <-ch2; value != 3 {
		t.Fatalf("third value = %d, want 3", value)
	}
}
