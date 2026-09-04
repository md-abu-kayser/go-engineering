// Lesson 41: recover Basics
//
// Goal: Use `recover` correctly — including the precise, easy-to-miss
// rule that recover() only catches an in-progress panic when called
// DIRECTLY inside a deferred function, not one level removed through a
// helper function call.
package main

import "fmt"

// safeCall runs fn, converting any panic into a returned error instead
// of letting it crash the caller. This is the general shape behind
// almost every real-world defer+recover usage in Go.
func safeCall(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	fn()
	return nil
}

// recoverHelper calls recover(), but recoverHelper ITSELF is not the
// deferred function — it's a regular function CALLED BY one. This
// distinction is exactly what the rule below hinges on.
func recoverHelper() any {
	return recover()
}

func demonstrateRecoverMustBeDirect() {
	defer func() {
		// recover() called INDIRECTLY, through a helper — this does
		// NOT catch the panic below, even though one is genuinely
		// in progress right now.
		r := recoverHelper()
		fmt.Printf("recoverHelper() during a real panic: %v (nil — did NOT catch it)\n", r)

		// recover() called DIRECTLY, right here in the deferred
		// function itself — THIS one actually catches it.
		real := recover()
		fmt.Printf("recover() called DIRECTLY in this deferred func: %v (this ONE works)\n", real)
	}()
	panic("a panic used to demonstrate recover's placement rule")
}

func main() {
	fmt.Println("=== recover Basics ===")
	fmt.Println("----------------------------------")

	err := safeCall(func() {
		fmt.Println("  running safely, no panic here")
	})
	fmt.Printf("safeCall (no panic)  -> err = %v\n", err)

	err = safeCall(func() {
		panic("deliberate panic for this demo")
	})
	fmt.Printf("safeCall (panicking) -> err = %v\n", err)

	// recover() called with NO panic in progress simply returns nil —
	// calling it "just in case" is always safe, never an error itself.
	fmt.Println("\n--- recover() with no panic in progress ---")
	r := recover()
	fmt.Printf("recover() called with nothing panicking: %v\n", r)

	// THE precise placement rule: recover() only has an effect when
	// called DIRECTLY inside a deferred function — one level of
	// indirection through a helper function defeats it completely.
	fmt.Println("\n--- recover() must be called DIRECTLY inside the deferred function ---")
	demonstrateRecoverMustBeDirect()
	fmt.Println("\nProgram continued normally — the direct recover() call caught the panic.")
}
