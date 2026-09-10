package main

import "testing"

// TestClassifyScore exercises classifyScore's actual behavior — an
// indirect but concrete way to confirm the function's INTERNAL block
// scoping (the `bonus` variable inside the if-block) produces the
// correct externally-visible result in every branch.
func TestClassifyScore(t *testing.T) {
	cases := []struct {
		name  string
		score int
		want  string
	}{
		{"high honors score", 95, "A (with honors)"},
		{"exactly the honors threshold", 90, "A (with honors)"},
		{"mid-range score", 75, "B"},
		{"just below B threshold", 69, "unknown"},
		{"low score", 40, "unknown"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := classifyScore(tc.score); got != tc.want {
				t.Errorf("classifyScore(%d) = %q, want %q", tc.score, got, tc.want)
			}
		})
	}
}
