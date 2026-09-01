// Lesson 27: switch Statements
//
// Goal: Use Go's `switch` — which, unlike C/Java/JavaScript, does NOT
// fall through to the next case by default, and supports multiple
// values per case directly.
package main

import "fmt"

func dayType(day string) string {
	switch day {
	case "Saturday", "Sunday": // MULTIPLE values in one case, comma-separated
		return "Weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "Weekday"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println("=== switch Statements ===")
	fmt.Println("----------------------------------")

	for _, day := range []string{"Saturday", "Monday", "Someday"} {
		fmt.Printf("%-10s -> %s\n", day, dayType(day))
	}

	// NO fallthrough by default — each case is complete on its own.
	// This is the single biggest difference from C/Java/JavaScript's
	// switch, where forgetting `break` causes execution to "fall
	// through" into the next case unintentionally.
	fmt.Println("\n--- No fallthrough by default ---")
	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two") // execution STOPS here — does NOT continue into case 3
	case 3:
		fmt.Println("three")
	}

	// The explicit `fallthrough` keyword exists for the RARE cases you
	// genuinely want C-style fall-through behavior — it's opt-in, not
	// the default.
	fmt.Println("\n--- Explicit fallthrough (opt-in, rare) ---")
	switch n {
	case 2:
		fmt.Println("two")
		fallthrough // explicitly continue into the NEXT case, unconditionally
	case 3:
		fmt.Println("three (reached via explicit fallthrough, not because n == 3)")
	case 4:
		fmt.Println("four (NOT reached — fallthrough only continues ONE case)")
	}
}
