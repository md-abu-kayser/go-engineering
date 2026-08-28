//go:build !windows

package main

// platformName is defined here for every target EXCEPT windows. Exactly
// one of this file and platform_windows.go is compiled into any given
// build — never both, never neither.
func platformName() string {
	return "Not Windows (compiled in via //go:build !windows)"
}
