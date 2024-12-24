package main

import (
	"fmt"
)

// reverseString takes a string and returns a reversed version.
func reverseString(input string) string {
	// Convert the string to a slice of runes for correct Unicode handling.
	runes := []rune(input)
	left, right := 0, len(runes)-1

	// Reverse the runes.
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}

func main() {
	input := "главрыба"
	reversed := reverseString(input)
	fmt.Printf("Original: %s\nReversed: %s\n", input, reversed)
}

/*
 - Output: -
Original: главрыба
Reversed: абырвалг
*/
