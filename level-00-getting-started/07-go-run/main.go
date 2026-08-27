// Lesson 07: go run
//
// Goal: See, from inside the program itself, how it was invoked — and use
// that to understand the difference between `go run`, `go build`, and
// `go install`.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== go run vs go build vs go install ===")
	fmt.Println("----------------------------------")
	fmt.Printf("This process was started as: %s\n", os.Args[0])
	fmt.Println()
	fmt.Println("Try running this file three different ways and compare the line above:")
	fmt.Println("  1. go run main.go")
	fmt.Println("  2. go build && ./07-go-run        (07-go-run.exe on Windows)")
	fmt.Println("  3. go install && 07-go-run         (once $GOBIN is on your PATH)")
}
