// Lesson 56: Module Cache Overview
//
// Goal: Locate Go's local module cache, understand its structure, and
// know the command for clearing it.
package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("=== Module Cache Overview ===")
	fmt.Println("----------------------------------")

	out, err := exec.Command("go", "env", "GOMODCACHE").Output()
	if err != nil {
		fmt.Println("could not run `go env GOMODCACHE`:", err)
		return
	}
	cacheDir := strings.TrimSpace(string(out))
	fmt.Printf("GOMODCACHE: %s\n", cacheDir)

	fmt.Println("\nThis repository currently depends on zero external modules, so this")
	fmt.Println("lesson's own module cache entry (if any) is just the standard library,")
	fmt.Println("which isn't cached the same way third-party modules are.")
	fmt.Println("\nSee the README for the cache's on-disk layout and how to clear it.")
}
