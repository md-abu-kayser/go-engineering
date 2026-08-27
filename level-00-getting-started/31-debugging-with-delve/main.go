// Lesson 31: Debugging with Delve
//
// Goal: Use Delve (dlv) directly from the command line — the same
// debugger VS Code's "Run and Debug" panel drives for you under the
// hood in lesson 29 — to inspect a running Go program step by step.
package main

import "fmt"

// applyDiscount reduces priceCents by percentOff percent, rounding down.
// It has a subtle bug on purpose: see the README's debugging walkthrough
// for how to find it with dlv instead of just reading the code.
func applyDiscount(priceCents, percentOff int) int {
	discount := priceCents * percentOff / 100
	return priceCents - discount
}

func main() {
	cart := []int{1999, 500, 1250}
	total := 0

	for _, price := range cart {
		discounted := applyDiscount(price, 10)
		fmt.Printf("price: %d -> after 10%% off: %d\n", price, discounted)
		total += discounted
	}

	fmt.Printf("Total: %d cents\n", total)
	fmt.Println("\nSee the README for a step-by-step dlv session on this exact program.")
}
