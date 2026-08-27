// Lesson 39: Exit Status
//
// Goal: Understand how a Go program communicates success or failure to
// whatever ran it — the shell, a CI pipeline, another process — via its
// exit status.
package main

import (
	"fmt"
	"os"
)

// validateAge returns an error if age is out of a sane range.
func validateAge(age int) error {
	if age < 0 || age > 150 {
		return fmt.Errorf("age %d is out of a plausible range", age)
	}
	return nil
}

func main() {
	age := 15 // hardcoded for this lesson's demonstration

	if err := validateAge(age); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Age %d is valid.\n", age)
	fmt.Println("\nSee the README for how to check this program's exit status from your shell,")
	fmt.Println("and how it would differ if validateAge had failed instead.")
	// Falling off the end of main() is equivalent to os.Exit(0) — success.
}
