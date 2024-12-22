package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		// Calculate the middle index.
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid // Element found, return index.
		}

		if nums[mid] < target {
			left = mid + 1 // Search in the right half.
		} else {
			right = mid - 1 // Search in the left half.
		}
	}

	return -1 // Element not found.
}

func main() {
	nums := []int{1, 3, 5, 7, 9, 11}
	target := 7

	index := binarySearch(nums, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d.\n", target, index)
	} else {
		fmt.Printf("Element %d not found.\n", target)
	}
}

/*
 - Output: -
Element 7 found at index 3.
*/
