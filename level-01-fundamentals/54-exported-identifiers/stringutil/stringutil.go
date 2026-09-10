// Package stringutil exists specifically to demonstrate exported vs
// unexported identifiers with a REAL, separately-imported package —
// not just a same-file example.
package stringutil

import "strings"

// Shout is EXPORTED — its name starts with an uppercase letter, so
// code in OTHER packages (like this lesson's main.go) can call it via
// stringutil.Shout(...).
func Shout(s string) string {
	return strings.ToUpper(s) + normalizedSuffix()
}

// normalizedSuffix is UNEXPORTED — lowercase first letter. It's a
// perfectly ordinary function, usable freely WITHIN this package
// (Shout calls it directly, right above), but main.go in the
// importing lesson package CANNOT reach it at all — not even by name.
func normalizedSuffix() string {
	return "!"
}

// DefaultGreeting is an EXPORTED package-level variable.
var DefaultGreeting = "Hello"

// internalVersion is an UNEXPORTED package-level variable — invisible
// outside this package, exactly like normalizedSuffix above.
var internalVersion = "v1"
