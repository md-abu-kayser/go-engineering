# 07 — go run

## 🎯 Learning Objectives

- Explain what `go run` actually does, step by step.
- Distinguish `go run` from `go build` and `go install`.
- Know when to reach for each one in a real development workflow.
- Inspect `os.Args` to see how a program was invoked.

## 📖 Concept

All three commands ultimately do the same core thing — **compile your source code** — but they
differ in what happens to the result.

| Command          | Compiles? | Leaves a binary behind?           | Runs it?                                        | Typical use                                       |
| ---------------- | --------- | --------------------------------- | ----------------------------------------------- | ------------------------------------------------- |
| `go run main.go` | ✅        | ❌ (temp file, deleted after)     | ✅ immediately                                  | Fast iteration during development                 |
| `go build`       | ✅        | ✅ in the current directory       | ❌ (you run it yourself)                        | Producing a binary to test, ship, or containerize |
| `go install`     | ✅        | ✅ in `$GOBIN` (or `$GOPATH/bin`) | ❌ (you run it yourself, from anywhere on PATH) | Installing a CLI tool you'll use repeatedly       |

### `go run` — compile, execute, discard

```bash
go run main.go
```

Under the hood, Go compiles your file(s) into a **temporary binary** in a system temp directory,
executes that binary immediately, streams its output back to your terminal, and then deletes the
temporary binary. You never see the intermediate file. This is the fastest loop for "I changed
one line, let me see if it works," which is why it's the command you'll use constantly while
learning.

### `go build` — compile, keep the binary

```bash
go build
./07-go-run          # macOS/Linux
07-go-run.exe         # Windows
```

This produces a real, persistent executable named after the module or folder, in your current
directory. You then run it yourself, as many times as you like, with no recompilation. This is
what you'd do to produce an artifact for testing, Docker packaging, or distribution.

### `go install` — compile, install globally

```bash
go install
```

This behaves like `go build`, but instead of leaving the binary in the current folder, it places
it in `$GOBIN` (or `$GOPATH/bin` if `$GOBIN` isn't set). If that directory is on your system
`PATH`, the tool becomes runnable from **anywhere**, by name — exactly how you install Go-based
CLI tools (like `golangci-lint` or `air`) system-wide.

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Printf("This process was started as: %s\n", os.Args[0])
```

`os.Args` is a slice of the command-line arguments the running program was invoked with.
`os.Args[0]` is always the path to the program itself — comparing this value across `go run`,
`go build`, and `go install` is a concrete way to see the difference between them, rather than
just reading about it.

## ▶️ How to Run — try all three!

```bash
cd level-00-getting-started/07-go-run

# 1. go run
go run main.go

# 2. go build
go build
./07-go-run

# 3. go install (optional — installs into $GOBIN)
go install
```

## ✅ Expected Output (shape)

```
=== go run vs go build vs go install ===
----------------------------------
This process was started as: /tmp/go-buildXXXXXXXXX/b001/exe/main
...
```

With `go run`, `os.Args[0]` will point into a temporary build directory. With `go build`, it will
point at `./07-go-run` in your current folder. Compare them side by side.

## 🧠 Key Takeaways

- `go run` = compile + execute + discard the binary — best for fast iteration.
- `go build` = compile + keep the binary locally — best for producing an artifact.
- `go install` = compile + place the binary on your PATH — best for CLI tools you use often.
- `os.Args[0]` reveals exactly how (and from where) a running program was started.

## 🛠️ Try It Yourself

1. Run all three commands above and compare the printed path each time.
2. After `go build`, check the file size and permissions of the generated binary (`ls -la`).
3. Delete the binary you built (`rm 07-go-run` / `del 07-go-run.exe`) to keep the folder clean —
   built binaries are already excluded by the repo's root `.gitignore`.

## ⚠️ Common Mistakes

- Committing built binaries to version control — they're large, platform-specific, and
  regenerable, so they don't belong in git (this repo's `.gitignore` already excludes them).
- Expecting `go run` to leave a binary behind — it never does, by design.
- Forgetting `$GOBIN`/`$GOPATH/bin` isn't on `PATH` and then wondering why `go install` "did
  nothing" — it worked, the binary just isn't reachable by name yet.
