package channels

import (
	"testing"
	"time"
)

func TestSelectResultBeforeTimeout(t *testing.T) {

	resultCh := make(chan string, 1)

	resultCh <- "completed"

	select {
	case result := <-resultCh:
		if result != "completed" {
			t.Fatalf("Channel received value %s, want completed", result)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("took too long to receive")
	}
}

func TestSelectTimeoutWhenNoResultArrives(t *testing.T) {
	resultCh := make(chan string)
	timedOut := false

	select {
	case result := <-resultCh:
		if result != "" {
			t.Fatal("received a result when timeout was expected")
		}
	case <-time.After(10 * time.Millisecond):
		timedOut = true
	}

	if !timedOut {
		t.Fatal("timeout path did not run")
	}

}

func TestSelectBetweenReadyChannels(t *testing.T) {
	primaryCh := make(chan string, 1)
	fallbackCh := make(chan string, 1)

	fallbackCh <- "fallback"
	primaryCh <- "primary"

	select {
	case value := <-primaryCh:
		if value != "primary" {
			t.Fatalf("primary value = %q, want %q", value, "primary")
		}
	case value := <-fallbackCh:
		if value != "fallback" {
			t.Fatalf("fallback value = %q, want %q", value, "fallback")
		}
	}
}
