package main

import (
	"fmt"
	"sync"
)

// A function to calculate the square of a number.
// Writes the square to the slice, and tells WaitGroup that the goroutine is complete.
func square(i, num int, res []int, wg *sync.WaitGroup) {
	// After the function completes, wg.Done() reports that the goroutine is complete.
	defer wg.Done()
	// Writes to the slice at the index square of a number
	// because all copies of the same slice point to the same array in memory.
	res[i] = num * num
}

func main() {
	// Array of numbers whose squares should be calculated and output.
	numbers := [5]int{2, 4, 6, 8, 10}

	// A slice with results, i.e. squares of numbers.
	squares := make([]int, len(numbers))
	// WaitGroup is needed to wait for all goroutines to complete without terminating the main program.
	var wg sync.WaitGroup

	// Going through the numbers in the array.
	for i, num := range numbers {
		// Add 1 to the number of goroutines to wait for
		wg.Add(1)
		// The goroutine is started based on the square function.
		go square(i, num, squares, &wg)
	}

	// Awaiting completion of all goroutines.
	// The program will not move until it catches an error or all goroutines are completed correctly.
	wg.Wait()

	// Output squares.
	for _, sqr := range squares {
		fmt.Println(sqr)
	}
}

/*
 - Output: -
4
16
36
64
100
*/
