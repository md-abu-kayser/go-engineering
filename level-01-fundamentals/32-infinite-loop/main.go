// Lesson 32: Infinite Loop
//
// Goal: Use `for {}` — Go's explicit, deliberate infinite loop shape —
// and the correct way to exit one with `break`, the pattern behind
// every event loop, server loop, and polling loop.
package main

import "fmt"

func main() {
	fmt.Println("=== Infinite Loop ===")
	fmt.Println("----------------------------------")

	// `for` with NOTHING after it — no init, no condition, no post —
	// loops forever, by design. This is Go's explicit way to say
	// "keep going until something inside the loop decides to stop."
	count := 0
	for {
		count++
		fmt.Printf("iteration %d\n", count)
		if count >= 5 {
			break // the ONLY thing stopping this loop — remove it, and it never ends
		}
	}
	fmt.Println("\nLoop exited cleanly via break.")

	// A more realistic shape: simulating "keep processing items from a
	// queue until it's empty" — the loop itself has no natural
	// "count" to bound it; it just runs until the WORK runs out.
	fmt.Println("\n--- Realistic pattern: process until empty ---")
	queue := []string{"task-1", "task-2", "task-3"}
	for {
		if len(queue) == 0 {
			fmt.Println("queue is empty, stopping")
			break
		}
		item := queue[0]
		queue = queue[1:] // "pop" the front item
		fmt.Printf("processing %s (queue has %d item(s) left)\n", item, len(queue))
	}

	fmt.Println("\nThis exact shape — for {} with an internal break — is how event loops,")
	fmt.Println("server accept loops, and polling loops are almost always written in Go.")
}
