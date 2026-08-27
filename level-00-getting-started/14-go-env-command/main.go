// Lesson 14: go env
//
// Goal: Understand the difference between reading a raw shell environment
// variable and asking Go for its *effective* configuration via `go env`.
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("=== From inside the running program ===")
	fmt.Println("----------------------------------")
	fmt.Printf("runtime.GOROOT() : %s\n", runtime.GOROOT())
	fmt.Printf("os.Getenv(\"GOPATH\") : %q (raw shell value — may be empty even if Go has a default)\n", os.Getenv("GOPATH"))
	fmt.Printf("os.Getenv(\"GOOS\")   : %q (raw shell value — usually empty; GOOS is normally implicit)\n", os.Getenv("GOOS"))

	fmt.Println("\n=== Compare against the command line ===")
	fmt.Println("Run these in your terminal and compare:")
	fmt.Println("  go env GOROOT")
	fmt.Println("  go env GOPATH")
	fmt.Println("  go env GOOS GOARCH")
}
