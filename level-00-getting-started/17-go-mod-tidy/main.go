// Lesson 17: go mod tidy
//
// Goal: Understand what `go mod tidy` does to keep go.mod and go.sum
// accurate — even though this particular lesson has no external
// dependencies to demonstrate it with directly.
package main

import "fmt"

func main() {
	fmt.Println("=== go mod tidy ===")
	fmt.Println("----------------------------------")
	fmt.Println("This module currently has zero external dependencies, so `go mod tidy`")
	fmt.Println("has nothing to add or remove here. See the README for what it does on a")
	fmt.Println("module that *does* depend on third-party packages.")
}
