package main

import (
	"context"
	"fmt"
	"time"
)

// Stopping goroutine by context.
func main() {
	fmt.Println("Solution-2")

	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped: ", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*
 - Output: -
Solution-2
Working...
Working...
Working...
Working...
Worker stopped:  context canceled
*/