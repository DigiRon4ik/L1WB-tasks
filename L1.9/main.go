package main

import (
	"fmt"
)

func main() {
	// Create two channels.
	inputChan := make(chan int)
	outputChan := make(chan int)

	// Array of numbers.
	numbers := []int{1, 2, 3, 4, 5}

	// Goroutine for sending numbers to the first channel.
	go func() {
		for _, num := range numbers {
			inputChan <- num
		}
		close(inputChan) // Close the channel after sending all numbers.
	}()

	// Goroutine for reading from the first channel, multiplying by 2, and sending to the second channel.
	go func() {
		for num := range inputChan {
			outputChan <- num * 2
		}
		close(outputChan) // Close the channel after processing all numbers.
	}()

	// Read from the second channel and print to stdout.
	for result := range outputChan {
		fmt.Println(result)
	}
}

// The order will always be correct.
/*
 - Output: -
2
4
6
8
10
*/
