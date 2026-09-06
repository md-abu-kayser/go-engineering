// Lesson 47: Function Values
//
// Goal: Treat functions as first-class VALUES — assign one to a
// variable, declare a variable of a function TYPE, and pass a function
// as an argument to another function (a "higher-order" function).
package main

import "fmt"

func add(a, b int) int      { return a + b }
func multiply(a, b int) int { return a * b }

// applyOp is a HIGHER-ORDER function: one of its parameters is itself
// a function. `op` here has the TYPE `func(int, int) int` — any
// function matching that exact signature can be passed in.
func applyOp(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	fmt.Println("=== Function Values ===")
	fmt.Println("----------------------------------")

	// Functions are values — assign one directly to a variable, just
	// like an int or a string. No special syntax beyond the function's
	// own name (without parentheses — that would CALL it instead).
	var operation func(int, int) int
	operation = add
	fmt.Printf("operation = add;      operation(3, 4) = %d\n", operation(3, 4))

	operation = multiply
	fmt.Printf("operation = multiply; operation(3, 4) = %d\n", operation(3, 4))

	// Passing a function AS AN ARGUMENT to another function — this is
	// what makes applyOp "higher-order": it operates ON functions, not
	// just on plain values.
	fmt.Println("\n--- Passing functions as arguments ---")
	fmt.Printf("applyOp(5, 6, add)      = %d\n", applyOp(5, 6, add))
	fmt.Printf("applyOp(5, 6, multiply) = %d\n", applyOp(5, 6, multiply))

	// A slice of functions — since functions are ordinary values, they
	// can be stored in any data structure a normal value can be.
	fmt.Println("\n--- A slice of functions ---")
	operations := []func(int, int) int{add, multiply}
	for i, op := range operations {
		fmt.Printf("operations[%d](2, 3) = %d\n", i, op(2, 3))
	}
}
