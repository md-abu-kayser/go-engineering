// Lesson 19: Test Cleanup
//
// Goal: Put behavior in a small pure function so a table-driven test,
// benchmark, fuzz target, or example test can exercise it deterministically.
package main

import (
	"fmt"
	"strings"
)

const lesson = "Test Cleanup"

func normalizeName(input string) (string, error) {
	name := strings.TrimSpace(input)
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	return strings.ToLower(name), nil
}

func main() {
	name, err := normalizeName("  Asha  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("normalized value: %s\n", name)
}