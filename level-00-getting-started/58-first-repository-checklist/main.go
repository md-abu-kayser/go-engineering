// Lesson 58: First Repository Checklist
//
// Goal: A capstone check — this program mechanically verifies several
// of the checklist items from the README against THIS lesson's own
// folder, so "does my repo pass the checklist" has at least a partial,
// automatable answer.
package main

import (
	"fmt"
	"os"
)

// checkFile reports whether name exists in the current directory.
func checkFile(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

func main() {
	fmt.Println("=== First Repository Checklist (partial, automated check) ===")
	fmt.Println("----------------------------------")

	checks := []struct {
		label string
		ok    bool
	}{
		{"README.md exists in this folder", checkFile("README.md")},
		{"main.go exists in this folder", checkFile("main.go")},
	}

	allPassed := true
	for _, c := range checks {
		status := "✅"
		if !c.ok {
			status = "❌"
			allPassed = false
		}
		fmt.Printf("%s %s\n", status, c.label)
	}

	fmt.Println("\nThe REST of the checklist (LICENSE, .gitignore, gofmt/vet/test cleanliness,")
	fmt.Println("a clear README structure) isn't mechanically checkable from inside one lesson")
	fmt.Println("folder — see the README for the full list to apply at the repository root.")

	if !allPassed {
		os.Exit(1)
	}
}
