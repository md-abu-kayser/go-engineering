# 15 — go list

## 🎯 Learning Objectives

- Use `go list` to inspect packages, their imports, and module information.
- Use the `-m` flag to work with modules instead of packages.
- Use `-f` (format) and `-json` to extract exactly the information you need, script-friendly.

## 📖 Concept

`go list` is Go's introspection command — it answers questions about your code's structure
without you having to read every file by hand.

### Listing packages

```bash
go list ./...                       # every package import path in the current module
go list -f '{{.ImportPath}}: {{.GoFiles}}' ./...   # custom formatted output
go list -json ./level-00-getting-started/15-go-list-command   # full package metadata as JSON
```

`go list ./...` is one of the most-used Go commands in real projects — CI pipelines commonly use
it to discover every package to test, lint, or build.

### Listing module information

```bash
go list -m                # the current module's path and version
go list -m all             # the current module plus every dependency, with resolved versions
go list -m -versions <path> # every published version of a given module
```

### Inspecting one package's imports

```bash
go list -f '{{.Imports}}' ./level-00-getting-started/15-go-list-command
```

This prints exactly what packages this Go file imports — useful for understanding a codebase you
didn't write, or for scripting a dependency check.

## 🔍 Code Walkthrough (`main.go`)

This lesson's code is intentionally simple — a couple of standard library imports (`os`,
`strings`) so that `go list`'s import-inspection commands above have something real to report.
The **command line itself** is the actual lesson content here, more than the Go source.

## ▶️ How to Run

```bash
cd level-00-getting-started/15-go-list-command
go run main.go
go list -f '{{.ImportPath}}: {{.Imports}}' .
```

From the **repository root** (`GO-ENGINEERING/`), also try:

```bash
go list ./...
go list -m
```

## ✅ Expected Output (shape)

```
=== go list ===
----------------------------------
This program itself isn't the point of this lesson — try the
commands in the README against this repository instead.
/tmp/go-buildXXXX/b001/exe/main
```

And from the repo root, `go list ./...` prints one import path per lesson package, e.g.:

```
go-engineering/level-00-getting-started/01-install-and-verify-go
go-engineering/level-00-getting-started/02-go-version-and-toolchain
...
```

## 🧠 Key Takeaways

- `go list ./...` enumerates every package in a module — a CI/tooling staple.
- `go list -m [all]` works at the **module** level instead of the package level.
- `-f`/`-json` let you extract exactly the fields you need for scripting.
- `go list` reads metadata without compiling anything — it's fast even on large codebases.

## 🛠️ Try It Yourself

1. From the repo root, run `go list ./...` and count how many packages are listed.
2. Run `go list -json .` inside this lesson folder and skim the full JSON structure it returns.
3. Try `go list -m all` and notice this module currently has **no external dependencies** — just
   itself.

## ⚠️ Common Mistakes

- Confusing package listing (`go list ./...`) with module listing (`go list -m all`) — packages
  are the code in this repo; modules include every dependency too.
- Forgetting the `./...` pattern means "this package and every package below it, recursively" —
  a single `.` only lists the current directory's package.
