// Lesson 44: Multiple Return Values
//
// Goal: Return more than one value from a function — the mechanism
// behind Go's idiomatic (value, error) and (value, ok) patterns — and
// use the blank identifier to discard a return value you don't need.
package main

import (
	"fmt"
)

// divide returns TWO values: the result, and an error if the divisor
// was zero. This IS the idiomatic Go shape for "an operation that can
// fail" — no exceptions, just an extra return value.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

// minMax returns two values that are BOTH genuinely useful results,
// not an error pattern — multiple return values aren't only for
// error handling.
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// lookup returns a value and a boolean "ok" — the OTHER extremely
// common multi-return pattern, used for map lookups, type assertions,
// and anywhere "did this succeed?" matters more than a specific error.
func lookup(m map[string]int, key string) (int, bool) {
	value, ok := m[key]
	return value, ok
}

func main() {
	fmt.Println("=== Multiple Return Values ===")
	fmt.Println("----------------------------------")

	result, err := divide(10, 2)
	fmt.Printf("divide(10, 2) = %d, err = %v\n", result, err)

	result, err = divide(10, 0)
	fmt.Printf("divide(10, 0) = %d, err = %v\n", result, err)

	fmt.Println("\n--- Two genuinely useful values, not an error pattern ---")
	lo, hi := minMax([]int{5, 2, 8, 1, 9, 3})
	fmt.Printf("minMax([5 2 8 1 9 3]) = min %d, max %d\n", lo, hi)

	fmt.Println("\n--- The (value, ok) pattern ---")
	ages := map[string]int{"Alice": 30}
	age, ok := lookup(ages, "Alice")
	fmt.Printf("lookup(ages, \"Alice\") = %d, ok = %t\n", age, ok)
	age, ok = lookup(ages, "Bob")
	fmt.Printf("lookup(ages, \"Bob\")   = %d, ok = %t\n", age, ok)

	// The blank identifier discards a return value you genuinely don't
	// need — common when you only care whether an operation SUCCEEDED,
	// not its actual result value.
	fmt.Println("\n--- Discarding a return value with _ ---")
	_, err = divide(20, 4)
	fmt.Printf("divide(20, 4): only checking err, ignoring the quotient -> err = %v\n", err)
}
