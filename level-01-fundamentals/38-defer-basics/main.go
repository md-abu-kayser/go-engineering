// Lesson 38: defer Basics
//
// Goal: Use `defer` to schedule a function call to run when the
// surrounding function returns — and understand the genuinely
// surprising rule that a deferred call's ARGUMENTS are evaluated
// IMMEDIATELY, at the defer statement itself, not later when it runs.
package main

import "fmt"

func main() {
	fmt.Println("=== defer Basics ===")
	fmt.Println("----------------------------------")

	// A deferred call runs when the SURROUNDING FUNCTION returns — not
	// immediately, and not at any point you'd expect from reading top
	// to bottom.
	fmt.Println("1. This prints first.")
	defer fmt.Println("3. This prints LAST — scheduled by defer, runs when main() returns.")
	fmt.Println("2. This prints second.")

	// THE classic gotcha: a deferred function call's ARGUMENTS are
	// evaluated RIGHT NOW, at the defer statement — only the CALL
	// itself is postponed, not the evaluation of its arguments.
	fmt.Println("\n--- Arguments are evaluated immediately, at defer time ---")
	n := 1
	defer fmt.Printf("deferred: n was %d AT THE TIME OF defer (not later)\n", n)
	n = 100
	fmt.Printf("n is now %d (changed AFTER the defer statement ran)\n", n)

	// If you genuinely need the LATEST value, wrap the call in a
	// closure — now the closure captures the VARIABLE, and reads its
	// value only when the closure itself finally runs.
	fmt.Println("\n--- Using a closure to capture the LATEST value instead ---")
	m := 1
	defer func() {
		fmt.Printf("deferred (closure): m is %d — the value AT THE TIME defer actually RUNS\n", m)
	}()
	m = 100
	fmt.Printf("m is now %d\n", m)
}
