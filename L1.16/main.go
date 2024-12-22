package main

import (
	"fmt"
)

// quicksort sorts the array in-place using the QuickSort algorithm.
func quicksort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index.
		pivotIndex := partition(arr, low, high)

		// Recursively sort the left and right subarrays.
		quicksort(arr, low, pivotIndex-1)
		quicksort(arr, pivotIndex+1, high)
	}
}

// partition rearranges elements around the pivot and returns its final index.
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Choose the last element as pivot.

	i := low - 1 // Pointer for the smaller element.

	for j := low; j < high; j++ {
		// If the current element is less than or equal to the pivot.
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap.
		}
	}

	// Place the pivot in the correct position.
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{1, 1, 2, 3, 5, 1, 2, -1}
	fmt.Println("Before sorting:", arr)

	quicksort(arr, 0, len(arr)-1)

	fmt.Println("After sorting:", arr)
}

/*
 - Output: -
Before sorting: [1 1 2 3 5 1 2 -1]
After sorting: [-1 1 1 1 2 2 3 5]
*/
