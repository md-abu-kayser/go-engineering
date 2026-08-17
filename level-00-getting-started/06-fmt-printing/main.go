// Lesson 06: fmt & Printing
//
// Goal: Understand the fmt package's core printing functions and the most
// commonly used formatting verbs.
package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Gopher"
	age := 15
	height := 1.75
	isAwesome := true

	fmt.Println("=== Print vs Println ===")
	fmt.Print("No newline here...")
	fmt.Print(" ...still the same line.\n")
	fmt.Println("Println adds a newline automatically.")

	fmt.Println("\n=== Printf verbs ===")
	fmt.Printf("%%s (string)              : %s\n", name)
	fmt.Printf("%%d (integer)             : %d\n", age)
	fmt.Printf("%%f (float, default)      : %f\n", height)
	fmt.Printf("%%.2f (float, 2 decimals) : %.2f\n", height)
	fmt.Printf("%%t (boolean)             : %t\n", isAwesome)
	fmt.Printf("%%v (default format)      : %v\n", name)
	fmt.Printf("%%T (type of the value)   : %T\n", age)
	fmt.Printf("%%q (quoted string)       : %q\n", name)

	fmt.Println("\n=== Sprintf: build a string instead of printing it ===")
	summary := fmt.Sprintf("%s is %d years old and %.2fm tall.", name, age, height)
	fmt.Println(summary)

	fmt.Println("\n=== Fprintf: write to a specific destination ===")
	fmt.Fprintf(os.Stdout, "Written explicitly to standard output.\n")
}
