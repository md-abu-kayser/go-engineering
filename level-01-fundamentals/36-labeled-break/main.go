// Lesson 36: Labeled break
//
// Goal: Use a LABELED break to exit an OUTER loop from inside a nested
// switch or inner loop — solving exactly the gotcha demonstrated in
// lesson 34.
package main

import "fmt"

func main() {
	fmt.Println("=== Labeled break ===")
	fmt.Println("----------------------------------")

	// Recall lesson 34's problem: a plain `break` inside a switch that's
	// inside a loop only exits the SWITCH. A LABEL, placed directly
	// before the loop, gives `break` something more specific to target.
	fmt.Println("--- Fixing lesson 34's gotcha with a label ---")
outer: // this label refers to the FOR loop immediately below it
	for i := 0; i < 5; i++ {
		switch {
		case i == 2:
			fmt.Printf("  i=%d: breaking the OUTER LOOP this time, via the label\n", i)
			break outer // exits the LOOP, not just the switch — this is the fix
		default:
			fmt.Printf("  i=%d: default case\n", i)
		}
		fmt.Printf("  (this line will NOT print for i=2 — the loop already exited)\n")
	}
	fmt.Println("  loop fully exited")

	// The same label mechanism ALSO works to break out of multiple
	// NESTED loops at once — not just a switch.
	fmt.Println("\n--- Breaking out of nested loops entirely ---")
search:
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if row == 1 && col == 1 {
				fmt.Printf("  found target at row=%d, col=%d — stopping BOTH loops\n", row, col)
				break search // exits BOTH loops immediately, not just the inner one
			}
			fmt.Printf("  checking row=%d, col=%d\n", row, col)
		}
	}
	fmt.Println("  search finished")
}
