// Lesson 36: Compile Errors
//
// Goal: Read a Go compiler error message confidently — file, line, column,
// and the specific complaint — using real (correct) code as a baseline
// and a catalog of common mistakes in the README.
package main

import "fmt"

// safeDivide returns a/b, and an error instead of panicking if b is zero.
// It exists here as valid, compiling code to contrast against the broken
// snippets discussed in the README.
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

func main() {
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("10 / 2 =", result)

	fmt.Println("\nThis file compiles cleanly. See the README for the exact compiler")
	fmt.Println("errors you'd get from several common, deliberate mistakes.")
}
