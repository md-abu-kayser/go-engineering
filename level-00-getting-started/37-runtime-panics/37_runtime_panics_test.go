package main

import (
	"strings"
	"testing"
)

// TestRiskyDivide_Panics confirms riskyDivide really does panic on
// division by zero, using recover() inside the TEST itself to catch it
// — otherwise this test would crash the whole test binary.
func TestRiskyDivide_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected riskyDivide(10, 0) to panic, but it did not")
		}
	}()
	riskyDivide(10, 0)
	t.Fatal("unreachable: riskyDivide should have panicked before this line")
}

// TestSafeDivide confirms the recover-based wrapper turns that same
// panic into an ordinary, checkable error instead.
func TestSafeDivide(t *testing.T) {
	if result, err := safeDivide(10, 2); err != nil || result != 5 {
		t.Errorf("safeDivide(10, 2) = %d, %v; want 5, <nil>", result, err)
	}

	_, err := safeDivide(10, 0)
	if err == nil {
		t.Fatal("expected an error from safeDivide(10, 0), got nil")
	}
	if !strings.Contains(err.Error(), "recovered from panic") {
		t.Errorf("error message = %q, want it to mention %q", err.Error(), "recovered from panic")
	}
}
