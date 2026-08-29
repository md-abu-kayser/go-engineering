// Lesson 16: String Conversion
//
// Goal: Convert numbers to and from strings using the strconv package —
// and see the classic gotcha of string(anInt), which does NOT do what
// most beginners expect.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== String Conversion ===")
	fmt.Println("----------------------------------")

	// The IDIOMATIC way to convert a number to its string representation.
	n := 42
	s := strconv.Itoa(n) // "Integer to ASCII" — an old, C-derived name
	fmt.Printf("strconv.Itoa(42)        = %q\n", s)

	// The IDIOMATIC way to parse a string back into a number — note it
	// returns an error too, since not every string is a valid number.
	parsed, err := strconv.Atoi("123")
	fmt.Printf("strconv.Atoi(\"123\")     = %d, err = %v\n", parsed, err)

	_, err = strconv.Atoi("not a number")
	fmt.Printf("strconv.Atoi(\"not a number\") errors: %v\n", err)

	// Floats: FormatFloat and ParseFloat.
	f := 3.14159
	fs := strconv.FormatFloat(f, 'f', 2, 64) // 'f' format, 2 decimal places, 64-bit
	fmt.Printf("strconv.FormatFloat(3.14159, 2 places) = %q\n", fs)

	parsedFloat, _ := strconv.ParseFloat("2.71828", 64)
	fmt.Printf("strconv.ParseFloat(\"2.71828\")          = %v\n", parsedFloat)

	// THE classic gotcha: string(someInt) does NOT give you the number
	// as text — it converts the int as a UNICODE CODE POINT, giving
	// you whatever CHARACTER has that code point instead.
	fmt.Println("\n--- The classic gotcha ---")
	code := 65
	wrong := string(rune(code)) // Go requires the explicit rune() now, precisely BECAUSE of this confusion
	right := strconv.Itoa(code)
	fmt.Printf("string(rune(65)) = %q (the CHARACTER with code point 65 — 'A')\n", wrong)
	fmt.Printf("strconv.Itoa(65) = %q (the actual TEXT \"65\" — almost always what you want)\n", right)
}
