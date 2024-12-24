package main

import (
	"fmt"
	"strings"
)

// Function to check if all characters in a string are unique (case-insensitive)
func areAllCharactersUnique(input string) bool {
	// Convert the string to lowercase for case-insensitive comparison
	input = strings.ToLower(input)

	// Create a map to track characters
	charMap := make(map[rune]bool)

	// Check each character in the string
	for _, char := range input {
		// If the character is already in the map, return false
		if charMap[char] {
			return false
		}
		// Otherwise, add the character to the map
		charMap[char] = true
	}

	// If all characters are unique, return true
	return true
}

func main() {
	// Test cases
	testCases := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, testCase := range testCases {
		fmt.Printf("String: %s, All Unique: %t\n", testCase, areAllCharactersUnique(testCase))
	}
}

/*
 - Output: -
String: abcd, All Unique: true
String: abCdefAaf, All Unique: false
String: aabcd, All Unique: false
*/
