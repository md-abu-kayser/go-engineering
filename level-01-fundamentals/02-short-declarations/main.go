// Lesson 02: Short Declarations
//
// Goal: Use := correctly, understand its scope rules, and know exactly
// when Go allows "redeclaring" a name with := and when it doesn't.
package main

import "fmt"

func main() {
	// The idiomatic, everyday way to declare a LOCAL variable with an
	// inferred type. This is by far the most common declaration style
	// inside function bodies.
	name := "Gopher"
	age := 15

	fmt.Println("=== Short Declarations ===")
	fmt.Println("----------------------------------")
	fmt.Printf("name: %s, age: %d\n", name, age)

	// := also works for MULTIPLE variables at once — extremely common
	// with functions that return two values (a result and an error).
	city, country := "Dhaka", "Bangladesh"
	fmt.Printf("city: %s, country: %s\n", city, country)

	// A SPECIAL RULE: := can "redeclare" as long as AT LEAST ONE
	// variable on the left is genuinely new. Here, `age` is reused
	// (assigned, not redeclared) while `job` is newly declared.
	age, job := 16, "Student"
	fmt.Printf("age (updated): %d, job (new): %s\n", age, job)

	// := creates a NEW variable in each new scope — this `name` is a
	// DIFFERENT variable from the outer one, only visible inside this
	// block. This is called "shadowing".
	{
		name := "Shadowed Gopher"
		fmt.Printf("inside block, name: %s\n", name)
	}
	fmt.Printf("outside block, name is still: %s\n", name)
}
