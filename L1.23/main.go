package main

import "fmt"

func main() {
	// Example slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", slice)

	// Index of the element to remove
	i := 2 // We will remove the element at index 2 (value 3)

	// Removing the i-th element
	slice = append(slice[:i], slice[i+1:]...)

	fmt.Println("Slice after removal:", slice)
}

/*
 - Output: -
Original slice: [1 2 3 4 5]
Slice after removal: [1 2 4 5]
*/
