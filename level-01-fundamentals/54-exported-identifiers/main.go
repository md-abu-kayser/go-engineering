// Lesson 54: Exported Identifiers
//
// Goal: Use an EXPORTED identifier — one starting with an uppercase
// letter — from a genuinely separate package, and confirm exactly what
// "exported" means: reachable from outside the package that declares it.
package main

import (
	"fmt"

	"go-engineering/level-01-fundamentals/54-exported-identifiers/stats"
)

func main() {
	fmt.Println("=== Exported Identifiers ===")
	fmt.Println("----------------------------------")

	nums := []int{60, 75, 40, 90}

	// stats.Average and stats.Threshold are usable here ONLY because
	// their names start with an uppercase letter — this file is in
	// package main, a COMPLETELY different package from stats.
	avg := stats.Average(nums)
	fmt.Printf("stats.Average(%v) = %.1f\n", nums, avg)
	fmt.Printf("stats.Threshold   = %.1f\n", stats.Threshold)

	if avg >= stats.Threshold {
		fmt.Println("Average meets the exported threshold.")
	} else {
		fmt.Println("Average is below the exported threshold.")
	}
}
