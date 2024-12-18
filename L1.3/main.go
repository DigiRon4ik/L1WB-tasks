package main

import (
	"fmt"
	"sync"
)

// Function for calculating squares of numbers and sending them to the channel.
func generateSquares(nums []int, ch chan<- int, wg *sync.WaitGroup) {
	// Decrease the WaitGroup counter when the function is finished.
	defer wg.Done()
	// Iterate over each number in the input slice.
	for _, num := range nums {
		// Send the square of the number to the channel.
		ch <- num * num
	}
}

func main() {
	// A sequence of numbers to find the sum of squares.
	numbers := []int{2, 4, 6, 8, 10}
	// Create an unbuffered channel for sending squares.
	// If the channel closes immediately after squares generation is complete, buffering becomes redundant.
	squareChan := make(chan int)
	// Variable for sum of squares.
	var sqrSum int
	// WaitGroup to synchronize goroutines.
	var wg sync.WaitGroup
	// Mutex to safely update the shared sqrSum variable.
	var mu sync.Mutex

	// Increment WaitGroup counter for the first goroutine.
	wg.Add(1)
	// Start of the first goroutine with an anonymous function.
	go func() {
		// To close the channel after all squares have been sent to the channel.
		defer close(squareChan)
		// // Call generateSquares in a goroutine
		generateSquares(numbers, squareChan, &wg)
	}()

	// Increment WaitGroup counter for the second goroutine
	wg.Add(1)
	// Start of the second goroutine with an anonymous function.
	go func() {
		// Decrement the WaitGroup counter when the goroutine finishes
		defer wg.Done()
		// Read squares from the channel
		for square := range squareChan {
			// Lock the mutex to ensure exclusive access to sqrSum
			mu.Lock()
			// Add the square to the total sum
			sqrSum += square
			// Unlock the mutex to allow access by other goroutines
			mu.Unlock()
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	// Print the final sum of squares
	fmt.Println(sqrSum)
}
