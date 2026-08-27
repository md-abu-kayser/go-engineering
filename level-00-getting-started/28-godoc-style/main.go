// Lesson 28: godoc Style
//
// Goal: See a small, complete package documented consistently — package
// overview, zero-value behavior, every exported symbol — and read it
// back the same way any consumer of the package would.
package main

import (
	"fmt"

	"go-engineering/level-00-getting-started/28-godoc-style/kvstore"
)

func main() {
	fmt.Println("=== godoc Style ===")
	fmt.Println("----------------------------------")

	s := kvstore.New()
	s.Set("name", "Gopher")
	s.Set("language", "Go")

	if value, ok := s.Get("name"); ok {
		fmt.Printf("name = %s\n", value)
	}
	fmt.Printf("stored keys: %d\n", s.Len())

	s.Delete("language")
	fmt.Printf("stored keys after delete: %d\n", s.Len())

	fmt.Println("\nRun `go doc ./kvstore` and `go doc -all ./kvstore` to read this")
	fmt.Println("package exactly the way a new consumer of it would.")
}
