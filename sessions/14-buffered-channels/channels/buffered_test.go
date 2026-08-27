package channels

import "testing"

func TestBufferedChannelCapacity(t *testing.T) {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	t.Logf("Length of channel: %d", len(ch))
	t.Logf("Capacity of channel: %d", cap(ch))

	value := <-ch

	t.Logf("received value from channel: %d", value)
	t.Logf("Length of channel: %d", len(ch))
	t.Logf("Capacity of channel: %d", cap(ch))
}