// Lesson 30: gopls Overview
//
// Goal: Understand what gopls actually is, and deliberately exercise a
// few of its features on this file inside your editor.
package main

import "fmt"

// invoice holds line items for a simple, contrived example — practice
// hovering, go-to-definition, and rename on these declarations.
type invoice struct {
	items []lineItem
}

type lineItem struct {
	description string
	cents       int
}

func (inv invoice) totalCents() int {
	total := 0
	for _, item := range inv.items {
		total += item.cents
	}
	return total
}

func main() {
	inv := invoice{
		items: []lineItem{
			{description: "Coffee", cents: 350},
			{description: "Muffin", cents: 275},
		},
	}

	fmt.Printf("Total: $%d.%02d\n", inv.totalCents()/100, inv.totalCents()%100)
	fmt.Println("\nSee the README for specific gopls features to try on this file.")
}
