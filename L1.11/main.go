package main

import "fmt"

// intersect returns the intersection of two unordered sets.
func intersect(set1, set2 []int) []int {
	// Create a map to track the elements of the first set.
	elementMap := make(map[int]struct{})
	for _, num := range set1 {
		elementMap[num] = struct{}{}
	}

	// Check which elements from the second set are in the map.
	var intersection []int
	for _, num := range set2 {
		if _, ok := elementMap[num]; ok {
			intersection = append(intersection, num)
			// Remove from map to avoid duplicates.
			delete(elementMap, num)
		}
	}

	return intersection
}

func main() {
	set1 := []int{1, 5, 4, 8, 10, 55, 41, 0, 0}
	set2 := []int{6, 4, 1, 8, 9, -1, 2, 14, 68, 1, 10}

	result := intersect(set1, set2)
	fmt.Println("Intersection of sets:", result)
}

/*
 - Output: -
Intersection of sets: [4 1 8 10]
*/
