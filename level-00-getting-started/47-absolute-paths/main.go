// Lesson 47: Absolute Paths
//
// Goal: Convert a relative path to an absolute one, check whether a path
// is already absolute, and understand why programs that run from
// unpredictable working directories should prefer absolute paths.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== Absolute Paths ===")
	fmt.Println("----------------------------------")

	fmt.Printf("filepath.IsAbs(\"README.md\")     : %t\n", filepath.IsAbs("README.md"))
	fmt.Printf("filepath.IsAbs(\"/etc/hosts\")     : %t\n", filepath.IsAbs("/etc/hosts"))

	// filepath.Abs resolves a relative path against the CURRENT WORKING
	// DIRECTORY (lesson 45) and returns a clean, absolute result.
	abs, err := filepath.Abs("README.md")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving absolute path:", err)
		os.Exit(1)
	}
	fmt.Printf("filepath.Abs(\"README.md\")       : %s\n", abs)

	// An already-absolute path is returned unchanged (after cleaning).
	absAlready, _ := filepath.Abs("/etc/hosts")
	fmt.Printf("filepath.Abs(\"/etc/hosts\")       : %s\n", absAlready)
}
