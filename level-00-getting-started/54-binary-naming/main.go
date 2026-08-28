// Lesson 54: Binary Naming
//
// Goal: Understand where a built binary's default name comes from, how
// to override it, and the naming convention used for cross-compiled
// release binaries.
package main

import "fmt"

func main() {
	fmt.Println("=== Binary Naming ===")
	fmt.Println("----------------------------------")
	fmt.Println("This folder is named '54-binary-naming' — that's exactly what")
	fmt.Println("`go build` (with no -o flag) will call the resulting binary.")
	fmt.Println()
	fmt.Println("See the README for overriding it with -o, and for the")
	fmt.Println("<name>-<os>-<arch> convention used for released binaries.")
}
