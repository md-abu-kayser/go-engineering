// Lesson 15: go list
//
// Goal: Understand `go list` — the command that answers questions like
// "what packages does this module contain?" and "what does this package
// import?" This lesson is command-line-focused; the code here is a small,
// realistic package to run `go list` against.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== go list ===")
	fmt.Println("----------------------------------")
	fmt.Println("This program itself isn't the point of this lesson — try the")
	fmt.Println("commands in the README against this repository instead.")

	// A tiny bit of real logic, so `go list` has genuine imports to report on.
	fmt.Println(strings.Join(os.Args, " "))
}
