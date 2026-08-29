// Lesson 22: Assignment Operators
//
// Goal: Use every compound assignment operator Go supports — the
// "do an operation AND assign the result back" shorthand for each of
// the arithmetic and bitwise operators.
package main

import "fmt"

func main() {
	fmt.Println("=== Assignment Operators ===")
	fmt.Println("----------------------------------")

	n := 10
	fmt.Printf("n := 10          -> n = %d\n", n)

	n += 5 // equivalent to: n = n + 5
	fmt.Printf("n += 5           -> n = %d\n", n)

	n -= 3 // equivalent to: n = n - 3
	fmt.Printf("n -= 3           -> n = %d\n", n)

	n *= 2 // equivalent to: n = n * 2
	fmt.Printf("n *= 2           -> n = %d\n", n)

	n /= 4 // equivalent to: n = n / 4
	fmt.Printf("n /= 4           -> n = %d\n", n)

	n %= 3 // equivalent to: n = n %% 3
	fmt.Printf("n %%= 3           -> n = %d\n", n)

	// Bitwise compound assignments follow the exact same pattern.
	fmt.Println("\n--- Bitwise compound assignments ---")
	flags := 0b1100
	flags &= 0b1010 // flags = flags & 0b1010
	fmt.Printf("flags &= 0b1010  -> %04b\n", flags)

	flags |= 0b0001 // flags = flags | 0b0001
	fmt.Printf("flags |= 0b0001  -> %04b\n", flags)

	flags ^= 0b1111 // flags = flags ^ 0b1111
	fmt.Printf("flags ^= 0b1111  -> %04b\n", flags)

	flags &^= 0b0010 // flags = flags &^ 0b0010 (bit clear, lesson 21)
	fmt.Printf("flags &^= 0b0010 -> %04b\n", flags)

	shifted := 1
	shifted <<= 3 // shifted = shifted << 3
	fmt.Printf("shifted <<= 3    -> %d\n", shifted)

	shifted >>= 1 // shifted = shifted >> 1
	fmt.Printf("shifted >>= 1    -> %d\n", shifted)

	// += also works for string concatenation, since + does (lesson 18).
	fmt.Println("\n--- += on strings ---")
	msg := "Hello"
	msg += ", Gopher!"
	fmt.Printf("msg += \", Gopher!\" -> %q\n", msg)
}
