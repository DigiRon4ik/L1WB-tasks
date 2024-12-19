package main

import "fmt"

// Swapping variables. Just a trick from Python.
func main() {
	a, b := 1, 2
	fmt.Println(a, b)
	// First the right part is calculated, then it is assigned to the left part.
	a, b = b, a
	fmt.Println(a, b)
}

/*
 - Output: -
1 2
2 1
*/
