// Lesson 53: Reproducible Build Basics
//
// Goal: Understand what makes a build "reproducible" (bit-for-bit
// identical given the same inputs), and the specific flag that removes
// the most common source of unwanted variation: local file paths.
package main

import "fmt"

func main() {
	fmt.Println("=== Reproducible Build Basics ===")
	fmt.Println("----------------------------------")
	fmt.Println("This program's own logic doesn't change between builds — but by")
	fmt.Println("default, Go embeds your LOCAL FILE PATH into the compiled binary.")
	fmt.Println()
	fmt.Println("See the README for how -trimpath removes that, and why it matters")
	fmt.Println("for producing byte-identical binaries from identical source.")
}
