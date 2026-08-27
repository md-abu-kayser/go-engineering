// Lesson 22: internal/ Packages
//
// Goal: Use a real internal package, and understand exactly what the Go
// compiler enforces about where it can be imported from.
package main

import (
	"fmt"

	"go-engineering/level-00-getting-started/22-internal-packages/internal/greeting"
)

func main() {
	fmt.Println(greeting.Hello("Gopher"))
	fmt.Println()
	fmt.Println("Try importing 'internal/greeting' from a DIFFERENT lesson folder —")
	fmt.Println("the compiler refuses. See the README for why.")
}
