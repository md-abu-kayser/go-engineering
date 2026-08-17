// Lesson 01: Install & Verify Go
//
// Goal: Prove that Go is installed correctly by asking the Go runtime
// itself for information about the environment it's running in — instead
// of just trusting `go version` on the command line.
package main

import (
	"fmt"
	"runtime"
)

// envInfo holds a snapshot of the current Go environment.
// Using a small struct (instead of loose variables) keeps the data
// together and makes it trivial to test — see the _test.go file
// next to this one.
type envInfo struct {
	GoVersion string // e.g. "go1.22.0"
	OS        string // e.g. "linux", "windows", "darwin"
	Arch      string // e.g. "amd64", "arm64"
	NumCPU    int    // number of logical CPUs available to the Go runtime
	Compiler  string // "gc" (standard) or "gccgo"
}

// collectEnvInfo asks the runtime package for details about the current
// Go installation and machine. It's separated from main() so it can be
// tested independently of any printing.
func collectEnvInfo() envInfo {
	return envInfo{
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		NumCPU:    runtime.NumCPU(),
		Compiler:  runtime.Compiler,
	}
}

func main() {
	info := collectEnvInfo()

	fmt.Println("✅ Go is installed and working!")
	fmt.Println("----------------------------------")
	fmt.Printf("Go version : %s\n", info.GoVersion)
	fmt.Printf("Operating system : %s\n", info.OS)
	fmt.Printf("Architecture : %s\n", info.Arch)
	fmt.Printf("Logical CPUs : %d\n", info.NumCPU)
	fmt.Printf("Compiler : %s\n", info.Compiler)
}
