package main

import "testing"

// sumSlice adds up a []int using range — a small, testable helper
// pulled out specifically so this lesson has real logic to verify,
// rather than only testing printed output.
func sumSlice(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// countRunes counts the actual Unicode characters in s using range —
// contrasted with len(s), which counts bytes (lesson 08/09).
func countRunes(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func TestSumSlice(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"several elements", []int{1, 2, 3, 4}, 10},
		{"nil slice", nil, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := sumSlice(tc.in); got != tc.want {
				t.Errorf("sumSlice(%v) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}

func TestCountRunes(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int
	}{
		{"ascii only", "hello", 5},
		{"empty string", "", 0},
		{"multi-byte characters", "Hi, 世界", 6},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := countRunes(tc.in); got != tc.want {
				t.Errorf("countRunes(%q) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
