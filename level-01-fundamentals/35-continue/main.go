// Lesson 35: continue
//
// Goal: Use `continue` to skip to the next iteration, and see that in
// NESTED loops, `continue` always targets the INNERMOST enclosing loop
// — never an outer one, by default.
package main

import "fmt"

func main() {
	fmt.Println("=== continue ===")
	fmt.Println("----------------------------------")

	// The straightforward case, already seen in lesson 30: skip one
	// iteration's remaining body, keep looping.
	fmt.Println("--- continue in a plain loop ---")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Printf("  i = %d (odd)\n", i)
	}

	// NESTED loops: continue only affects the INNERMOST loop it's
	// written in — it has no effect on any OUTER loop at all.
	fmt.Println("\n--- continue in nested loops (targets the INNER loop only) ---")
	for row := 0; row < 3; row++ {
		fmt.Printf("row %d: ", row)
		for col := 0; col < 4; col++ {
			if col == 1 {
				continue // skips just THIS inner iteration; the outer `row` loop is unaffected
			}
			fmt.Printf("%d ", col)
		}
		fmt.Println() // this line runs once per ROW, proving the outer loop's structure is intact
	}

	fmt.Println("\nSee lesson 37 (labeled continue) for how to continue an OUTER loop")
	fmt.Println("from inside a nested inner loop.")
}
