// Lesson 30: for Loop
//
// Goal: Use Go's classic three-part `for` loop — init; condition; post
// — the ONLY looping keyword Go has (there is no separate `while` or
// `do-while`, previewed further in lesson 31).
package main

import "fmt"

func main() {
	fmt.Println("=== for Loop ===")
	fmt.Println("----------------------------------")

	// The classic three-part form: init; condition; post.
	// - init runs ONCE, before the loop starts (i := 0)
	// - condition is checked BEFORE every iteration (i < 5)
	// - post runs AFTER every iteration's body (i++)
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// The init variable is scoped to the loop ONLY — exactly like
	// if/switch's init statements (lessons 25, 28).
	fmt.Println("\nSee the README: `i` above is NOT visible out here.")

	// Counting DOWN instead of up — just change all three parts
	// accordingly.
	fmt.Println("\n--- Counting down ---")
	for i := 3; i > 0; i-- {
		fmt.Printf("i = %d\n", i)
	}

	// Stepping by more than 1 — the "post" part can be any statement,
	// not just ++ / --.
	fmt.Println("\n--- Stepping by 2 ---")
	for i := 0; i < 10; i += 2 {
		fmt.Printf("i = %d\n", i)
	}

	// break exits the loop entirely; continue skips to the NEXT
	// iteration's condition check, without running the rest of the body.
	fmt.Println("\n--- break and continue ---")
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // skip printing 3, but keep looping
		}
		if i == 6 {
			break // stop looping entirely once we reach 6
		}
		fmt.Printf("i = %d\n", i)
	}
}
