package main

import "testing"

// TestIsPalindrome uses a table-driven test: a slice of input/expected-output
// pairs, looped over with t.Run so each case shows up individually in
// `go test -v` output. This is the idiomatic way to test many cases in Go.
func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"empty string", "", true},
		{"single char", "a", true},
		{"simple palindrome", "level", true},
		{"racecar", "racecar", true},
		{"not a palindrome", "gopher", false},
		{"two different chars", "go", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isPalindrome(tc.in)
			if got != tc.want {
				t.Errorf("isPalindrome(%q) = %t, want %t", tc.in, got, tc.want)
			}
		})
	}
}

// BenchmarkIsPalindrome demonstrates a Go benchmark — run with:
//
//	go test -bench=. -run=^$
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("racecar")
	}
}
