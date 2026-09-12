// Lesson 05: Union Constraints
//
// Goal: Use a type constraint to write one operation that remains type-safe
// for every numeric type it accepts.
package main

import "fmt"

const lesson = "Union Constraints"

// Number is intentionally small: callers may use integer or floating-point
// values without giving up compile-time checking.
type Number interface {
	~int | ~int64 | ~float64
}

func total[T Number](values ...T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

func main() {
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("integer total: %d\n", total(3, 5, 8))
	fmt.Printf("decimal total: %.1f\n", total(1.5, 2.5))
}