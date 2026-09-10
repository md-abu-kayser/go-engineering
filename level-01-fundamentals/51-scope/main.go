// Lesson 51: Scope
//
// Goal: See exactly where a variable is visible — a BLOCK ({ }) is the
// fundamental unit of scope in Go, and every variable is visible from
// its declaration to the end of the innermost block containing it.
package main

import "fmt"

// classifyScore is a PURE function (no side effects, same input always
// gives the same output) — pulled out specifically so this lesson has
// something real to unit test, demonstrating scope rules through
// actual passing behavior, not just printed output.
func classifyScore(score int) string {
	// `label` is scoped to this FUNCTION BODY — visible everywhere
	// inside classifyScore, but nowhere outside it.
	label := "unknown"

	if score >= 90 {
		// `bonus` is scoped to THIS if-block only — it does not exist
		// even a single line after the closing brace below.
		bonus := " (with honors)"
		label = "A" + bonus
	} else if score >= 70 {
		label = "B"
	}

	return label
}

func main() {
	fmt.Println("=== Scope ===")
	fmt.Println("----------------------------------")

	for _, score := range []int{95, 75, 40} {
		fmt.Printf("classifyScore(%d) = %s\n", score, classifyScore(score))
	}

	// Demonstrating block scope directly: `x` declared inside the `for`
	// loop's block is completely gone once the loop ends.
	fmt.Println("\n--- Block scope, made visible ---")
	for i := 0; i < 3; i++ {
		x := i * i
		fmt.Printf("  inside loop: i=%d, x=%d\n", i, x)
	}
	fmt.Println("  outside the loop: neither i nor x exist here anymore")

	// Nested blocks: an inner block can see everything an outer block
	// declared before it — but NOT the other way around.
	fmt.Println("\n--- Nested blocks see outward, never inward ---")
	outer := "I'm from the outer block"
	{
		fmt.Printf("  inner block CAN see: %q\n", outer)
		inner := "I'm from the inner block"
		fmt.Printf("  inner block can also see its own: %q\n", inner)
	}
	fmt.Println("  outer block canNOT see `inner` — it never existed out here")
}
