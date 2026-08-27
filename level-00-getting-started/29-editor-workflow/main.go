// Lesson 29: Editor Workflow (VS Code)
//
// Goal: Practice the everyday VS Code + Go extension workflow: run,
// format-on-save, go-to-definition, and debugging — on real code.
package main

import "fmt"

// splitBill divides amount evenly among people and returns each
// person's share, rounded down to the nearest cent... intentionally
// simplified (integer division) so it's easy to reason about while
// stepping through it in the debugger.
func splitBill(amountCents, people int) int {
	if people <= 0 {
		return 0
	}
	return amountCents / people
}

func main() {
	share := splitBill(2599, 4)
	fmt.Printf("Each person owes: $%d.%02d\n", share/100, share%100)

	fmt.Println("\nSee the README for a checklist of editor features to try on this file.")
}
