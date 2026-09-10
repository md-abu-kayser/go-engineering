// Lesson 55: Unexported Identifiers
//
// Goal: Use the "New + methods" pattern to work with a type whose
// FIELDS are entirely unexported — confirming that unexported
// identifiers genuinely cannot be reached from another package, full
// stop, not even to READ them.
package main

import (
	"fmt"

	"go-engineering/level-01-fundamentals/55-unexported-identifiers/counter"
)

func main() {
	fmt.Println("=== Unexported Identifiers ===")
	fmt.Println("----------------------------------")

	// We can ONLY create a Counter via the exported New() constructor —
	// counter.Counter{value: 0, step: 5} would NOT compile here, since
	// `value` and `step` are unexported and this file is in package main,
	// not package counter.
	c := counter.New(5)

	c.Increment()
	c.Increment()
	c.Increment()

	// We can only READ the current count through the exported Value()
	// method — there is no way to write `c.value` directly from here.
	fmt.Printf("After 3 increments of step 5: %d\n", c.Value())

	fmt.Println("\nSee the README for the exact compiler error you'd get trying to")
	fmt.Println("access counter.Counter's fields directly from this package.")
}
