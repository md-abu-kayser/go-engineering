// Lesson 25: if with Init
//
// Goal: Use Go's "if with a short init statement" form — extremely
// common with functions returning (value, error) or (value, ok) — and
// understand exactly what scope the init variable lives in.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== if with Init ===")
	fmt.Println("----------------------------------")

	// The IDIOMATIC Go pattern: run a statement, then immediately
	// check its result, all in one line. The semicolon separates the
	// INIT statement from the CONDITION.
	if n, err := strconv.Atoi("123"); err == nil {
		fmt.Printf("Parsed successfully: %d\n", n)
	} else {
		fmt.Printf("Failed to parse: %v\n", err)
	}

	// Same pattern, but this time the parse genuinely fails — note the
	// else branch can still use `err` (and even `n`, though it's 0 here)
	// since they're in scope for the WHOLE if/else chain.
	if n, err := strconv.Atoi("not a number"); err == nil {
		fmt.Printf("Parsed successfully: %d\n", n)
	} else {
		fmt.Printf("Failed to parse %q: %v\n", "not a number", err)
	}

	// THE SCOPE RULE: n and err declared in the init statement are
	// visible ONLY inside the if/else chain — NOT outside it. Trying to
	// use them here would be a compile error (see README for the exact
	// error message).
	fmt.Println("\nSee the README: `n` and `err` above are NOT visible out here.")

	// This pattern is especially common for avoiding namespace clutter
	// when you only need a value WITHIN the if — no separate variable
	// left lying around afterward for the rest of the function to trip over.
	if length := len("Gopher"); length > 3 {
		fmt.Printf("\n\"Gopher\" has %d characters — that's more than 3.\n", length)
	}
}
