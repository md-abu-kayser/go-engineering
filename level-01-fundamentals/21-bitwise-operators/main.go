// Lesson 21: Bitwise Operators
//
// Goal: Use Go's bitwise operators — including &^ ("AND NOT" / bit
// clear), a Go-specific operator most other C-family languages don't
// have a dedicated symbol for.
package main

import "fmt"

func main() {
	fmt.Println("=== Bitwise Operators ===")
	fmt.Println("----------------------------------")

	a, b := 0b1100, 0b1010 // binary literals: 12 and 10

	fmt.Printf("a = %04b (%d), b = %04b (%d)\n", a, a, b, b)
	fmt.Printf("a & b  = %04b (%d)  (AND: 1 where BOTH bits are 1)\n", a&b, a&b)
	fmt.Printf("a | b  = %04b (%d)  (OR: 1 where EITHER bit is 1)\n", a|b, a|b)
	fmt.Printf("a ^ b  = %04b (%d)  (XOR: 1 where the bits DIFFER)\n", a^b, a^b)
	fmt.Printf("^a     = %d  (unary NOT: flips every bit — sign matters for signed types)\n", ^a)

	// &^ is Go's OWN operator, not found in most other C-family
	// languages under a dedicated symbol: "AND NOT", aka BIT CLEAR.
	// a &^ b means: take a, but force off any bit that's set in b.
	fmt.Printf("a &^ b = %04b (%d)  (BIT CLEAR: a, with b's set bits forced OFF)\n", a&^b, a&^b)

	// Shifting: << moves bits left (multiplying by a power of 2),
	// >> moves bits right (dividing by a power of 2, for unsigned/
	// non-negative values).
	fmt.Println("\n--- Shifting ---")
	x := 1
	fmt.Printf("1 << 4 = %d (same as 1 * 2^4)\n", x<<4)
	fmt.Printf("16 >> 2 = %d (same as 16 / 2^2)\n", 16>>2)

	// The flag-combining pattern from lesson 06 (iota), shown again
	// with &^ specifically for REMOVING one flag from a combined set.
	fmt.Println("\n--- Practical use: removing one flag from a set ---")
	const (
		FlagRead  = 1 << iota // 1
		FlagWrite             // 2
		FlagExec              // 4
	)
	perms := FlagRead | FlagWrite | FlagExec // all three: 7
	permsNoWrite := perms &^ FlagWrite       // remove JUST the write bit
	fmt.Printf("perms          = %03b (all three flags)\n", perms)
	fmt.Printf("perms &^ Write = %03b (write flag cleanly removed)\n", permsNoWrite)
}
