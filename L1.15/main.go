package main

import (
	"fmt"
)

func createHugeString(size int) string {
	// Create a string of the specified size.
	return string(make([]byte, size))
}

func someFunc() string {
	v := createHugeString(1 << 10) // Create a string of size 1024 bytes.
	// Create a new string, copying only the first 100 characters.
	return string([]byte(v[:100]))
}

func main() {
	result := someFunc()
	fmt.Println("Result length:", len(result))
}

/*
 - Output: -
Result length: 100
*/

// Memory leak. Strings in Go are immutable and are stored as slices of bytes.
// When the v[:100] operation is executed, a new slice is created that references part of
// the original string v. However, the original string v is not freed from memory because the
// new slice still references its data.
