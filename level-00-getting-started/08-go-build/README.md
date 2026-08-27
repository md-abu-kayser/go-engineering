# 08 — go build

## 🎯 Learning Objectives

- Use the most common `go build` flags: `-o`, `-v`, `-race`, `-ldflags`.
- Cross-compile a Go binary for a different OS/architecture without leaving your machine.
- Understand why Go binaries are self-contained and easy to ship.

## 📖 Concept

Lesson 07 introduced `go build` at a basic level. This lesson focuses on the flags and
capabilities that make it genuinely useful in real projects.

### Useful flags

| Flag | Purpose |
|---|---|
| `-o <path>` | Choose the output binary's name/location instead of the default. |
| `-v` | Print the names of packages as they're compiled — useful for larger builds. |
| `-race` | Compile with the **race detector** enabled — catches concurrent data races at runtime. |
| `-ldflags "-s -w"` | Strip debug symbols to produce a smaller binary. |
| `-ldflags "-X pkg.Var=value"` | Inject a value into a variable at build time (common for embedding a version string). |

```bash
go build -o bin/app .
go build -v ./...
go build -race .
go build -ldflags "-X main.version=1.0.0" .
```

### Cross-compilation

Go can compile for a different operating system or CPU architecture than the one you're running
on, with **no extra toolchain to install** — just two environment variables:

```bash
GOOS=windows GOARCH=amd64 go build -o app.exe .
GOOS=linux   GOARCH=arm64 go build -o app-linux-arm64 .
GOOS=darwin  GOARCH=arm64 go build -o app-macos-arm64 .
```

This works because the Go compiler ships with support for every platform Go targets — you don't
need a Windows machine to produce a Windows binary. This is one of the biggest practical reasons
Go is popular for CLI tools and deployment artifacts.

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Printf("This binary was compiled for: %s/%s\n", runtime.GOOS, runtime.GOARCH)
```

`runtime.GOOS` and `runtime.GOARCH` report the platform the **running binary** was built for —
these are baked in at compile time, which is exactly what makes the cross-compilation experiment
below interesting.

## ▶️ How to Run

```bash
cd level-00-getting-started/08-go-build
go build -o bin/app .
./bin/app
```

## ✅ Expected Output

```
=== go build, in depth ===
----------------------------------
This binary was compiled for: linux/amd64

Try cross-compiling this file for another platform (see README).
```

## 🧠 Key Takeaways

- `go build` flags like `-o`, `-race`, and `-ldflags` cover the vast majority of real needs.
- `GOOS`/`GOARCH` environment variables enable cross-compilation with zero extra tooling.
- `runtime.GOOS`/`runtime.GOARCH` reflect the target the binary was actually built for.

## 🛠️ Try It Yourself

1. Cross-compile this lesson for Windows: `GOOS=windows GOARCH=amd64 go build -o app.exe .`
2. Check the resulting file with `file app.exe` (Linux/macOS) — note it's correctly identified
   as a Windows PE executable, even though you built it on a different OS.
3. Build with `-ldflags "-s -w"` and compare the binary size to a normal build.

## ⚠️ Common Mistakes

- Forgetting that cross-compiled binaries **cannot run** on the machine that built them (a
  Windows `.exe` won't execute on Linux) — they're for the *target* platform.
- Using `-race` in production builds — it adds real runtime overhead and is meant for testing.
