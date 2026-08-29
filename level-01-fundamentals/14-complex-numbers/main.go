// Lesson 14: Complex Numbers
//
// Goal: Use Go's built-in complex64/complex128 types — a genuinely
// unusual feature for a systems-adjacent language — including
// arithmetic and the real()/imag() built-ins.
package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	fmt.Println("=== Complex Numbers ===")
	fmt.Println("----------------------------------")

	// Complex literals use the imaginary suffix `i`. complex128 (based
	// on two float64s) is the default, just like float64 is the default
	// real type.
	c1 := 3 + 4i
	c2 := complex(1.5, -2.5) // building one from real and imaginary parts explicitly

	fmt.Printf("c1 = %v (type %T)\n", c1, c1)
	fmt.Printf("c2 = %v\n", c2)

	// real() and imag() extract the two components as plain float64s.
	fmt.Printf("real(c1) = %v, imag(c1) = %v\n", real(c1), imag(c1))

	// Ordinary arithmetic operators work directly on complex numbers.
	fmt.Println("\n--- Arithmetic ---")
	fmt.Printf("c1 + c2 = %v\n", c1+c2)
	fmt.Printf("c1 * c2 = %v\n", c1*c2)

	// math/cmplx provides complex-number-aware versions of common math
	// functions — e.g. Abs computes the MAGNITUDE (distance from origin
	// in the complex plane), which for 3+4i is the classic 3-4-5 triangle.
	fmt.Println("\n--- math/cmplx ---")
	fmt.Printf("cmplx.Abs(3+4i) = %v (the classic 3-4-5 right triangle)\n", cmplx.Abs(c1))

	// complex64 is the smaller variant, built from two float32s —
	// exists for the same "specific memory/format constraint" reasons
	// float32 does.
	var small complex64 = 1 + 2i
	fmt.Printf("\ncomplex64 example: %v (type %T)\n", small, small)
}
