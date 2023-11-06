package unraveling_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	// Create an integer channel
	ch := make(chan int)

	// Goroutine 1: Send data to the channel
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // Send values 1 to 5 to the channel
			time.Sleep(100 * time.Millisecond)
		}
		close(ch) // Close the channel when done
	}()

	// Goroutine 2: Receive data from the channel
	go func() {
		for value := range ch {
			fmt.Printf("Received: %d\n", value)
		}
	}()

	// Allow time for the goroutines to complete
	time.Sleep(1 * time.Second)
}
