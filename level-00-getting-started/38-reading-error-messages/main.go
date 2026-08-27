// Lesson 38: Reading Error Messages
//
// Goal: Read and construct Go's `error` values confidently — including
// wrapping errors with context and unwrapping them again with the
// standard library's errors package.
package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is a SENTINEL ERROR — a specific, comparable error value
// other code can check for by identity, using errors.Is.
var ErrNotFound = errors.New("item not found")

var inventory = map[string]int{
	"apple":  12,
	"banana": 0,
}

// lookup returns the stock count for name, or a wrapped ErrNotFound if
// name isn't in the inventory at all.
func lookup(name string) (int, error) {
	count, ok := inventory[name]
	if !ok {
		// %w (not %s or %v) wraps ErrNotFound, preserving its identity
		// for errors.Is, while adding human-readable context.
		return 0, fmt.Errorf("lookup %q: %w", name, ErrNotFound)
	}
	return count, nil
}

func main() {
	for _, item := range []string{"apple", "banana", "cherry"} {
		count, err := lookup(item)
		switch {
		case err == nil:
			fmt.Printf("%-8s: %d in stock\n", item, count)
		case errors.Is(err, ErrNotFound):
			fmt.Printf("%-8s: %v\n", item, err)
		default:
			fmt.Printf("%-8s: unexpected error: %v\n", item, err)
		}
	}

	fmt.Println("\nSee the README for exactly what %w does, and why errors.Is beats ==.")
}
