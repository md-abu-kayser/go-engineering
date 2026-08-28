// Lesson 46: Relative Paths
//
// Goal: Build and interpret relative paths correctly and portably using
// path/filepath, instead of hand-concatenating strings with "/".
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("=== Relative Paths ===")
	fmt.Println("----------------------------------")

	// filepath.Join builds a path using the CORRECT separator for the
	// current OS ("/" on Linux/macOS, "\" on Windows) and cleans up any
	// redundant slashes or ".." segments along the way.
	joined := filepath.Join("data", "reports", "2026", "summary.csv")
	fmt.Printf("filepath.Join(...)        : %s\n", joined)

	messy := filepath.Join("data//reports/../reports", "./summary.csv")
	fmt.Printf("filepath.Join (messy in)  : %s\n", messy)

	// filepath.Rel computes a relative path FROM one location TO another.
	rel, err := filepath.Rel("/home/gopher/project", "/home/gopher/project/data/reports")
	if err == nil {
		fmt.Printf("filepath.Rel(...)         : %s\n", rel)
	}

	// filepath.Dir and filepath.Base split a path into its directory and
	// final element.
	fmt.Printf("filepath.Dir(joined)      : %s\n", filepath.Dir(joined))
	fmt.Printf("filepath.Base(joined)     : %s\n", filepath.Base(joined))
	fmt.Printf("filepath.Ext(joined)      : %s\n", filepath.Ext(joined))
}
