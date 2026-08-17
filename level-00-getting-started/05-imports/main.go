// Lesson 05: Imports
//
// Goal: Understand Go's import styles — grouped, aliased, and blank —
// and why unused imports are a compile error.
package main

import (
	"fmt"           // a plain, standard import
	m "math"        // an aliased import: refer to it as `m` instead of `math`
	_ "time/tzdata" // a blank import: kept only for its side effect (see below)
)

func main() {
	fmt.Println("=== Import Styles ===")
	fmt.Println("----------------------------------")

	// Because "math" was imported as `m`, we call it via `m.` instead of `math.`
	fmt.Printf("Pi (via aliased import `m \"math\"`) : %.5f\n", m.Pi)
	fmt.Printf("Square root of 2 : %.5f\n", m.Sqrt(2))

	fmt.Println()
	fmt.Println("The blank import `_ \"time/tzdata\"` above contributes nothing we call")
	fmt.Println("directly. It runs only for its side effect: it embeds the IANA time zone")
	fmt.Println("database into this binary, so time.LoadLocation() works even on a machine")
	fmt.Println("that has no system time zone data installed.")
}
