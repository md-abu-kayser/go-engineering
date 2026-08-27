// Lesson 41: Standard Output
//
// Goal: Understand os.Stdout as the destination behind fmt.Println, and
// practice redirecting it from the shell.
package main

import (
	"fmt"
	"os"
)

func main() {
	// These three lines are equivalent in effect — fmt.Println always
	// writes to os.Stdout; Fprintln just makes that destination explicit.
	fmt.Println("Written via fmt.Println (implicitly to stdout)")
	fmt.Fprintln(os.Stdout, "Written via fmt.Fprintln(os.Stdout, ...) (explicitly)")

	os.Stdout.WriteString("Written directly via os.Stdout.WriteString\n")

	fmt.Println("\nSee the README for redirecting this program's stdout to a file,")
	fmt.Println("and for why that's different from redirecting stderr (lesson 42).")
}
