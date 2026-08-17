// This test file lives in the same package as main.go, which means it can
// call collectEnvInfo() directly even though that function is unexported
// (lowercase). This is a common Go pattern: tests live right next to the
// code they test, in the same package.
package main

import (
	"runtime"
	"strings"
	"testing"
)

// TestCollectEnvInfo checks that the values returned by collectEnvInfo()
// line up with what the runtime package reports directly. This is mostly
// a sanity check that our wrapper struct isn't losing or mangling data.
func TestCollectEnvInfo(t *testing.T) {
	info := collectEnvInfo()

	if info.GoVersion != runtime.Version() {
		t.Errorf("GoVersion = %q, want %q", info.GoVersion, runtime.Version())
	}

	if !strings.HasPrefix(info.GoVersion, "go") {
		t.Errorf("GoVersion = %q, expected it to start with %q", info.GoVersion, "go")
	}

	if info.OS != runtime.GOOS {
		t.Errorf("OS = %q, want %q", info.OS, runtime.GOOS)
	}

	if info.Arch != runtime.GOARCH {
		t.Errorf("Arch = %q, want %q", info.Arch, runtime.GOARCH)
	}

	if info.NumCPU < 1 {
		t.Errorf("NumCPU = %d, want at least 1", info.NumCPU)
	}

	if info.Compiler != runtime.Compiler {
		t.Errorf("Compiler = %q, want %q", info.Compiler, runtime.Compiler)
	}
}

// Run this test with:
//
//	go test ./...
//
// or, for verbose output showing each test name:
//
//	go test -v ./...
