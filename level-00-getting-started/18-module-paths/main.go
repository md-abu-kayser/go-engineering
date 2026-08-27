// Lesson 18: Module Paths
//
// Goal: Understand what a module path is, how import paths are derived
// from it, and how to read one at a glance.
package main

import "fmt"

func main() {
	modulePath := "go-engineering"
	packageDir := "level-00-getting-started/18-module-paths"
	fullImportPath := modulePath + "/" + packageDir

	fmt.Println("=== Module Paths ===")
	fmt.Println("----------------------------------")
	fmt.Printf("Module path            : %s\n", modulePath)
	fmt.Printf("Package's folder       : %s\n", packageDir)
	fmt.Printf("Package's import path  : %s\n", fullImportPath)
	fmt.Println()
	fmt.Println("Every package's import path = <module path> + <its folder path>.")
}
