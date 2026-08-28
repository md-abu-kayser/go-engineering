// Lesson 07: Booleans
//
// Goal: Use bool, the three logical operators, and observe Go's
// short-circuit evaluation directly.
package main

import "fmt"

// hasSideEffect is called only if short-circuit evaluation DOESN'T skip
// it — used below to make short-circuiting visible, not just theoretical.
func hasSideEffect(label string, value bool) bool {
	fmt.Printf("  (evaluating %s)\n", label)
	return value
}

func main() {
	fmt.Println("=== Booleans ===")
	fmt.Println("----------------------------------")

	isRaining := true
	haveUmbrella := false

	fmt.Printf("isRaining              : %t\n", isRaining)
	fmt.Printf("haveUmbrella           : %t\n", haveUmbrella)
	fmt.Printf("!isRaining             : %t\n", !isRaining)
	fmt.Printf("isRaining && haveUmbrella : %t\n", isRaining && haveUmbrella)
	fmt.Printf("isRaining || haveUmbrella : %t\n", isRaining || haveUmbrella)

	fmt.Println("\n--- Short-circuit evaluation ---")
	fmt.Println("false && hasSideEffect(...):")
	result := false && hasSideEffect("right side of &&", true)
	fmt.Printf("  result = %t (right side was NEVER evaluated — see above, no line printed)\n", result)

	fmt.Println("true || hasSideEffect(...):")
	result = true || hasSideEffect("right side of ||", true)
	fmt.Printf("  result = %t (right side was NEVER evaluated — see above, no line printed)\n", result)

	fmt.Println("true && hasSideEffect(...):")
	result = true && hasSideEffect("right side of && (this time it DOES run)", true)
	fmt.Printf("  result = %t\n", result)
}
