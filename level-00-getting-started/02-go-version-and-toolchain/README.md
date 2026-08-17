# 02 — Go Version & Toolchain

## 🎯 Learning Objectives

- Understand what the **Go toolchain** is (compiler, linker, formatter, module tooling — all
  bundled together).
- Understand what `go.mod` does and what the `go 1.22` directive means.
- Learn the difference between the Go version **installed on your machine** and the Go version
  a **module requires**.
- Use `runtime/debug.ReadBuildInfo()` to inspect a program's build metadata from inside itself.

## 📖 Concept

When people say "the Go toolchain," they mean the full set of tools that ship with a Go
installation: the compiler, the linker, `gofmt`, the module downloader, the test runner, and
more — all accessible through the single `go` command (`go build`, `go run`, `go test`,
`go fmt`, `go mod`, …).

### `go.mod` — the module manifest

At the root of this repository is a `go.mod` file:

```go
module go-engineering

go 1.22
```

- **`module go-engineering`** — declares the module's import path. Every package inside this
  repo is addressed relative to this.
- **`go 1.22`** — declares the **minimum** Go language version this module requires. It's not
  the version you must have installed — Go will use the version you have, as long as it's new
  enough, and modern Go (1.21+) can even **auto-download** a newer toolchain if your `go.mod`
  requires one you don't have (controlled by the `GOTOOLCHAIN` environment variable).

### Checking versions from the command line

```bash
go version          # the Go version installed on this machine
go env GOVERSION     # same info, script-friendly
go env GOTOOLCHAIN   # toolchain selection policy (e.g. "auto")
```

## 🔍 Code Walkthrough (`main.go`)

```go
info, ok := debug.ReadBuildInfo()
```

This is a **two-value return** — a pattern you'll see constantly in Go. `debug.ReadBuildInfo()`
returns the build metadata _and_ a boolean telling you whether it succeeded. Checking `ok` before
using `info` avoids working with a zero-value struct that doesn't mean anything.

```go
for _, dep := range info.Deps {
    fmt.Printf("  - %s@%s\n", dep.Path, dep.Version)
}
```

A `for range` loop over a slice. The blank identifier `_` throws away the index because we only
care about each dependency's value here — more on `_` in [lesson 05](../05-imports).

## ▶️ How to Run

```bash
cd level-00-getting-started/02-go-version-and-toolchain
go run main.go
```

## ✅ Expected Output

With `go build && ./02-go-version-and-toolchain`:

```
=== Go Toolchain Info ===
----------------------------------
Go version in use : go1.22.2
Target OS/Arch : linux/amd64

=== Build Info (via runtime/debug) ===
Main module path : go-engineering
Go version recorded at build : go1.22.2
Dependencies : none — this lesson only uses the standard library
```

> Note: with `go run` instead of `go build`, `Main module path` commonly comes back **empty**
> (or as `command-line-arguments` on some Go versions), because `go run` compiles into a
> throwaway temp build rather than a "real" build of your module. Try both and compare —
> `go build` always gives you the complete picture.

## 🧠 Key Takeaways

- The **toolchain** is the whole `go` command suite, not just the compiler.
- `go.mod` declares your module's name and its **minimum** required Go version.
- `runtime/debug.ReadBuildInfo()` lets a program inspect its own build metadata.
- Two-value returns (`value, ok`) are idiomatic Go for "did this succeed?"

## 🛠️ Try It Yourself

1. Run `go env` with no arguments and skim the full list of environment variables Go uses.
2. Change `go 1.22` in the root `go.mod` to a version newer than what you have installed, then
   run `go run main.go` again and read the error message carefully.
3. Revert your change afterward so the rest of the repo keeps working.

## ⚠️ Common Mistakes

- Confusing the **installed** Go version with the **required** version in `go.mod` — they are
  independent numbers that both matter.
- Forgetting that `go.mod` lives at the **module root**, not inside every subfolder.
