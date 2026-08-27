// Lesson 35: Stack Traces
//
// Goal: Read a Go stack trace confidently — both the one Delve shows you
// on demand, and the one Go itself prints automatically when a program
// panics.
package main

import (
	"fmt"
	"runtime/debug"
)

// factorial calls itself recursively, giving us a genuinely deep call
// stack to inspect — the same shape of problem you'd see debugging real
// recursive or deeply-layered code.
func factorial(n int) int {
	if n <= 1 {
		printStackOnce(n)
		return 1
	}
	return n * factorial(n-1)
}

// printStackOnce prints the current goroutine's stack trace exactly
// once, at the deepest point of the recursion, using the standard
// library instead of a debugger — so you can see the same information
// dlv's `stack` command would show you, without needing dlv running.
var printed bool

func printStackOnce(n int) {
	if printed {
		return
	}
	printed = true
	fmt.Println("=== Stack trace at the deepest recursive call ===")
	debug.PrintStack()
}

func main() {
	result := factorial(5)
	fmt.Printf("\n5! = %d\n", result)
	fmt.Println("\nSee the README for reading this trace, and for the `dlv stack` equivalent.")
}
