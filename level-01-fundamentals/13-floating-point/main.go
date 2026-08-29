// Lesson 13: Floating-Point
//
// Goal: Use float32/float64, and see IEEE 754's classic precision
// surprise (0.1 + 0.2 != 0.3) directly, plus the correct way to compare
// floats for "close enough" equality.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Floating-Point ===")
	fmt.Println("----------------------------------")

	// float64 is Go's default floating-point type (what a bare literal
	// like 3.14 becomes) — use it unless you have a specific reason
	// (matching a file format, saving memory in a huge array) for float32.
	var f64 float64 = 3.14159265358979
	var f32 float32 = 3.14159265358979
	fmt.Printf("float64 : %.14f (full precision)\n", f64)
	fmt.Printf("float32 : %.14f (visibly LESS precise — fewer bits to work with)\n", f32)

	// THE classic floating-point surprise: 0.1 and 0.2 cannot be
	// represented EXACTLY in binary floating point, so their sum isn't
	// EXACTLY 0.3 either — but this needs to be demonstrated CAREFULLY.
	fmt.Println("\n--- The classic surprise (and a Go-specific subtlety) ---")

	// If you write 0.1 + 0.2 as bare LITERALS, Go's compiler evaluates
	// untyped constant expressions using ARBITRARY PRECISION (see
	// lesson 05) — the rounding to float64 happens only ONCE, at the
	// very end, which can accidentally land EXACTLY on the same
	// rounding 0.3 would get. So this does NOT reliably show the surprise:
	constSum := 0.1 + 0.2
	fmt.Printf("0.1 + 0.2 (const-folded)   = %.20f\n", constSum)
	fmt.Printf("  == 0.3 ?                   %t (misleading! see below for why)\n", constSum == 0.3)

	// To see the REAL surprise, the values must actually be float64
	// VARIABLES at the time of addition, so each one is independently
	// rounded to float64 BEFORE they're added together:
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3
	runtimeSum := a + b
	fmt.Printf("a + b (real float64 runtime arithmetic) = %.20f\n", runtimeSum)
	fmt.Printf("  == c (float64 0.3) ?                     %t (the REAL, reliable surprise)\n", runtimeSum == c)

	// The correct way to compare floats: check they're within a small
	// EPSILON of each other, never use == directly.
	const epsilon = 1e-9
	closeEnough := math.Abs(runtimeSum-c) < epsilon
	fmt.Printf("  within epsilon of 0.3 ?                   %t (this is the reliable check)\n", closeEnough)

	// Special values: division by zero on FLOATS doesn't panic (unlike
	// integers) — it produces +Inf, -Inf, or NaN. Note: this only works
	// with VARIABLES, not literal constants — 1.0/0.0 as a literal
	// expression is itself a COMPILE-TIME error (division by zero is
	// caught at compile time when both operands are constants).
	fmt.Println("\n--- Special values ---")
	var zero float64 = 0.0
	posInf := 1.0 / zero
	negInf := -1.0 / zero
	notANumber := math.NaN()
	fmt.Printf("1.0 / 0.0  = %v\n", posInf)
	fmt.Printf("-1.0 / 0.0 = %v\n", negInf)
	fmt.Printf("NaN        = %v\n", notANumber)
	fmt.Printf("NaN == NaN ? %t (NaN is NEVER equal to anything, including itself)\n", notANumber == notANumber)
	fmt.Printf("math.IsNaN(NaN) ? %t (use this instead of == to check for NaN)\n", math.IsNaN(notANumber))
}
