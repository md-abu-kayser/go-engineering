//go:build windows

package main

// platformName is defined here only when building for GOOS=windows.
// See platform_other.go for the complementary definition.
func platformName() string {
	return "Windows (compiled in via //go:build windows)"
}
