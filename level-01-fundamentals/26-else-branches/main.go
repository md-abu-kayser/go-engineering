// Lesson 26: else Branches
//
// Goal: Chain if/else if/else correctly, understand Go's mandatory
// brace-placement rule for `else`, and see why idiomatic Go often
// avoids `else` entirely in favor of early returns.
package main

import "fmt"

// grade demonstrates a classic if/else-if/else chain.
func grade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else {
		return "F"
	}
}

// validateAge shows the IDIOMATIC alternative: early returns, with NO
// else at all. Once a function returns, there's nothing left to be
// "else" about — this avoids nesting entirely.
func validateAge(age int) (ok bool, reason string) {
	if age < 0 {
		return false, "age cannot be negative"
	}
	if age > 150 {
		return false, "age is not plausible"
	}
	// No else needed here — if execution reaches this point, both
	// checks above have already passed.
	return true, ""
}

func main() {
	fmt.Println("=== else Branches ===")
	fmt.Println("----------------------------------")

	for _, score := range []int{95, 82, 71, 40} {
		fmt.Printf("score %3d -> grade %s\n", score, grade(score))
	}

	fmt.Println("\n--- Early return instead of else ---")
	for _, age := range []int{-5, 30, 200} {
		ok, reason := validateAge(age)
		if ok {
			fmt.Printf("age %4d: valid\n", age)
		} else {
			fmt.Printf("age %4d: invalid (%s)\n", age, reason)
		}
	}
}
