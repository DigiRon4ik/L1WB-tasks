package main

import (
	"fmt"
	"time"
)

// Sleep causes the program to wait the specified number of milliseconds
func Sleep(d time.Duration) {
	// Create a timer and block execution until its completion
	<-time.After(d)
}

func main() {
	fmt.Println("Start")
	Sleep(2 * time.Second) // 2-second delay
	fmt.Println("End")
}

/*
 - Output: -
Start
End
*/
