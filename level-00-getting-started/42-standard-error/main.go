// Lesson 42: Standard Error
//
// Goal: Understand why diagnostics and error messages belong on
// os.Stderr, separate from a program's real output on os.Stdout — and
// see the two streams redirected independently from the shell.
package main

import (
	"fmt"
	"os"
)

func processItem(name string, price int) error {
	if price < 0 {
		return fmt.Errorf("item %q has an invalid negative price: %d", name, price)
	}
	return nil
}

func main() {
	items := map[string]int{
		"Keyboard": 4999,
		"Mouse":    -100, // deliberately invalid, to produce a real diagnostic below
		"Monitor":  15999,
	}

	names := []string{"Keyboard", "Mouse", "Monitor"} // fixed order, for predictable output

	for _, name := range names {
		price := items[name]
		if err := processItem(name, price); err != nil {
			// Diagnostics go to STDERR, not stdout — this is the whole point
			// of this lesson.
			fmt.Fprintln(os.Stderr, "warning:", err)
			continue
		}
		// Real, intended output goes to STDOUT.
		fmt.Printf("%-10s $%.2f\n", name, float64(price)/100)
	}
}
