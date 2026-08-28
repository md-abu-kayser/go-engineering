// Lesson 10: Bytes
//
// Goal: Use `byte` (an alias for uint8) and []byte, and convert cleanly
// between strings and byte slices — the most common real-world use of
// byte-level string handling.
package main

import "fmt"

func main() {
	fmt.Println("=== Bytes ===")
	fmt.Println("----------------------------------")

	// byte is LITERALLY just an alias for uint8 — same type, different
	// conventional name, used when you mean "raw byte data" rather than
	// "a small non-negative number".
	var b byte = 65
	fmt.Printf("var b byte = 65   : %d, as a character: %c\n", b, b)

	// Converting a string to []byte gives you its RAW UTF-8 bytes,
	// exactly what you'd get from len() in lesson 08 — this conversion
	// is cheap (no decoding), unlike converting to []rune (lesson 09).
	greeting := "Hello, 世界"
	data := []byte(greeting)
	fmt.Printf("[]byte(greeting)  : %d bytes, first few: %v\n", len(data), data[:5])

	// []byte is MUTABLE, unlike string — you can change individual
	// bytes in place. This is a genuinely different value from the
	// original string once you do.
	data[0] = 'h' // lowercase the first byte
	fmt.Printf("after data[0]='h' : %s (converted back to string)\n", string(data))
	fmt.Printf("original greeting : %s (completely unaffected — data is a SEPARATE copy)\n", greeting)

	// []byte is the type most I/O functions in Go actually work with
	// (file reads, network reads, etc.) — converting to/from string is
	// how you bridge "raw bytes from the outside world" and "text".
	fmt.Printf("\nstring(data)      : %s\n", string(data))
}
