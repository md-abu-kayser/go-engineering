// Lesson 12: go vet
//
// Goal: Understand what `go vet` catches that the compiler does not —
// suspicious code that compiles fine but is probably a bug.
package main

import "fmt"

func main() {
	name := "Gopher"
	age := 15

	// This line is correct: the verb (%s) matches the argument type (string).
	fmt.Printf("%s is %d years old\n", name, age)

	fmt.Println("\nSee the README for an example of the kind of bug `go vet` catches —")
	fmt.Println("this file is intentionally correct so it builds and vets cleanly.")
}
