// Lesson 54: Batching
//
// Goal: Coordinate a bounded set of goroutines, collect every result, and
// close the result channel only after all producers have stopped.
package main

import (
	"fmt"
	"sync"
)

const lesson = "Batching"

func squareAll(values []int) []int {
	type result struct {
		index int
		value int
	}
	results := make(chan result, len(values))
	var group sync.WaitGroup
	for index, value := range values {
		group.Add(1)
		go func(index, n int) {
			defer group.Done()
			results <- result{index: index, value: n * n}
		}(index, value)
	}
	group.Wait()
	close(results)

	output := make([]int, len(values))
	for result := range results {
		output[result.index] = result.value
	}
	return output
}

func main() {
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("completed work: %v\n", squareAll([]int{2, 3, 4}))
}