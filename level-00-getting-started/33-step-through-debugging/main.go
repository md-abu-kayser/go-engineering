// Lesson 33: Step-Through Debugging
//
// Goal: Distinguish Delve's three stepping commands — next, step, and
// stepout — by practicing all three on real nested function calls.
package main

import "fmt"

func formatPrice(cents int) string {
	dollars := cents / 100
	remainder := cents % 100
	return fmt.Sprintf("$%d.%02d", dollars, remainder)
}

func applyTax(cents int, taxPercent float64) int {
	tax := float64(cents) * taxPercent / 100
	return cents + int(tax)
}

func checkout(cents int) string {
	withTax := applyTax(cents, 8.5)
	return formatPrice(withTax)
}

func main() {
	result := checkout(1999)
	fmt.Println("Final price:", result)
	fmt.Println("\nSee the README to practice next/step/stepout on this exact call chain.")
}
