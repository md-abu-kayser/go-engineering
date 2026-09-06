// Lesson 45: Named Results
//
// Goal: Name return values in a function's signature — documenting
// what each one MEANS — and use the "naked return" they enable. Also
// see the genuine gotcha: a deferred function can still modify a named
// result AFTER a naked return has "returned" it.
package main

import "fmt"

// divide's return values are NAMED: result and err. This documents
// what each one means directly in the signature — genuinely useful
// even before considering the naked-return feature below.
func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("cannot divide %d by zero", a)
		return // a NAKED return — sends back result and err AS THEY CURRENTLY STAND
	}
	result = a / b
	return // naked return again — result and err are already set correctly
}

// rectangleStats shows named results purely for DOCUMENTATION value —
// `go doc` and any editor's signature hint will show "area" and
// "perimeter" by name, which is far more informative than two bare ints.
func rectangleStats(width, height int) (area, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return area, perimeter // an EXPLICIT return also still works, naming or not
}

// safeDivide demonstrates the genuine gotcha: a DEFERRED function can
// still modify a named result AFTER a naked (or even explicit) return
// has set it — because the return doesn't ACTUALLY exit until every
// deferred call has run.
func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = -1 // this OVERWRITES whatever `return` already set, since
			// the deferred function runs AFTER return but BEFORE the
			// function truly exits.
		}
	}()
	result = a / b // panics here if b == 0 (integer divide by zero)
	return
}

func main() {
	fmt.Println("=== Named Results ===")
	fmt.Println("----------------------------------")

	result, err := divide(10, 2)
	fmt.Printf("divide(10, 2) = %d, err = %v\n", result, err)

	result, err = divide(10, 0)
	fmt.Printf("divide(10, 0) = %d, err = %v\n", result, err)

	fmt.Println("\n--- Named results as documentation ---")
	area, perimeter := rectangleStats(4, 6)
	fmt.Printf("rectangleStats(4, 6) = area %d, perimeter %d\n", area, perimeter)

	fmt.Println("\n--- The defer + named result gotcha ---")
	fmt.Printf("safeDivide(10, 2) = %d (normal case)\n", safeDivide(10, 2))
	fmt.Printf("safeDivide(10, 0) = %d (deferred recover OVERWROTE the named result)\n", safeDivide(10, 0))
}
