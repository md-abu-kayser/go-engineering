// Lesson 37: Runtime Panics
//
// Goal: Understand what a panic actually is, the handful of common ways
// one happens, and how `defer` + `recover` turns a panic into an
// ordinary, handleable error instead of a crash.
package main

import "fmt"

// riskyDivide panics if b is zero, instead of returning an error — a
// deliberately "bad citizen" function so this lesson has something real
// to recover from below.
func riskyDivide(a, b int) int {
	return a / b // panics: "runtime error: integer divide by zero" if b == 0
}

// safeDivide wraps riskyDivide with a deferred recover, converting a
// panic into an ordinary error return instead of crashing the whole
// program.
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	result = riskyDivide(a, b)
	return result, nil
}

func main() {
	if result, err := safeDivide(10, 2); err == nil {
		fmt.Println("10 / 2 =", result)
	}

	if _, err := safeDivide(10, 0); err != nil {
		fmt.Println("Handled gracefully:", err)
	}

	fmt.Println("\nSee the README for the most common causes of a Go panic, and why")
	fmt.Println("recover() only works when called directly inside a deferred function.")
}
