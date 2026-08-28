# 48 — Cross-Compilation Overview

## 🎯 Learning Objectives

- List every `GOOS`/`GOARCH` combination Go supports, directly from your toolchain.
- Cross-compile for several real targets and verify each binary's platform.
- Understand `CGO_ENABLED` and why it matters for reliable cross-compilation.

## 📖 Concept

[Lesson 08](../08-go-build) introduced cross-compilation briefly. This lesson is a dedicated,
deeper look — because it's one of Go's most genuinely useful features for shipping software.

### Listing every supported target

```bash
go tool dist list
```

This prints every `GOOS/GOARCH` pair your installed Go toolchain can build for — commonly over
40 combinations, including things like `linux/amd64`, `windows/arm64`, `darwin/arm64`,
`js/wasm`, and more. No extra downloads needed; support for every listed target ships with a
standard Go installation.

### A practical matrix

| GOOS | GOARCH | Produces a binary for |
|---|---|---|
| `linux` | `amd64` | Most servers, most desktop Linux |
| `linux` | `arm64` | Raspberry Pi (64-bit), AWS Graviton, many modern ARM servers |
| `windows` | `amd64` | Most Windows PCs |
| `darwin` | `amd64` | Older/Intel Macs |
| `darwin` | `arm64` | Apple Silicon Macs (M1/M2/M3/…) |

```bash
GOOS=linux   GOARCH=arm64 go build -o app-linux-arm64   .
GOOS=windows GOARCH=amd64 go build -o app-windows.exe   .
GOOS=darwin  GOARCH=arm64 go build -o app-macos-arm64   .
```

### `CGO_ENABLED` — the thing that quietly breaks cross-compilation

If your code (or a dependency) uses `cgo` (calling into C code), cross-compilation can fail or
silently produce a binary that only works on the exact machine it was built on, because `cgo`
needs a real C compiler for the **target** platform, not just Go's own toolchain. The fix, when
your code doesn't genuinely need `cgo`:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app .
```

`CGO_ENABLED=0` produces a fully **statically linked** binary with no C dependency at all — this
is exactly why most Go Docker images can be `FROM scratch` (an empty base image) and still work:
the binary needs nothing else.

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Printf("This binary targets: GOOS=%s GOARCH=%s\n", runtime.GOOS, runtime.GOARCH)
```

Exactly like [lesson 08](../08-go-build), this reports the platform baked into **this specific
binary** at build time — build it several different ways and compare.

## ▶️ How to Run

```bash
cd level-00-getting-started/48-cross-compilation-overview
go run main.go
go tool dist list | head -20
```

## ✅ Expected Output (normal run)

```
=== Cross-Compilation Overview ===
----------------------------------
This binary targets: GOOS=linux GOARCH=amd64

See the README for the full GOOS/GOARCH matrix, and how to list it yourself
with `go tool dist list`.
```

## 🧠 Key Takeaways

- `go tool dist list` shows every target your Go toolchain can cross-compile for, out of the box.
- Cross-compiling is just two environment variables (`GOOS`, `GOARCH`) plus `go build`.
- `CGO_ENABLED=0` avoids the most common cross-compilation failure mode, at the cost of not being
  able to use `cgo`-dependent packages.
- Statically-linked, `CGO_ENABLED=0` binaries are why minimal Docker images (`FROM scratch`) work
  for Go programs.

## 🛠️ Try It Yourself

1. Run `go tool dist list` and count how many total `GOOS/GOARCH` combinations your Go
   installation supports.
2. Cross-compile this lesson for `windows/amd64` and `darwin/arm64`, and use the `file` command
   (Linux/macOS) to confirm each binary is correctly identified as belonging to its target platform.
3. Look up whether any package you've used elsewhere in this repository uses `cgo` (hint: none of
   them do — that's exactly why every lesson so far has cross-compiled without any extra flags).

## ⚠️ Common Mistakes

- Assuming a cross-compiled binary can be tested on the machine that built it — it generally
  cannot; a `windows/amd64` binary won't run on Linux, and vice versa.
- Forgetting `CGO_ENABLED=0` when a dependency uses `cgo` and cross-compilation either fails
  outright or produces a binary that mysteriously only works on the build machine.
