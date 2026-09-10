// Lesson 53: Package Scope
//
// Goal: See that package-level declarations are visible EVERYWHERE in
// the package — every function, without needing to pass them as
// arguments or import anything — contrasted directly with the
// function-local scope from lesson 51.
package main

import "fmt"

// appName is declared at PACKAGE scope — outside any function. It's
// visible to EVERY function in this file (and, in a real multi-file
// package, every OTHER file in the same package too — see below).
var appName = "GO-ENGINEERING"

// requestCount is also package-scoped — multiple functions can read
// AND modify it directly, with no need to pass it around as a
// parameter or return value.
var requestCount int

// logRequest and resetCount both use requestCount directly — neither
// receives it as a parameter. This is ONLY possible because
// requestCount lives at package scope, not function scope.
func logRequest(path string) {
	requestCount++
	fmt.Printf("  [%s] request #%d to %s\n", appName, requestCount, path)
}

func resetCount() {
	requestCount = 0
}

func main() {
	fmt.Println("=== Package Scope ===")
	fmt.Println("----------------------------------")
	fmt.Printf("appName is visible right here too, in main(): %s\n", appName)

	fmt.Println("\n--- Multiple functions sharing package-scoped state ---")
	logRequest("/home")
	logRequest("/about")
	logRequest("/contact")
	fmt.Printf("requestCount after 3 requests: %d\n", requestCount)

	resetCount()
	fmt.Printf("requestCount after resetCount(): %d\n", requestCount)

	fmt.Println("\nSee the README: this file's package-level declarations would ALSO be")
	fmt.Println("visible from a second .go file, if this package had one — no import needed.")
}
