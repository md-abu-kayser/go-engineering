// Lesson 15: Numeric Conversion
//
// Goal: Convert explicitly between numeric types — Go NEVER converts
// implicitly — and see exactly what happens on truncation (float->int)
// and overflow (large type -> smaller type).
package main

import "fmt"

func main() {
	fmt.Println("=== Numeric Conversion ===")
	fmt.Println("----------------------------------")

	// Go requires an EXPLICIT conversion between ANY two different
	// numeric types — even int to int64, or int32 to int. There is no
	// implicit widening or narrowing, unlike many other languages.
	var i int = 42
	var i64 int64 = int64(i) // explicit, even though this is always safe
	var f float64 = float64(i)
	fmt.Printf("int(42) -> int64: %d, -> float64: %v\n", i64, f)

	// float -> int TRUNCATES (discards the fractional part) — it does
	// NOT round, and it does NOT error, even for a value with a large
	// fractional part. Note: this must go through VARIABLES — an
	// untyped float CONSTANT with a fractional part cannot convert to
	// int at all (a COMPILE error, per lesson 05's untyped-constant rules).
	fmt.Println("\n--- float -> int truncates, never rounds ---")
	var f1, f2, f3 float64 = 3.99, 3.01, -3.99
	fmt.Printf("int(3.99)  = %d\n", int(f1))
	fmt.Printf("int(3.01)  = %d\n", int(f2))
	fmt.Printf("int(-3.99) = %d (truncates TOWARD ZERO, not down)\n", int(f3))

	// Converting a LARGE value into a SMALLER integer type silently
	// discards the high-order bits — this can produce a wildly
	// different, seemingly nonsensical number, with NO error or panic.
	fmt.Println("\n--- Narrowing conversion can silently lose data ---")
	var big int32 = 300
	var narrow int8 = int8(big) // int8 can only hold up to 127
	fmt.Printf("int32(300) -> int8 = %d (NOT 300 — the high bits were simply discarded)\n", narrow)

	var big2 int64 = 40000
	var narrow2 int16 = int16(big2) // int16 max is 32767
	fmt.Printf("int64(40000) -> int16 = %d (wrapped, same silent truncation)\n", narrow2)
}
