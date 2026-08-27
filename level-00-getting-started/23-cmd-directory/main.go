// Lesson 23: cmd/ Directory
//
// Goal: Understand the cmd/ convention for modules that produce more than
// one executable.
package main

import "fmt"

func main() {
	fmt.Println("=== cmd/ directory ===")
	fmt.Println("----------------------------------")
	fmt.Println("This lesson's real example lives in ./cmd/greeter — a second,")
	fmt.Println("independent `package main`. Run it with:")
	fmt.Println("  go run ./cmd/greeter")
}
