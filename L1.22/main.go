package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Initialization of large numbers a and b
	a := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(21), nil) // a = 2^21
	b := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(22), nil) // b = 2^22

	// Operations
	sum := new(big.Int).Add(a, b) // Addition
	sub := new(big.Int).Sub(a, b) // Subtraction
	mul := new(big.Int).Mul(a, b) // Multiplication
	div := new(big.Int).Div(b, a) // Division

	// Output results
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("Sum (a + b) =", sum)
	fmt.Println("Difference (a - b) =", sub)
	fmt.Println("Product (a * b) =", mul)
	fmt.Println("Quotient (b / a) =", div)
}

/*
 - Output: -
a = 2097152
b = 4194304
Sum (a + b) = 6291456
Difference (a - b) = -2097152
Product (a * b) = 8796093022208
Quotient (b / a) = 2
*/
