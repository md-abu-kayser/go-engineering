// Lesson 05: Typed Constants
//
// Goal: Distinguish TYPED constants from UNTYPED constants, and see how
// untyped constants get a "default type" only when they finally need one
// — letting the same literal flexibly work as different numeric types.
package main

import "fmt"

// An UNTYPED constant — just a bare literal. It doesn't commit to a
// specific type until it's actually used somewhere that needs one.
const untypedNumber = 100

// A TYPED constant — explicitly pinned to float64, and ONLY usable
// as a float64 from here on.
const typedNumber float64 = 100

func main() {
	fmt.Println("=== Typed Constants ===")
	fmt.Println("----------------------------------")

	// untypedNumber flexibly becomes whatever numeric type context
	// requires — no explicit conversion needed.
	var asInt int = untypedNumber
	var asFloat float64 = untypedNumber
	var asInt64 int64 = untypedNumber

	fmt.Printf("untypedNumber as int     : %d (%T)\n", asInt, asInt)
	fmt.Printf("untypedNumber as float64 : %v (%T)\n", asFloat, asFloat)
	fmt.Printf("untypedNumber as int64   : %d (%T)\n", asInt64, asInt64)

	// typedNumber is ALREADY float64 — assigning it to an int variable
	// requires an explicit conversion, exactly like any other float64.
	var asFloatDirect float64 = typedNumber
	var asIntConverted int = int(typedNumber) // explicit conversion required
	fmt.Printf("typedNumber as float64   : %v (%T)\n", asFloatDirect, asFloatDirect)
	fmt.Printf("typedNumber as int       : %d (%T, needed int(...))\n", asIntConverted, asIntConverted)

	// An untyped FLOATING constant used somewhere expecting int is fine
	// too, AS LONG AS it has no fractional part.
	const wholeFloat = 4.0
	var asIntFromFloat int = wholeFloat
	fmt.Printf("wholeFloat (4.0) as int  : %d\n", asIntFromFloat)
}
