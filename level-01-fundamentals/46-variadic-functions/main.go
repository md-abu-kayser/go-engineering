// Lesson 46: Variadic Functions
//
// Goal: Declare and call a function that accepts a variable number of
// arguments — the exact mechanism behind fmt.Println itself — including
// "spreading" an existing slice into one with `slice...`.
package main

import "fmt"

// sum accepts ANY number of ints — zero, one, or a hundred — via the
// `...int` syntax. Inside the function, nums is just an ordinary
// []int; the variadic part only affects how CALLERS can invoke it.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// A variadic parameter must be the LAST parameter — you can freely mix
// ordinary parameters before it.
func joinWithPrefix(prefix string, items ...string) string {
	result := prefix
	for _, item := range items {
		result += " " + item
	}
	return result
}

func main() {
	fmt.Println("=== Variadic Functions ===")
	fmt.Println("----------------------------------")

	// Call with zero, one, or several individual arguments — all valid.
	fmt.Printf("sum()        = %d\n", sum())
	fmt.Printf("sum(5)       = %d\n", sum(5))
	fmt.Printf("sum(1,2,3,4) = %d\n", sum(1, 2, 3, 4))

	// Mixing an ordinary parameter with a variadic one.
	fmt.Println()
	fmt.Println(joinWithPrefix("Items:", "apple", "banana", "cherry"))

	// SPREADING an existing slice into a variadic call with `...` —
	// this is how you pass an already-built slice where individual
	// arguments are expected.
	fmt.Println("\n--- Spreading a slice with ... ---")
	nums := []int{10, 20, 30, 40}
	fmt.Printf("sum(nums...) = %d (spread an existing []int)\n", sum(nums...))

	// fmt.Println/Printf are THEMSELVES variadic — this is not a
	// special case just for this lesson's own functions.
	fmt.Println("\n--- fmt.Println is variadic too ---")
	fmt.Println("this", "call", "has", "five", "separate", "arguments")
}
