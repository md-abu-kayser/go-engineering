// Lesson 04: Constants
//
// Goal: Declare constants with `const`, understand how they differ from
// `var` (immutable, must be computable at compile time), and use a
// const block for related values.
package main

import "fmt"

// Constants can live at package level, just like var — and often make
// MORE sense there, since a constant's whole point is that it never
// changes for the lifetime of the program.
const appVersion = "v1.0.0"

const (
	maxRetries   = 3
	timeoutSecs  = 30
	defaultDebug = false
)

func main() {
	const greeting = "Hello"

	fmt.Println("=== Constants ===")
	fmt.Println("----------------------------------")
	fmt.Printf("appVersion   : %s\n", appVersion)
	fmt.Printf("maxRetries   : %d\n", maxRetries)
	fmt.Printf("timeoutSecs  : %d\n", timeoutSecs)
	fmt.Printf("defaultDebug : %t\n", defaultDebug)
	fmt.Printf("greeting     : %s\n", greeting)

	// Constants must be computable AT COMPILE TIME — this expression is
	// fine, because both operands are themselves constants.
	const totalTimeout = timeoutSecs * maxRetries
	fmt.Printf("totalTimeout : %d (computed from two other constants)\n", totalTimeout)
}
