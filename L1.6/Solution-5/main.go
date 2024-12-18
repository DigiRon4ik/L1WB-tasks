package main

import (
	"fmt"
	"time"
)

// Stopping the goroutine with a channel and timer.
func main() {
	fmt.Println("Solution-5")

	stopChan := make(chan struct{})
	go worker(stopChan)

	time.AfterFunc(time.Duration(2*time.Second), func() {
		stopChan <- struct{}{}
	})

	time.Sleep(5 * time.Second)
}

func worker(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Worker stopped!")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
 - Output: -
Solution-5
Working...
Working...
Working...
Working...
Worker stopped!
*/
