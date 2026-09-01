// Lesson 28: Expression Switches
//
// Goal: Use "tagless" switch — switch with NO expression after the
// keyword — as a cleaner alternative to a long if/else-if chain
// (lesson 26), plus a switch WITH an init statement.
package main

import "fmt"

// classify uses a TAGLESS switch — `switch` with nothing after it,
// which is exactly equivalent to `switch true`. Each `case` is a full
// boolean expression, evaluated top to bottom until one matches.
func classify(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "F"
	}
}

func main() {
	fmt.Println("=== Expression Switches ===")
	fmt.Println("----------------------------------")

	for _, score := range []int{95, 82, 71, 40} {
		fmt.Printf("score %3d -> grade %s (via tagless switch)\n", score, classify(score))
	}

	// Compare this directly to lesson 26's if/else-if/else chain for
	// the EXACT same logic — many Go developers find the tagless
	// switch reads more cleanly once there are more than two or three
	// conditions.
	fmt.Println("\n--- Same logic, if/else-if/else style (lesson 26) vs switch (this lesson) ---")
	fmt.Println("Both are equally valid Go — this is a genuine STYLE choice, not a rule.")

	// switch ALSO supports an init statement, exactly like if (lesson 25).
	fmt.Println("\n--- switch with an init statement ---")
	switch length := len("Gopher"); {
	case length > 10:
		fmt.Println("long")
	case length > 3:
		fmt.Println("medium")
	default:
		fmt.Println("short")
	}
}
