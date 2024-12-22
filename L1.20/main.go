package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	// Split the string into words.
	words := strings.Fields(input)

	// Reverse the order of words.
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the words back into a string.
	return strings.Join(words, " ")
}

func main() {
	input := "snow  dog    sun"
	result := reverseWords(input)

	fmt.Printf("Original: %s\nReversed: %s\n", input, result)
}

/*
 - Output: -
Original: snow  dog    sun
Reversed: sun dog snow
*/
