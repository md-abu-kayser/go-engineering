// Lesson 42: Functions
//
// Goal: Declare and call functions with Go's exact syntax — name,
// parameters (each with an explicit type), and a return type — the
// foundation everything in the rest of this module builds on.
package main

import "fmt"

// add takes two named parameters, both int, and returns a single int.
// The syntax is: func name(param type, param type) returnType { ... }
func add(a int, b int) int {
	return a + b
}

// When consecutive parameters share the same type, you can write the
// type ONCE at the end — `a, b int` means both a and b are int. This
// is purely a shorthand; it behaves identically to `a int, b int`.
func multiply(a, b int) int {
	return a * b
}

// A function with NO return value simply omits the return type
// entirely — there's no `void` keyword in Go.
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// A function with NO parameters just has empty parentheses.
func currentVersion() string {
	return "v1.0.0"
}

func main() {
	fmt.Println("=== Functions ===")
	fmt.Println("----------------------------------")

	fmt.Printf("add(3, 4)      = %d\n", add(3, 4))
	fmt.Printf("multiply(3, 4) = %d\n", multiply(3, 4))

	greet("Gopher")

	fmt.Printf("currentVersion() = %s\n", currentVersion())
}
