// Lesson 17: Operator Precedence
//
// Goal: See Go's operator precedence table applied directly — including
// a couple of expressions that read one way but EVALUATE another,
// making the case for parentheses even when technically optional.
package main

import "fmt"

func main() {
	fmt.Println("=== Operator Precedence ===")
	fmt.Println("----------------------------------")

	// * binds TIGHTER than + — exactly like ordinary arithmetic (PEMDAS).
	result1 := 2 + 3*4
	fmt.Printf("2 + 3 * 4        = %d (NOT 20 — * happens first)\n", result1)

	// Comparison operators (<, >, ==, ...) bind LOOSER than arithmetic
	// operators (+, -, *, /) — so arithmetic always happens before the
	// comparison, with no parentheses needed.
	result2 := 1+2 == 3
	fmt.Printf("1 + 2 == 3       = %t (arithmetic happens BEFORE the comparison)\n", result2)

	// && binds TIGHTER than || — this one genuinely surprises people
	// coming from some other backgrounds' mental models. Using
	// VARIABLES (not bare literals) here so the expression is
	// evaluated for real, rather than being flagged by static analysis
	// as a foregone conclusion.
	c1, c2, c3 := true, false, false
	result3 := c1 || c2 && c3
	fmt.Printf("true || false && false = %t (&& binds tighter: true || (false && false))\n", result3)

	// Bitwise operators have their OWN precedence levels, interleaved
	// with arithmetic — this is a common source of subtle bugs, since
	// it's easy to assume they're all "the same level".
	result4 := 1 | 2&3 // & binds tighter than |, so this is 1 | (2 & 3)
	fmt.Printf("1 | 2 & 3        = %d (& binds tighter: 1 | (2 & 3) = 1 | 2 = 3)\n", result4)

	fmt.Println("\n--- When parentheses are technically unnecessary but STILL a good idea ---")
	// This expression is UNAMBIGUOUS to the compiler, but arguably NOT
	// unambiguous to a human reading it quickly.
	confusing := 1+2 == 3 && 4-1 == 3
	clear := (1+2 == 3) && (4-1 == 3) // IDENTICAL result, much easier to read
	fmt.Printf("without parens: %t, with parens: %t (same result, but which was easier to read?)\n", confusing, clear)
}
