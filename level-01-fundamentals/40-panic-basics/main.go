// Lesson 40: panic Basics
//
// Goal: Understand panic as a control-flow mechanism — including
// panicking with a CUSTOM value (not just a string), and seeing that
// deferred calls still run while a panic unwinds the stack, even
// before anything actually recovers it.
package main

import "fmt"

// lookupError is a custom panic VALUE type — panic() accepts any(),
// not just strings or errors. Using a structured value like this lets
// whatever eventually recovers it inspect real fields, not just a
// human-readable message.
type lookupError struct {
	Index, Length int
}

func (e lookupError) Error() string {
	return fmt.Sprintf("index %d out of range for length %d", e.Index, e.Length)
}

func riskyLookup(items []string, index int) string {
	if index < 0 || index >= len(items) {
		panic(lookupError{Index: index, Length: len(items)})
	}
	return items[index]
}

// runWithCleanup demonstrates that a DEFERRED call still runs during a
// panic's stack unwinding — even though the function never reaches its
// normal return statement at all.
func runWithCleanup() {
	defer fmt.Println("  runWithCleanup: cleanup ran (deferred, even though we're about to panic)")
	fmt.Println("  runWithCleanup: about to panic")
	panic("something went wrong inside runWithCleanup")
}

func main() {
	fmt.Println("=== panic Basics ===")
	fmt.Println("----------------------------------")

	items := []string{"a", "b", "c"}
	fmt.Printf("riskyLookup(items, 1) = %q (no panic — valid index)\n", riskyLookup(items, 1))

	fmt.Println("\n--- panic with a CUSTOM value type, not just a string ---")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recovered a panic value of type %T: %v\n", r, r)
			}
		}()
		riskyLookup(items, 10)
	}()

	fmt.Println("\n--- Deferred calls run DURING unwinding, before anything recovers ---")
	func() {
		defer func() {
			recover() // just enough to keep this demo running — lesson 41 covers recover properly
		}()
		runWithCleanup()
	}()

	fmt.Println("\nSee the README: an UNRECOVERED panic keeps unwinding all the way up and,")
	fmt.Println("if nothing ever recovers it, crashes the whole program.")
}
