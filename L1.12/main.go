package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := Set(sequence)
	fmt.Println(set)
}

// Set function creates a set for a sequence of strings.
func Set(sequence []string) []string {
	// Create a map to track elements and remove duplicates.
	temp := make(map[string]struct{})
	var set []string
	// Iterate over each element in the sequence.
	for _, v := range sequence {
		// If the element is not in the map, it's unique.
		if _, ok := temp[v]; !ok {
			temp[v] = struct{}{}
			set = append(set, v)
		}
	}
	return set
}

/*
 - Output: -
[cat dog tree]
*/
