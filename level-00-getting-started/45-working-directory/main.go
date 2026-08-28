// Lesson 45: Working Directory
//
// Goal: Understand what a process's "current working directory" is,
// read it with os.Getwd, and see how it affects relative paths.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Working Directory ===")
	fmt.Println("----------------------------------")

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error getting working directory:", err)
		os.Exit(1)
	}
	fmt.Printf("Current working directory: %s\n", wd)

	// A relative path is always resolved AGAINST the working directory
	// above — this same "README.md" would resolve completely differently
	// if this program were run from a different directory.
	if _, err := os.Stat("README.md"); err == nil {
		fmt.Println(`"README.md" found relative to the working directory above.`)
	} else {
		fmt.Println(`"README.md" NOT found relative to the working directory above.`)
		fmt.Println("(Try running this from inside the lesson folder itself.)")
	}
}
