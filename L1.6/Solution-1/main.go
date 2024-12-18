package main

import (
	"fmt"
	"time"
)


// Stopping goroutine by channel.
func main() {
	fmt.Println("Solution-1")

	channel := make(chan struct{})
	go worker(channel)

	time.Sleep(2 * time.Second)
	close(channel)
	// or
	// channel <- struct{}{}
	time.Sleep(500 * time.Millisecond)
}

func worker(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Worker Stopped!")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
 - Output: -
Solution-1
Working...
Working...
Working...
Working...
Worker stopped!
*/