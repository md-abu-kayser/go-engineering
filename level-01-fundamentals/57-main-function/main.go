// Lesson 57: The main Function
//
// Goal: Understand exactly what's special about func main() itself —
// no parameters, no return value, called automatically by the Go
// runtime rather than by any code you write, and unique per program.
package main

import (
	"fmt"
	"os"
)

// setup is an ORDINARY function main() calls — main() itself does the
// calling; nothing calls main() from within this program's own code.
func setup() {
	fmt.Println("setup() ran — called explicitly, by main() itself")
}

// func main() takes NO parameters and returns NOTHING — this exact
// signature is mandatory; Go will not compile a "func main(args
// []string) int" or any other variation, no matter how reasonable it
// might look coming from another language.
func main() {
	fmt.Println("=== The main Function ===")
	fmt.Println("----------------------------------")

	fmt.Println("main() itself was called by the Go RUNTIME, not by any")
	fmt.Println("line of code in this program — it's the program's entry point.")

	setup()

	// os.Args (level 00, lesson 44) is how main() ACCESSES command-line
	// arguments — since main() itself can't declare parameters, this is
	// the mechanism Go provides instead.
	fmt.Printf("\nos.Args (main() can't take parameters, so this is the workaround): %v\n", os.Args)

	// The program's lifetime is bounded by main(): the moment main()
	// returns (normally, or via os.Exit), the whole program ends —
	// even if OTHER goroutines were still running (a later-level topic).
	fmt.Println("\nOnce main() returns here, the entire program ends.")
}
