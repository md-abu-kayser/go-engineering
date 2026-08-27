// Lesson 32: Breakpoints
//
// Goal: Go beyond a single, unconditional breakpoint (lesson 31) and use
// Delve's conditional breakpoints and tracepoints — the tools you reach
// for once "just stop here every time" isn't precise enough.
package main

import "fmt"

// findFirstOver returns the first value in nums that is strictly greater
// than threshold, and its index. It returns (-1, -1) if none qualifies.
func findFirstOver(nums []int, threshold int) (value, index int) {
	for i, n := range nums {
		if n > threshold {
			return n, i
		}
	}
	return -1, -1
}

func main() {
	nums := []int{3, 7, 2, 9, 4, 15, 6}

	value, index := findFirstOver(nums, 8)
	fmt.Printf("First value over 8: %d at index %d\n", value, index)
	fmt.Println("\nSee the README for how to break ONLY on the loop iteration that matters,")
	fmt.Println("instead of stepping through every single one.")
}
