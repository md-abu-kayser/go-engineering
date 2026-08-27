// Lesson 11: go fmt
//
// Goal: Understand what `gofmt` / `go fmt` actually normalizes, and why Go
// treats formatting as a solved, non-debatable problem.
package main

import "fmt"

// Point is intentionally formatted exactly as gofmt would produce it:
// tabs for indentation, aligned struct fields, one blank line of
// breathing room. Try messing up the spacing below and running
// `gofmt -l .` to see it get flagged.
type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 3, Y: 4}
	fmt.Printf("Point{X: %d, Y: %d}\n", p.X, p.Y)
	fmt.Println("\nThis file is already gofmt-clean. See the README for how to prove it.")
}
