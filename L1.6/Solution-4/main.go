package main

import (
	"fmt"
	"time"
)

var running = true

// Stopping goroutine by a logical condition.
func main() {
	fmt.Println("Solution-4")

	go worker()

	time.Sleep(2 * time.Second)
	running = false
	time.Sleep(500 * time.Millisecond)
}

func worker() {
	for running {
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Worker stopped!")
}

/*
 - Output: -
Solution-4
Working...
Working...
Working...
Working...
Worker stopped!
*/