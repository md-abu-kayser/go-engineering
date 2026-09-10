// This is a SECOND file in the SAME package as main.go. Everything
// declared here at package level is visible in main.go — and every
// other file in this package — with NO import needed. The package,
// not the file, is the real unit of scope for these declarations.
package main

// sharedGreeting is declared HERE, in helper.go...
var sharedGreeting = "Hello from helper.go's package-level variable"

// formatCount is ALSO declared here...
func formatCount(n int) string {
	if n == 1 {
		return "1 item"
	}
	return "many items"
}
