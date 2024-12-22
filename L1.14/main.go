package main

import (
	"fmt"
)

func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("The type of the variable is: int")
	case string:
		fmt.Println("The type of the variable is: string")
	case bool:
		fmt.Println("The type of the variable is: bool")
	case chan int:
		fmt.Println("The type of the variable is: channel (chan int)")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	// Example variables
	var a int = 42
	var b string = "Hello, Go!"
	var c bool = true
	var d chan int = make(chan int)

	// Check variable types
	determineType(a) // Expected: The type of the variable is: int
	determineType(b) // Expected: The type of the variable is: string
	determineType(c) // Expected: The type of the variable is: bool
	determineType(d) // Expected: The type of the variable is: channel (chan int)

	// Example with an unknown type
	var e float64 = 3.14
	determineType(e)
}

/*
 - Output: -
The type of the variable is: int
The type of the variable is: string
The type of the variable is: bool
The type of the variable is: channel (chan int)
Unknown type
*/
