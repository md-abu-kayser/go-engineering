// Lesson 49: Closures
//
// Goal: Write a CLOSURE — a function literal that captures a variable
// from its surrounding scope — and use the classic "counter generator"
// pattern. Also confirm Go's (modern, 1.22+) per-iteration loop
// variable behavior directly, since it specifically affects closures
// created inside a loop.
package main

import "fmt"

// makeCounter returns a NEW function each time it's called — and that
// returned function CLOSES OVER `count`, meaning it keeps its own
// private reference to that specific variable, even after makeCounter
// itself has returned and count would normally have gone out of scope.
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	fmt.Println("=== Closures ===")
	fmt.Println("----------------------------------")

	// Each call to makeCounter() creates a BRAND NEW `count` variable —
	// the two counters below are completely independent of each other.
	counterA := makeCounter()
	counterB := makeCounter()

	fmt.Printf("counterA(): %d\n", counterA())
	fmt.Printf("counterA(): %d\n", counterA())
	fmt.Printf("counterA(): %d\n", counterA())
	fmt.Printf("counterB(): %d (independent of counterA — its OWN count)\n", counterB())

	// A closure over a variable declared in main() itself — the
	// returned function keeps working with THIS SPECIFIC variable.
	fmt.Println("\n--- Closing over an outer variable directly ---")
	total := 0
	addToTotal := func(n int) {
		total += n // modifies the OUTER `total`, not a copy of it
	}
	addToTotal(5)
	addToTotal(10)
	fmt.Printf("total = %d (modified via the closure, twice)\n", total)

	// Go 1.22+ loop variable semantics: EACH iteration gets its OWN
	// copy of the loop variable, so closures created inside a loop
	// correctly capture that specific iteration's value — this was
	// NOT always true in Go (versions before 1.22 shared ONE variable
	// across all iterations, a well-known historical gotcha).
	fmt.Println("\n--- Closures inside a loop (Go 1.22+ per-iteration variables) ---")
	var funcs []func()
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Printf("  captured i = %d\n", i)
		})
	}
	for _, f := range funcs {
		f()
	}
}
