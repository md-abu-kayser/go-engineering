// Lesson 34: break
//
// Goal: Use `break` to exit a loop early, and see the classic gotcha:
// `break` inside a `switch` that's nested inside a loop only exits the
// SWITCH, not the loop — it targets its NEAREST enclosing construct.
package main

import "fmt"

func main() {
	fmt.Println("=== break ===")
	fmt.Println("----------------------------------")

	// The straightforward case: break exits the loop immediately.
	fmt.Println("--- break in a plain loop ---")
	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}
		fmt.Printf("  i = %d\n", i)
	}

	// THE GOTCHA: break inside a switch, which is itself inside a
	// loop, only breaks out of the SWITCH — the loop keeps going.
	// This surprises people coming from languages where a bare `break`
	// always means "exit the nearest loop", full stop.
	fmt.Println("\n--- The gotcha: break inside switch-inside-loop ---")
	for i := 0; i < 5; i++ {
		switch {
		case i == 2:
			fmt.Printf("  i=%d: breaking the SWITCH, not the loop\n", i)
			break // exits the SWITCH only — the for loop is UNAFFECTED
		default:
			fmt.Printf("  i=%d: default case\n", i)
		}
		fmt.Printf("  (loop continues after switch, i=%d)\n", i)
	}

	fmt.Println("\nSee lesson 36 (labeled break) for how to ACTUALLY break the outer loop")
	fmt.Println("from inside a nested switch or inner loop.")
}
