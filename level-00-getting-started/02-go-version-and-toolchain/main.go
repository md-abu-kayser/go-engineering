// Lesson 02: Go Version & Toolchain
//
// Goal: Understand what the "Go toolchain" actually is by inspecting build
// information from inside a running program — not just reading `go.mod`.
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	fmt.Println("=== Go Toolchain Info ===")
	fmt.Println("----------------------------------")
	fmt.Printf("Go version in use : %s\n", runtime.Version())
	fmt.Printf("Target OS/Arch : %s/%s\n", runtime.GOOS, runtime.GOARCH)

	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("\nNo build info available for this binary.")
		return
	}

	fmt.Println("\n=== Build Info (via runtime/debug) ===")
	fmt.Printf("Main module path : %s\n", info.Main.Path)
	fmt.Printf("Go version recorded at build : %s\n", info.GoVersion)

	if len(info.Deps) == 0 {
		fmt.Println("Dependencies : none — this lesson only uses the standard library")
	} else {
		fmt.Println("Dependencies:")
		for _, dep := range info.Deps {
			fmt.Printf("  - %s@%s\n", dep.Path, dep.Version)
		}
	}
}
