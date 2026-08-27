// Lesson 19: Semantic Import Versioning
//
// Goal: Understand how Go encodes major version numbers directly into
// import paths, and why that's a deliberate design choice.
package main

import "fmt"

func main() {
	examples := []struct {
		version    string
		importPath string
	}{
		{"v0.x.x or v1.x.x", "github.com/example/project"},
		{"v2.x.x", "github.com/example/project/v2"},
		{"v3.x.x", "github.com/example/project/v3"},
	}

	fmt.Println("=== Semantic Import Versioning ===")
	fmt.Println("----------------------------------")
	for _, ex := range examples {
		fmt.Printf("%-16s -> %s\n", ex.version, ex.importPath)
	}
	fmt.Println("\nNotice: only v2 and above change the import path itself.")
}
