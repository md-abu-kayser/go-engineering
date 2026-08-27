// Lesson 21: Workspace Layout
//
// Goal: Understand what a Go workspace (go.work) is for — developing
// across multiple modules at once — as distinct from a single module's
// internal folder layout.
package main

import "fmt"

func main() {
	fmt.Println("=== Workspace Layout (go.work) ===")
	fmt.Println("----------------------------------")
	fmt.Println("This repository is a single module, so it has no go.work file.")
	fmt.Println("See the README for when multi-module workspaces are worth using.")
}
