package main

import (
	"fmt"
)

func main() {
	a := 20
	b := 10

	if a > b {
		fmt.Printf("%v is greater than %v\n", a, b)
	} else if a < b {
		fmt.Printf("%v is less than %v\n", a, b)
	} else {
		fmt.Printf("%v is equal to %v\n", a, b)
	}

	fmt.Println("Basic Calculator Operations:")
	operation := "add"
	switch operation {
	case "add":
		fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	case "subtract":
		fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
	case "multiply":
		fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
	case "divide":
		if b != 0 {
			fmt.Printf("Division: %d / %d = %d\n", a, b, a/b)
		} else {
			fmt.Println("Division by zero is not allowed")
		}
	default:
		fmt.Println("Invalid operation")
	}

	//for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//for loop with range
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	//for loops as while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// // for loop as infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }
}
