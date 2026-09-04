// Lesson 37: Labeled continue
//
// Goal: Use a LABELED continue to skip to the next iteration of an
// OUTER loop from inside a nested inner loop — solving exactly the
// scoping limitation demonstrated in lesson 35.
package main

import "fmt"

func main() {
	fmt.Println("=== Labeled continue ===")
	fmt.Println("----------------------------------")

	// Recall lesson 35: a plain `continue` inside a nested inner loop
	// only affects that inner loop. A label lets us target the OUTER
	// loop's next iteration instead, abandoning the rest of the inner
	// loop entirely.
	fmt.Println("--- Skipping an entire outer iteration from an inner loop ---")
rows:
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if col == 2 && row%2 == 0 {
				fmt.Printf("  row=%d, col=%d: skipping the REST of this row entirely\n", row, col)
				continue rows // jumps straight to the OUTER loop's next iteration
			}
			fmt.Printf("  row=%d, col=%d\n", row, col)
		}
		fmt.Printf("  (finished row %d normally)\n", row)
	}

	fmt.Println("\n--- Contrast: a PLAIN continue only skips within the inner loop ---")
	for row := 0; row < 2; row++ {
		for col := 0; col < 4; col++ {
			if col == 2 {
				continue // plain continue — only affects THIS inner loop
			}
			fmt.Printf("  row=%d, col=%d\n", row, col)
		}
		fmt.Printf("  (finished row %d — notice col=3 STILL printed, unlike the labeled version above)\n", row)
	}
}
