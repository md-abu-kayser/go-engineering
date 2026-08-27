// Lesson 16: go mod init
//
// Goal: Understand exactly what `go mod init` creates and why the module
// path you choose matters.
package main

import "fmt"

func main() {
	fmt.Println("=== go mod init ===")
	fmt.Println("----------------------------------")
	fmt.Println("This repository's go.mod was created with:")
	fmt.Println("  go mod init go-engineering")
	fmt.Println()
	fmt.Println("See the README for what that command generates and why the module")
	fmt.Println("path matters beyond this repo.")
}
