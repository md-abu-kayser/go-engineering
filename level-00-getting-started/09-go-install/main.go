// Lesson 09: go install
//
// Goal: Understand `go install`, GOBIN, and how Go CLI tools get
// distributed and installed by version.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== go install ===")
	fmt.Println("----------------------------------")

	gobin := os.Getenv("GOBIN")
	if gobin == "" {
		gobin = "(not set — falls back to $GOPATH/bin, run `go env GOBIN` and `go env GOPATH`)"
	}
	fmt.Printf("GOBIN: %s\n", gobin)

	fmt.Println()
	fmt.Println("Run `go install` in this folder, then check the path above for a new binary.")
}
