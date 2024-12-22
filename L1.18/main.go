package main

import (
	"fmt"
	"sync"
)

// Counter is a structure representing a thread-safe counter.
type Counter struct {
	mu      sync.Mutex // Mutex to protect the counter.
	counter int        // The actual counter value.
}

// Increment increments the counter safely in a concurrent environment.
func (c *Counter) Increment() {
	c.mu.Lock()         // Lock the mutex to ensure exclusive access.
	defer c.mu.Unlock() // Unlock the mutex after updating the counter.
	c.counter++         // Increment the counter.
}

// Value returns the current value of the counter.
func (c *Counter) Value() int {
	c.mu.Lock()         // Lock the mutex to safely access the counter.
	defer c.mu.Unlock() // Unlock the mutex after reading the counter value.
	return c.counter    // Return the current counter value.
}

func main() {
	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish.
	counter := &Counter{} // Create a new Counter instance.

	// Number of goroutines to run concurrently.
	const numGoroutines = 100
	// Number of increments per goroutine.
	const numIncrements = 1000

	// Start multiple goroutines.
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Increment the WaitGroup counter.
		go func() {
			defer wg.Done() // Decrement the WaitGroup counter when done.
			for j := 0; j < numIncrements; j++ {
				counter.Increment() // Safely increment the counter.
			}
		}()
	}

	// Wait for all goroutines to complete.
	wg.Wait()

	// Output the final value of the counter.
	fmt.Printf("Final Counter Value: %d\n", counter.Value())
}

/*
 - Output: -
Final Counter Value: 100000
*/
