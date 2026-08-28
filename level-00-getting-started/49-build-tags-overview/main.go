// Lesson 49: Build Tags Overview
//
// Goal: See build constraints ("build tags") actually change which code
// gets compiled — both an OS-based constraint (windows vs everything
// else) and a custom, hand-defined one ("debug").
package main

import "fmt"

func main() {
	fmt.Println("=== Build Tags Overview ===")
	fmt.Println("----------------------------------")
	fmt.Printf("platformName() : %s\n", platformName())
	fmt.Printf("debugEnabled   : %t\n", debugEnabled)
	fmt.Println()
	fmt.Println("Try: go run -tags debug . -- and compare debugEnabled above.")
}
