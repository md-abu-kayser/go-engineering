// Lesson 12: Unsigned Integers
//
// Goal: Understand Go's unsigned integer types, when to reach for them
// deliberately, and the specific danger of unsigned UNDERFLOW —
// arguably the single most common integer bug in Go.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Unsigned Integers ===")
	fmt.Println("----------------------------------")

	// uint8 (== byte, lesson 10) and friends can only represent
	// NON-NEGATIVE values, but get double the positive range of their
	// signed counterpart in exchange.
	var u8 uint8 = 200
	fmt.Printf("uint8        : %d (range: 0 to %d)\n", u8, math.MaxUint8)
	fmt.Printf("int8 max     : %d  <- uint8 goes almost twice as high, using the same 8 bits\n", math.MaxInt8)

	// len() and cap() return `int`, not an unsigned type — a deliberate
	// Go design choice, precisely to avoid the underflow trap below.
	s := []int{1, 2, 3}
	fmt.Printf("\nlen(s)       : %d (type: %T — a signed int, on purpose)\n", len(s), len(s))

	// THE classic unsigned bug: subtracting past zero doesn't go
	// negative — it WRAPS AROUND to a huge positive number instead.
	fmt.Println("\n--- Unsigned underflow ---")
	var count uint = 0
	count--
	fmt.Printf("uint(0) - 1  = %d (NOT -1 — uint cannot represent negative numbers)\n", count)

	// This exact bug pattern is why a naive "count down to empty" loop
	// over an unsigned type can loop seemingly forever.
	var itemsLeft uint8 = 3
	fmt.Println("\n--- Why this matters in a loop ---")
	for i := 0; i < 5; i++ { // bounded manually to keep this demo finite
		fmt.Printf("  itemsLeft = %d\n", itemsLeft)
		if itemsLeft == 0 {
			fmt.Println("  (stopping the DEMO here — a naive `for itemsLeft >= 0` loop would NOT have stopped)")
			break
		}
		itemsLeft--
	}
}
