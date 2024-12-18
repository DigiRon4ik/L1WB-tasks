package main

import (
	"context"
	"fmt"
	"time"
)

// Stopping goroutine time work contex or deadline context.
func main() {
	fmt.Println("Solution-3")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go worker(ctx)

	time.Sleep(4 * time.Second)
	cancel() // Explicitly cancel (even though the timeout is about to expire)
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
Solution-3
Working...
Working...
Working...
Working...
Worker stopped:  context deadline exceeded
*/