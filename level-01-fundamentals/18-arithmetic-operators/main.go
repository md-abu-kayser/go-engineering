// Lesson 18: Arithmetic Operators
//
// Goal: Use Go's five arithmetic operators across both numbers and
// strings, and see exactly which types + supports beyond addition.
package main

import "fmt"

func main() {
	fmt.Println("=== Arithmetic Operators ===")
	fmt.Println("----------------------------------")

	a, b := 17, 5

	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d (integer division truncates — lesson 11)\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d (remainder)\n", a, b, a%b)

	// Unary minus negates a value.
	fmt.Printf("\n-%d = %d (unary minus)\n", a, -a)

	// + is ALSO Go's string concatenation operator — it's the ONE
	// arithmetic-looking operator that works on a non-numeric type.
	// No other arithmetic operator (-, *, /, %) works on strings at all.
	fmt.Println("\n--- + also means string concatenation ---")
	greeting := "Hello, " + "Gopher" + "!"
	fmt.Printf(`"Hello, " + "Gopher" + "!" = %q`+"\n", greeting)

	// Float arithmetic works the same way, except / is TRUE division,
	// not truncating — since there's no integer-truncation rule for floats.
	fmt.Println("\n--- Float arithmetic ---")
	fa, fb := 17.0, 5.0
	fmt.Printf("%.1f / %.1f = %v (true division, not truncated)\n", fa, fb, fa/fb)
}
