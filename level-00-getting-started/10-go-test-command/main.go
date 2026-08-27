// Lesson 10: go test
//
// Goal: Write a small function with real logic, then test it with a
// table-driven test — the idiomatic Go testing pattern — and explore the
// `go test` command's most useful flags.
package main

import "fmt"

// isPalindrome reports whether s reads the same forwards and backwards.
// It's simple on purpose — the point of this lesson is the *testing*,
// not the algorithm.
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"level", "gopher", "racecar", "go"}
	for _, w := range words {
		fmt.Printf("%-10s isPalindrome = %t\n", w, isPalindrome(w))
	}
	fmt.Println("\nRun `go test -v ./...` in this folder to see the table-driven tests.")
}
