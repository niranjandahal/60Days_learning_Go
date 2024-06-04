package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

// Function with multiple return values
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}

// Function with named return values
func divideNamed(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}

func main() {
	sum := add(3, 4)
	fmt.Println("Sum:", sum)

	// Calling the divide function with multiple return values
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)

	// Calling the divideNamed function with named return values
	q, r := divideNamed(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

}
