// Package stats is a small, REAL separate package (not just a
// function in main.go) — specifically so this lesson can demonstrate
// exported identifiers being used from OUTSIDE the package that
// declares them, which is the whole point of exporting something.
package stats

// Average is EXPORTED — its name starts with an uppercase letter, so
// any OTHER package that imports "stats" can call it directly, exactly
// like main.go does below.
func Average(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	total := 0
	for _, n := range nums {
		total += n
	}
	return float64(total) / float64(len(nums))
}

// Threshold is an EXPORTED package-level variable — also reachable
// from outside this package, as stats.Threshold.
var Threshold = 50.0
