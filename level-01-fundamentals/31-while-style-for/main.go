// Lesson 31: While-Style for
//
// Goal: Use `for` with ONLY a condition — Go's equivalent of a `while`
// loop, since Go has no separate `while` keyword at all.
package main

import "fmt"

func main() {
	fmt.Println("=== While-Style for ===")
	fmt.Println("----------------------------------")

	// Dropping the init and post parts (and their semicolons) leaves
	// just a condition — this IS Go's "while" loop. No separate
	// keyword exists; this is simply another shape of `for`.
	n := 1
	for n < 20 {
		fmt.Printf("n = %d\n", n)
		n *= 2
	}

	// A realistic use: reading/processing until some condition is met
	// — here, simulating "keep halving a number until it's small enough".
	fmt.Println("\n--- A more realistic example ---")
	value := 100
	steps := 0
	for value > 1 {
		value /= 2
		steps++
	}
	fmt.Printf("Halved 100 down to %d in %d steps.\n", value, steps)

	// Comparing this DIRECTLY to how you'd write "while" in a language
	// that has a separate keyword for it — the ONLY syntactic
	// difference is the missing `while` keyword itself; the shape of
	// the loop is otherwise identical.
	fmt.Println("\nOther languages: while (n < 20) { ... }")
	fmt.Println("Go:              for n < 20 { ... }        <- same idea, no separate keyword")
}
