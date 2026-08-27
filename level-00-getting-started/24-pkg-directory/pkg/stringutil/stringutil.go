// Package stringutil provides small, reusable string helpers — exactly
// the kind of general-purpose, exported code the pkg/ convention is
// meant to hold: code intended to be imported, not just an
// implementation detail of one binary.
package stringutil

import "strings"

// Reverse returns s with its characters in reverse order.
// It operates on runes rather than raw bytes, so multi-byte UTF-8
// characters are reversed correctly instead of being corrupted.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// TitleCase capitalizes the first letter of every word in s.
func TitleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if w == "" {
			continue
		}
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, " ")
}
