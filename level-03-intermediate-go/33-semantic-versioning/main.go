// Lesson 33: Semantic Versioning
//
// Goal: Model an explicit dependency boundary. Real module and build-tool
// commands belong outside application code; the application consumes an
// already-validated configuration value.
package main

import "fmt"

const lesson = "Semantic Versioning"

type dependency struct {
	Path    string
	Version string
}

func (d dependency) Valid() bool {
	return d.Path != "" && d.Version != ""
}

func main() {
	dep := dependency{Path: "example.com/catalog", Version: "v1.4.0"}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("%s at %s is valid: %t\n", dep.Path, dep.Version, dep.Valid())
}