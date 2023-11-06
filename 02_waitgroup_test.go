package unraveling_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Printf("Worker %d is done\n", id)
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add each goroutine to the WaitGroup
		go worker(i, &wg)
	}

	wg.Wait() // Block until the internal counter reaches zero
	fmt.Println("All workers have completed.")
}
