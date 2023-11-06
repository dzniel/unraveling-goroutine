package unraveling_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func regularFunction() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Regular Goroutine: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func TestGoroutines(t *testing.T) {
	// Starting a new goroutine using a regular function
	go regularFunction()

	// Starting a new goroutine using an anonymous function
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Anonymous Goroutine: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Allow the goroutines to run for a while
	time.Sleep(1 * time.Second)

	fmt.Println("Main function completed.")
}
