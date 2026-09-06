// Lesson 48: Anonymous Functions
//
// Goal: Write a FUNCTION LITERAL — a function with no name, defined
// right where it's used — and call one immediately (an "IIFE"). Note:
// this lesson is about the LITERAL syntax itself; capturing outer
// variables (what makes one a "closure") is lesson 49's dedicated topic.
package main

import "fmt"

func main() {
	fmt.Println("=== Anonymous Functions ===")
	fmt.Println("----------------------------------")

	// A function literal, assigned to a variable — identical in
	// spirit to lesson 47's function values, except the function
	// itself was never given a name via `func name(...)`.
	square := func(n int) int {
		return n * n
	}
	fmt.Printf("square(5) = %d\n", square(5))

	// Passed DIRECTLY as an argument, with no intermediate variable at
	// all — extremely common for one-off logic that doesn't deserve
	// its own top-level, named function.
	fmt.Println("\n--- Passed directly as an argument ---")
	numbers := []int{1, 2, 3, 4, 5}
	doubled := mapInts(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("mapInts(%v, double) = %v\n", numbers, doubled)

	// An IIFE — Immediately Invoked Function Expression: define AND
	// call a function literal in one statement, using it to scope some
	// setup logic without polluting the surrounding function with
	// extra named helper functions or leftover variables.
	fmt.Println("\n--- Immediately invoked (IIFE) ---")
	result := func(a, b int) int {
		sum := a + b
		return sum * sum
	}(3, 4) // called IMMEDIATELY, right here, with (3, 4)
	fmt.Printf("IIFE result = %d ((3+4)^2)\n", result)
}

// mapInts applies fn to every element of nums, returning a new slice
// of the results — a small, real use for accepting a function value
// (lesson 47) that's about to receive an ANONYMOUS one as its argument.
func mapInts(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}
