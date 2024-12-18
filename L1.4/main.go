package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Check if the number of arguments is less than 2 (program name + argument).
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number_of_workers>")
		return
	}

	// Parses the second argument to match the number of workers.
	numWorkers, err := strconv.Atoi(os.Args[1])
	// Validate input to ensure it's a positive integer.
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers. Please provide a positive integer.")
		return
	}

	// A channel for sending data to vorkers.
	dataChannel := make(chan int)
	// A channel to notify workers and the manufacturer of a shutdown.
	stopChannel := make(chan struct{})
	// Channel for listening to OS interrupt signals (e.g. Ctrl+C).
	signalChan := make(chan os.Signal, 1)
	// Subscribe to interrupt signals (SIGINT, SIGTERM).
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	// A WaitGroup to synchronize the shutdown of all workers.
	var wg sync.WaitGroup

	// Launch the specified number of worker goroutines.
	for i := 1; i <= numWorkers; i++ {
		// Increase the WaitGroup counter for each worker.
		wg.Add(1)
		// Start a worker goroutine.
		go worker(i, dataChannel, stopChannel, &wg)
	}

	// Start the producer goroutine.
	go producer(dataChannel, stopChannel)

	// Wait for an interrupt signal (e.g. Ctrl+C).
	<-signalChan
	fmt.Println("\nReceived interrupt signal. Shutting down...")
	// Close the signal channel to clean up resources.
	close(signalChan)

	// Notify the workers and producer that the job is over by closing the stopChannel channel.
	close(stopChannel)
	// Wait for all workers to finish.
	wg.Wait()
	// Close the data channel as no more data will be sent.
	close(dataChannel)
	fmt.Println("All workers stopped. Exiting.")
}

// Worker goroutine that processes data from the dataChannel.
func worker(id int, dataChannel <-chan int, stopChannel <-chan struct{}, wg *sync.WaitGroup) {
	// Decrease the WaitGroup counter when the worker at job completion.
	defer wg.Done()
	// Infinity cycle.
	for {
		select {
		// Listen for the stop signal.
		case <-stopChannel:
			fmt.Printf("Worker %d stopping.\n", id)
			return
		// Receive data from the data channel.
		case data, ok := <-dataChannel:
			// If the data channel is closed, terminate the worker.
			if !ok {
				fmt.Printf("Worker %d detected channel closure. Exiting.\n", id)
				return
			}
			// Process the received data.
			fmt.Printf("Worker %d received: %d\n", id, data)
		}
	}
}

// Producer goroutine that generates data and sends it to the dataChannel.
func producer(dataChannel chan int, stopChannel chan struct{}) {
	i := 0 // Initialize the counter for generated data.
	// Infinity cycle.
	for {
		select {
		// Listen for the stop signal.
		case <-stopChannel:
			fmt.Println("Producer stopped.")
			return
		// If no stop signal, produce data.
		default:
			i++                                // Increment the data counter.
			dataChannel <- i                   // Send the data to the dataСhannel.
			time.Sleep(500 * time.Millisecond) // Simulate production delay.
		}
	}
}

/*
 - Input : -
go run L1.4/main.go 4
 - Output: -
Worker 2 received: 1
Worker 3 received: 2
Worker 4 received: 3
Worker 1 received: 4
Worker 2 received: 5
Worker 3 received: 6
Worker 4 received: 7
Worker 1 received: 8
Worker 2 received: 9
Worker 3 received: 10
Worker 4 received: 11
Worker 1 received: 12
Worker 2 received: 13
Worker 3 received: 14
Worker 4 received: 15
Worker 1 received: 16
Worker 2 received: 17
Worker 3 received: 18
Worker 4 received: 19
Worker 1 received: 20
^C
Received interrupt signal. Shutting down...
Worker 2 stopping.
Worker 1 stopping.
Worker 4 stopping.
Worker 3 stopping.
All workers stopped. Exiting.
*/
