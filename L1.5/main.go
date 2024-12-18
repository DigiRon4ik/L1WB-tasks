package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// A channel to send data to the goroutine.
	ch := make(chan string)
	// WaitGroup to wait for the goroutine to finish.
	var wg sync.WaitGroup
	// A number to determine the running time of the goroutine.
	N := 5

	// Increase the number of goroutines by 1 that you have to wait for.
	wg.Add(1)
	// We start a goroutine with an anonymous function that will listen to the channel until it closes.
	go func() {
		// Decrease the WaitGroup counter when the goroutine finishes.
		defer wg.Done()
		// Read data from the channel.
		for value := range ch {
			fmt.Println(value)
		}
	}()

	// Define a new channel, which will automatically receive the current date with time,
	// after a certain time has elapsed.
	to := time.After(time.Duration(N) * time.Second)
	// Infinity cycle.
	for {
		// We'll listen to the channel with the date.
		select {
		// When the date from the channel is read, this case will be triggered.
		case <-to:
			// Closes the channel that is used to read the data in the goroutine above.
			close(ch)
			fmt.Println("Done")
			// After the canal closes, the goroutine will be completed. Here we will wait until it closes correctly.
			wg.Wait()
			// Let's terminate the main function, which will terminate the program accordingly.
			return
		// If the date from the channel is not read, this case will be triggered.
		default:
			// Send data to the channel.
			ch <- "Downloading..."
			// Wait a second to simulate a download.
			time.Sleep(time.Second)
		}
	}
}

/*
 - Output: -
Downloading...
Downloading...
Downloading...
Downloading...
Downloading...
Done
*/
