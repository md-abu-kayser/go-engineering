// Lesson 23: Increment & Decrement
//
// Goal: Use ++ and -- correctly — and understand the genuinely
// Go-specific restriction that trips up people coming from C, Java,
// JavaScript, or similar languages: ++ and -- are STATEMENTS in Go,
// never expressions.
package main

import "fmt"

func main() {
	fmt.Println("=== Increment & Decrement ===")
	fmt.Println("----------------------------------")

	n := 5
	fmt.Printf("n := 5    -> n = %d\n", n)

	n++ // equivalent to n += 1, which is equivalent to n = n + 1
	fmt.Printf("n++       -> n = %d\n", n)

	n--
	n-- // two separate decrements
	fmt.Printf("n--; n--  -> n = %d\n", n)

	// A common loop usage — though note lesson 20's material on
	// for-loop syntax hasn't been covered yet; this is just ++ in
	// isolation, used repeatedly.
	fmt.Println("\n--- Counting up with ++ ---")
	count := 0
	for count < 3 {
		fmt.Printf("  count = %d\n", count)
		count++
	}

	fmt.Println("\n--- What ++ and -- CANNOT do (see README) ---")
	fmt.Println("The following would NOT compile if uncommented:")
	fmt.Println(`  x := 5`)
	fmt.Println(`  y := x++       // ERROR: ++ is a STATEMENT, not an expression`)
	fmt.Println(`  fmt.Println(x++) // ERROR: same reason — you can't use its "result"`)
	fmt.Println(`  z := ++x        // ERROR: Go has NO pre-increment form at all`)
}
