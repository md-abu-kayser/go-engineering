// Lesson 52: Shadowing
//
// Goal: Go deeper on shadowing than lesson 02's brief introduction —
// including the genuinely dangerous real-world case of accidentally
// shadowing `err` inside a nested if-block, and the surprising fact
// that even built-in names like `len` or `true` can be shadowed.
package main

import (
	"fmt"
	"strconv"
)

// process demonstrates the CLASSIC, genuinely common shadowing bug:
// using := inside a nested if-block accidentally creates a NEW `err`,
// leaving the outer `err` variable untouched — silently discarding a
// real error.
func process(input string) error {
	value, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	if value < 0 {
		// BUG: this := creates a NEW, block-scoped `err` and `doubled`,
		// shadowing the outer `err` entirely. The outer `err` (still nil
		// from the successful Atoi above) is NEVER updated by this line.
		doubled, err := computeDoubled(value)
		if err != nil {
			fmt.Printf("  (bug demonstrated: inner err was %v, but caller never sees it)\n", err)
		}
		_ = doubled
	}

	// Because of the shadowing bug above, THIS still returns the
	// original, unaffected `err` — nil — even if the inner block's
	// `err` was actually non-nil.
	return err
}

func computeDoubled(n int) (int, error) {
	if n < -1000 {
		return 0, fmt.Errorf("value %d is too extreme to double safely", n)
	}
	return n * 2, nil
}

func main() {
	fmt.Println("=== Shadowing ===")
	fmt.Println("----------------------------------")

	fmt.Println("--- The classic err-shadowing bug ---")
	err := process("-2000")
	fmt.Printf("process(\"-2000\") returned err = %v (BUG: should have been non-nil!)\n", err)

	// Go allows shadowing even BUILT-IN identifiers like `len`, `true`,
	// or `int` — they are NOT reserved keywords, just ordinary,
	// PRE-DECLARED identifiers in the universe scope, which any inner
	// scope can shadow like anything else.
	fmt.Println("\n--- Even built-ins can be shadowed (rarely a good idea!) ---")
	{
		len := "surprise!" // shadows the built-in len() function, in THIS block only
		fmt.Printf("  len is now a variable holding: %q\n", len)
		// fmt.Println(len("hello")) would NOT compile in here — len is
		// no longer callable as a function inside this block.
	}
	fmt.Println("  outside the block, len is the built-in function again:", len("hello"))
}
