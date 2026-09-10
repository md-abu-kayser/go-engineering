// Lesson 56: Doc Comments
//
// Goal: See specifically WHICH identifiers a real Go codebase expects
// doc comments on, tied directly to lessons 54/55's exported vs.
// unexported distinction — not a formatting deep-dive (that's already
// covered by level 00's lessons 13, 27, and 28).
package main

import "fmt"

// Temperature represents a temperature reading in Celsius.
//
// This is EXPORTED (capital T), so it has a doc comment — anyone
// outside this package relying on Temperature deserves documentation,
// since they can't just "go read the unexported internals" to
// understand it themselves.
type Temperature struct {
	// Celsius is the EXPORTED field holding the reading.
	// It also gets a comment, for the same reason as the type itself.
	Celsius float64

	// source is UNEXPORTED — an internal implementation detail nobody
	// outside this package can even see, let alone needs documented
	// for THEM. A brief note here is still useful for future
	// maintainers of THIS package, but it's a different AUDIENCE than
	// Celsius's comment above.
	source string
}

// Fahrenheit converts t to degrees Fahrenheit. EXPORTED — documented.
func (t Temperature) Fahrenheit() float64 {
	return t.Celsius*9/5 + 32
}

// clampToRealistic is UNEXPORTED — no outside caller will ever see
// this function's name, so there's no PUBLIC API contract to document
// here. A short comment is still genuinely helpful for anyone editing
// this file later, but it need not follow the "start with the name"
// convention as strictly, since go doc won't even show it to external
// consumers of this package.
func clampToRealistic(c float64) float64 {
	if c < -273.15 { // absolute zero
		return -273.15
	}
	return c
}

func main() {
	fmt.Println("=== Doc Comments ===")
	fmt.Println("----------------------------------")

	t := Temperature{Celsius: clampToRealistic(21.5), source: "sensor-1"}
	fmt.Printf("%.1f°C = %.1f°F\n", t.Celsius, t.Fahrenheit())

	fmt.Println("\nRun `go doc .` and `go doc -all .` here — notice unexported")
	fmt.Println("identifiers like `source` and `clampToRealistic` never appear.")
}
