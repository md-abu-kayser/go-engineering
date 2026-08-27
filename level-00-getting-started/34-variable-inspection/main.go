// Lesson 34: Variable Inspection
//
// Goal: Use Delve's print, whatis, and set commands to inspect — and
// even change — variables of different shapes (struct, slice, map)
// mid-execution.
package main

import "fmt"

// order and its fields give us a struct, a slice, and a map to inspect —
// three of the most common shapes you'll want to look inside while
// debugging.
type order struct {
	ID       int
	Items    []string
	Metadata map[string]string
}

func summarize(o order) string {
	total := len(o.Items)
	return fmt.Sprintf("Order #%d has %d item(s)", o.ID, total)
}

func main() {
	o := order{
		ID:    42,
		Items: []string{"Keyboard", "Mouse", "Monitor"},
		Metadata: map[string]string{
			"priority": "high",
			"region":   "eu-west",
		},
	}

	summary := summarize(o)
	fmt.Println(summary)
	fmt.Println("\nSee the README to inspect (and even modify) o's fields with dlv print/set.")
}
