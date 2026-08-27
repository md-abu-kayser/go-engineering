// Lesson 26: examples/ Directory
//
// Goal: Understand why libraries ship a runnable examples/ directory,
// distinct from both their tests and their own source.
package main

import "fmt"

func main() {
	fmt.Println("=== examples/ directory ===")
	fmt.Println("----------------------------------")
	fmt.Println("This lesson's real example lives in ./examples/basic — a small,")
	fmt.Println("standalone program showing how to use ./greetlib. Run it with:")
	fmt.Println("  go run ./examples/basic")
}
