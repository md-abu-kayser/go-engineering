// Lesson 08: Strings
//
// Goal: Understand Go's string type: UTF-8 encoded, immutable, and
// measured in BYTES by len() — not necessarily "characters".
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Strings ===")
	fmt.Println("----------------------------------")

	greeting := "Hello, 世界"

	// len() on a string counts BYTES, not "characters" — and in UTF-8,
	// non-ASCII characters take MORE than one byte each.
	fmt.Printf("greeting          : %s\n", greeting)
	fmt.Printf("len(greeting)     : %d bytes\n", len(greeting))

	// Indexing a string with [i] gives you a single BYTE, not a
	// character — for ASCII text this happens to look right, but it's
	// genuinely wrong for multi-byte characters (see lesson 09 for the
	// correct way, using runes).
	fmt.Printf("greeting[0]       : %d (%q) — the BYTE at index 0\n", greeting[0], string(greeting[0]))

	// Strings are IMMUTABLE — there is no way to change a byte in
	// place. "Modifying" a string always means building a NEW one.
	upper := strings.ToUpper(greeting)
	fmt.Printf("strings.ToUpper() : %s (a NEW string — greeting itself is unchanged)\n", upper)
	fmt.Printf("greeting (after)  : %s (still the original)\n", greeting)

	// Concatenation with + also always builds a new string.
	full := greeting + "!"
	fmt.Printf("greeting + \"!\"    : %s\n", full)

	// Common strings package helpers.
	fmt.Printf("strings.Contains  : %t\n", strings.Contains(greeting, "世界"))
	fmt.Printf("strings.Split     : %q\n", strings.Split("a,b,c", ","))
	fmt.Printf("strings.TrimSpace : %q\n", strings.TrimSpace("  padded  "))
}
