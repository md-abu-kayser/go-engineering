// Lesson 55: Build Metadata
//
// Goal: Embed build-time metadata (version, commit, build date) into a
// binary using -ldflags "-X ...", with sensible defaults for plain
// `go run`/`go build` when no metadata is injected.
package main

import "fmt"

// These three variables are deliberately package-level, exported-looking
// (though unexported here since main doesn't need to export them), with
// safe defaults. -ldflags "-X main.version=..." overwrites them at BUILD
// time, before main() ever runs — there's no runtime cost or complexity.
var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
)

func main() {
	fmt.Println("=== Build Metadata ===")
	fmt.Println("----------------------------------")
	fmt.Printf("version   : %s\n", version)
	fmt.Printf("commit    : %s\n", commit)
	fmt.Printf("buildDate : %s\n", buildDate)
	fmt.Println()
	fmt.Println("See the README for the -ldflags -X command that overrides these at")
	fmt.Println("build time — a real release build would wire this into `go build`.")
}
