// Lesson 50: Recursion
//
// Goal: Write a recursive function — one that calls itself — with a
// clear base case and recursive case, and see why Go does NOT optimize
// tail-recursive calls the way some other languages do.
package main

import "fmt"

// factorial is the classic first recursion example: n! = n * (n-1)!,
// with the BASE CASE stopping the recursion at n <= 1.
func factorial(n int) int {
	if n <= 1 {
		return 1 // BASE CASE — stops the recursion
	}
	return n * factorial(n-1) // RECURSIVE CASE — calls itself with a SMALLER problem
}

// fibonacci is the second classic example — note this NAIVE version
// recomputes the same values repeatedly (fibonacci(5) calls
// fibonacci(3) TWICE, fibonacci(2) THREE times, etc.), which is fine
// for teaching but genuinely inefficient for larger n.
func fibonacci(n int) int {
	if n <= 1 {
		return n // base case: fibonacci(0) = 0, fibonacci(1) = 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// sumDigits recursively sums the digits of a non-negative number —
// a THIRD shape of recursion, operating on a number instead of
// counting down by 1 each time.
func sumDigits(n int) int {
	if n < 10 {
		return n // base case: a single digit sums to itself
	}
	return n%10 + sumDigits(n/10) // last digit + recurse on the rest
}

func main() {
	fmt.Println("=== Recursion ===")
	fmt.Println("----------------------------------")

	for _, n := range []int{0, 1, 5, 10} {
		fmt.Printf("factorial(%d)  = %d\n", n, factorial(n))
	}

	fmt.Println()
	for _, n := range []int{0, 1, 5, 10} {
		fmt.Printf("fibonacci(%d)  = %d\n", n, fibonacci(n))
	}

	fmt.Println()
	for _, n := range []int{7, 123, 9999} {
		fmt.Printf("sumDigits(%d) = %d\n", n, sumDigits(n))
	}

	fmt.Println("\nSee the README for why Go does NOT optimize tail-recursive calls,")
	fmt.Println("and what that means for how deep a recursive function can safely go.")
}
