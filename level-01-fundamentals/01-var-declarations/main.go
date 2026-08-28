// Lesson 01: var Declarations
//
// Goal: Use the `var` keyword in every form it supports — single,
// grouped, with an explicit type, and with an inferred type from an
// initial value.
package main

import "fmt"

// var works at package level too, outside any function — something
// := (lesson 02) cannot do.
var appName string = "GO-ENGINEERING"

func main() {
	// Form 1: explicit type, explicit value.
	var age int = 15

	// Form 2: type INFERRED from the value — still uses `var`, just
	// without writing the type yourself.
	var city = "Dhaka"

	// Form 3: declared with no initial value at all — gets its
	// type's ZERO VALUE automatically (see lesson 03 for the full story).
	var score int

	// Form 4: a grouped var block — idiomatic for declaring several
	// related variables together.
	var (
		width  = 1920
		height = 1080
		title  = "Monitor"
	)

	fmt.Println("=== var Declarations ===")
	fmt.Println("----------------------------------")
	fmt.Printf("appName (package-level) : %s\n", appName)
	fmt.Printf("age (explicit type)     : %d\n", age)
	fmt.Printf("city (inferred type)    : %s\n", city)
	fmt.Printf("score (no initializer)  : %d\n", score)
	fmt.Printf("grouped: %s is %dx%d\n", title, width, height)
}
