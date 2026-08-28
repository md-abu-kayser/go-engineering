// Lesson 57: Checksum Database Overview
//
// Goal: Understand what go.sum actually protects against, and how the
// public checksum database (sum.golang.org) backs it up independently.
package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("=== Checksum Database Overview ===")
	fmt.Println("----------------------------------")

	out, err := exec.Command("go", "env", "GOSUMDB", "GOPRIVATE", "GONOSUMCHECK").Output()
	if err != nil {
		fmt.Println("could not run `go env`:", err)
		return
	}
	lines := strings.Split(strings.TrimSuffix(string(out), "\n"), "\n")
	labels := []string{"GOSUMDB", "GOPRIVATE", "GONOSUMCHECK"}
	for i, line := range lines {
		if i < len(labels) {
			fmt.Printf("%-14s: %s\n", labels[i], line)
		}
	}

	fmt.Println("\nThis repository has no go.sum yet, since it has zero external")
	fmt.Println("dependencies. See the README for what go.sum records once it does,")
	fmt.Println("and how GOSUMDB independently verifies it.")
}
