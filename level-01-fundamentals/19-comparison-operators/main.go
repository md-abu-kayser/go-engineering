// Lesson 19: Comparison Operators
//
// Goal: Use Go's six comparison operators, and understand COMPARABILITY
// — which types support == at all, including the surprising case of
// structs (comparable, IF every field is) vs slices/maps (never
// comparable with ==).
package main

import "fmt"

type point struct{ X, Y int }

func main() {
	fmt.Println("=== Comparison Operators ===")
	fmt.Println("----------------------------------")

	a, b := 10, 20
	fmt.Printf("%d == %d : %t\n", a, b, a == b)
	fmt.Printf("%d != %d : %t\n", a, b, a != b)
	fmt.Printf("%d <  %d : %t\n", a, b, a < b)
	fmt.Printf("%d <= %d : %t\n", a, b, a <= b)
	fmt.Printf("%d >  %d : %t\n", a, b, a > b)
	fmt.Printf("%d >= %d : %t\n", a, b, a >= b)

	// Strings compare LEXICOGRAPHICALLY (dictionary order, by byte
	// value) — not by length.
	fmt.Println("\n--- String comparison is lexicographic ---")
	fmt.Printf("%q < %q : %t (compares byte-by-byte, like dictionary order)\n", "apple", "banana", "apple" < "banana")
	fmt.Printf("%q < %q  : %t (shorter isn't automatically \"less\")\n", "zoo", "apple", "zoo" < "apple")

	// STRUCTS are comparable with == IF (and only if) every one of
	// their fields is itself comparable.
	fmt.Println("\n--- Struct comparability ---")
	p1 := point{X: 1, Y: 2}
	p2 := point{X: 1, Y: 2}
	p3 := point{X: 3, Y: 4}
	fmt.Printf("p1 == p2 : %t (same field values)\n", p1 == p2)
	fmt.Printf("p1 == p3 : %t (different field values)\n", p1 == p3)

	// Slices, maps, and functions are NEVER comparable with == — this
	// is a COMPILE ERROR, not a runtime false. (The one exception:
	// comparing any of them to the literal `nil` IS allowed.)
	fmt.Println("\n--- Slices/maps are NOT comparable ---")
	fmt.Println("See the README: `mySlice == otherSlice` is a compile-time error, not `false`.")
	var s []int
	fmt.Printf("s == nil : %t (comparing to nil IS allowed, even though slice == slice is not)\n", s == nil)
}
