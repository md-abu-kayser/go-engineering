// Lesson 46: Cache Locality
//
// Goal: Make the allocation decision visible: preallocate the exact output
// size and keep the hot loop free of formatting and interface conversion.
package main

import "fmt"

const lesson = "Cache Locality"

func double(values []int) []int {
	output := make([]int, len(values))
	for index, value := range values {
		output[index] = value * 2
	}
	return output
}

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("input: %v output: %v\n", input, double(input))
}