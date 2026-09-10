// Lesson 58: Package Initialization
//
// Goal: See the EXACT order Go initializes a package in: package-level
// variables first (in dependency order), then every init() function
// (in the order they appear), and ALL of that before main() ever runs.
package main

import "fmt"

// Package-level variables are initialized BEFORE main() runs, and
// BEFORE any init() function runs. Go figures out the correct order
// automatically based on DEPENDENCIES between them — notice `total`
// depends on `base` and `multiplier`, so Go initializes those two
// first, regardless of the order they're WRITTEN in below.
var total = base * multiplier
var base = 10
var multiplier = 3

// A package can have MULTIPLE init() functions — even several in the
// SAME file. Every one of them runs automatically, with no explicit
// call anywhere, before main() starts.
func init() {
	fmt.Println("init() #1: package-level vars are ALREADY initialized here:")
	fmt.Printf("  total = %d (computed from base=%d, multiplier=%d)\n", total, base, multiplier)
}

func init() {
	fmt.Println("init() #2: multiple init() functions run in the order they APPEAR in the file")
}

func main() {
	fmt.Println("\n=== Package Initialization ===")
	fmt.Println("----------------------------------")
	fmt.Println("main() runs AFTER every package-level var AND every init() function above.")
	fmt.Printf("total is still %d here in main() — nothing re-runs it.\n", total)
}
