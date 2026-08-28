package main

import "testing"

// TestDefaultMetadata confirms the package-level defaults are exactly
// what an un-injected `go run`/`go build`/`go test` sees — the same
// values a plain `go test` (with no -ldflags) will always exercise.
func TestDefaultMetadata(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"version", version, "dev"},
		{"commit", commit, "none"},
		{"buildDate", buildDate, "unknown"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Errorf("%s = %q, want %q (did you build with -ldflags -X?)", tc.name, tc.got, tc.want)
			}
		})
	}
}
