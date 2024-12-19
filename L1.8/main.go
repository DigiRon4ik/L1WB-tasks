package main

import (
	"fmt"
)

func main() {
	var number int64 = 42    // Example number (00101010 in binary).
	var bitIndex uint = 27   // Bit index to modify (indexing starts from 0).
	var setToOne bool = true // Set the bit to 1.

	fmt.Printf("Initial number: %d (binary: %064b)\n", number, number)
	number = setBit(number, bitIndex, setToOne)
	fmt.Printf("After setting the %d-th bit to %v: %d (binary: %064b)\n", bitIndex, setToOne, number, number)

	// Example of setting the bit to 0.
	setToOne = false
	number = setBit(number, bitIndex, setToOne)
	fmt.Printf("After setting the %d-th bit to %v: %d (binary: %064b)\n", bitIndex, setToOne, number, number)
}

// setBit sets the i-th bit to 1 or 0.
func setBit(num int64, i uint, value bool) int64 {
	if value {
		// Set the i-th bit to 1.
		num |= 1 << i
	} else {
		// Set the i-th bit to 0.
		num &= ^(1 << i)
	}
	return num
}

/*
- Output: -
Initial number: 42 (binary: 0000000000000000000000000000000000000000000000000000000000101010)
After setting the 27-th bit to true: 134217770 (binary: 0000000000000000000000000000000000001000000000000000000000101010)
After setting the 27-th bit to false: 42 (binary: 0000000000000000000000000000000000000000000000000000000000101010)
*/
