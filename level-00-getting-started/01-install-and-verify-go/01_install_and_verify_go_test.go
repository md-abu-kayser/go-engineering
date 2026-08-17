package main

import (
	"runtime"
	"strings"
	"testing"
)

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

