// Lesson 48: Cross-Compilation Overview
//
// Goal: See what a running binary can tell you about the platform it
// was built for, and get an overview of the full GOOS/GOARCH matrix Go
// can target — building on the brief mention in lesson 08.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=== Cross-Compilation Overview ===")
	fmt.Println("----------------------------------")
	fmt.Printf("This binary targets: GOOS=%s GOARCH=%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println()
	fmt.Println("See the README for the full GOOS/GOARCH matrix, and how to list it yourself")
	fmt.Println("with `go tool dist list`.")
}
