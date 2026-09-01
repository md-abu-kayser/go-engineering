// Lesson 24: if Statements
//
// Goal: Use Go's `if` statement — no parentheses around the condition,
// mandatory braces, and a condition that must be a genuine bool
// expression (no truthy/falsy values, as lesson 07 already covered).
package main

import "fmt"

func main() {
	fmt.Println("=== if Statements ===")
	fmt.Println("----------------------------------")

	age := 20

	// Notice: NO parentheses around the condition — that would be
	// unusual, unidiomatic Go (though technically still legal, since
	// extra parens around any expression are allowed).
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	// Braces are MANDATORY, even for a single-statement body — unlike
	// C, Java, or JavaScript, where you CAN omit them for one statement.
	// This is a deliberate Go design choice (see README).
	if age < 0 {
		fmt.Println("This is impossible, but the braces are still required.")
	}

	// The condition must be an actual bool expression — comparisons,
	// boolean variables, function calls returning bool, or combinations
	// of these with && / || / !.
	isMember := true
	hasDiscount := age >= 65 || isMember
	if hasDiscount {
		fmt.Println("Eligible for a discount.")
	}

	fmt.Println("\nSee the README for exactly why Go REQUIRES braces, even for one-liners.")
}
