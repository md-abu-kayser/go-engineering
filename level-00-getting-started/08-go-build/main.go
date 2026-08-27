// Lesson 08: go build
//
// Goal: Go deeper into `go build` — flags, output naming, and
// cross-compilation — beyond the basic use already seen in lesson 07.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=== go build, in depth ===")
	fmt.Println("----------------------------------")
	fmt.Printf("This binary was compiled for: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println()
	fmt.Println("Try cross-compiling this file for another platform (see README).")
}
