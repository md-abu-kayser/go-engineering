# 09 — go install

## 🎯 Learning Objectives

- Understand what `go install` does differently from `go build`.
- Understand `GOBIN` and how it relates to `GOPATH`.
- Install a specific **version** of a third-party Go tool directly, without cloning it.

## 📖 Concept

`go install` compiles a package **and installs the resulting binary** into a well-known
directory — `$GOBIN` if set, otherwise `$GOPATH/bin` (or `~/go/bin` if `GOPATH` isn't set
either). If that directory is on your shell's `PATH`, the tool becomes runnable by name from
anywhere on your system.

```bash
go install .                 # install the current directory's main package
go env GOBIN                 # see where it went
go env GOPATH                # the fallback location if GOBIN is unset
```

### Installing versioned tools directly

One of `go install`'s most useful tricks: you can install a specific version of any published
Go tool directly from its module path, without cloning the repository first:

```bash
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.0
```

The `@version` suffix (`@latest`, `@v1.2.3`, `@main`) tells Go exactly which version to fetch,
build, and install — this is how most Go command-line tools are distributed today, instead of
platform-specific installers.

## 🔍 Code Walkthrough (`main.go`)

```go
gobin := os.Getenv("GOBIN")
```

This reads the `GOBIN` **environment variable** directly from the shell. Note this may be empty
even when Go *does* have an effective `GOBIN` — Go falls back to a computed default
(`$GOPATH/bin`) when the variable isn't explicitly set. `go env GOBIN` (the command, not the
environment variable) always shows the effective value, which is why the program points you at
it instead of trying to compute the fallback logic itself.

## ▶️ How to Run

```bash
cd level-00-getting-started/09-go-install
go run main.go
go install .
```

## ✅ Expected Output (shape)

```
=== go install ===
----------------------------------
GOBIN: (not set — falls back to $GOPATH/bin, run `go env GOBIN` and `go env GOPATH`)

Run `go install` in this folder, then check the path above for a new binary.
```

## 🧠 Key Takeaways

- `go install` = build + place the binary somewhere on your `PATH`.
- `GOBIN` (if set) wins; otherwise Go falls back to `$GOPATH/bin`.
- `go install module/path/cmd/tool@version` installs a specific version without manual cloning.
- `go env GOBIN` / `go env GOPATH` always reflect the *effective* value, unlike reading the raw
  environment variable yourself.

## 🛠️ Try It Yourself

1. Run `go env GOBIN` and `go env GOPATH` and compare them to what `os.Getenv("GOBIN")` printed.
2. Install a real tool: `go install golang.org/x/tools/cmd/goimports@latest`, then run
   `goimports -h` (assuming the install directory is on your `PATH`).
3. Add `$(go env GOPATH)/bin` to your shell's `PATH` permanently if it isn't already — most Go
   CLI tools assume it is.

## ⚠️ Common Mistakes

- Installing a tool successfully but not being able to run it — almost always a `PATH` issue,
  not an installation failure.
- Confusing `GOPATH` (the whole Go workspace, legacy-ish today) with `GOBIN` (just the binary
  install location) — they're related but distinct settings.
