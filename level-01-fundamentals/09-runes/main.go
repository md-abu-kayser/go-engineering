// Lesson 09: Runes
//
// Goal: Use `rune` (Go's name for a Unicode code point, really just an
// int32) to correctly handle text that might contain multi-byte
// characters — fixing the exact problem lesson 08 identified.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== Runes ===")
	fmt.Println("----------------------------------")

	greeting := "Hello, 世界"

	// The CORRECT way to count "characters" (Unicode code points):
	// utf8.RuneCountInString, NOT len().
	fmt.Printf("len(greeting)               : %d (bytes)\n", len(greeting))
	fmt.Printf("utf8.RuneCountInString(...) : %d (actual characters)\n", utf8.RuneCountInString(greeting))

	// Converting a string to []rune gives you a slice of actual
	// characters — now indexing works correctly, at the cost of an
	// upfront decoding pass over the whole string.
	runes := []rune(greeting)
	fmt.Printf("[]rune(greeting)[7]          : %c (the ACTUAL 8th character, 世)\n", runes[7])

	// `range` over a STRING (not a []rune) also decodes correctly,
	// WITHOUT needing to convert first — and gives you each rune's
	// starting BYTE INDEX alongside it.
	fmt.Println("\nranging over the string directly:")
	for i, r := range greeting {
		fmt.Printf("  byte index %2d: %c (%U)\n", i, r, r)
	}

	// A rune is really just an int32 — you can do arithmetic on it.
	var r rune = 'A'
	fmt.Printf("\n'A' as a rune: %d, +1 = %c\n", r, r+1)
}
