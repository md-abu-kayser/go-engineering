// Lesson 04: package main
//
// Goal: Understand what makes `package main` special, and how a package
// with multiple functions is organized.
package main

import "fmt"

// greet is an ordinary function that belongs to package main, just like
// main() does. A single package can — and usually does — contain many
// functions, types, and variables, not just main().
func greet(name string) string {
	return "Hello, " + name + "!"
}

// farewell is a second ordinary function, to make it obvious that a
// package is just "everything declared in these files", not "just main()".
func farewell(name string) string {
	return "Goodbye, " + name + "."
}

func main() {
	fmt.Println(greet("Gopher"))
	fmt.Println(farewell("Gopher"))

	fmt.Println()
	fmt.Println("This file declares 'package main', which is what tells the Go")
	fmt.Println("toolchain: 'this produces a runnable program', not a library.")
}
