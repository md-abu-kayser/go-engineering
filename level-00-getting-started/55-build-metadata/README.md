# 55 — Build Metadata

## 🎯 Learning Objectives

- Embed version/commit/build-date metadata into a binary at build time, not runtime.
- Use `-ldflags "-X pkg.Var=value"` to override package-level string variables during `go build`.
- Write a test that verifies the sensible defaults used when no metadata is injected.

## 📖 Concept

A common, genuinely useful pattern for CLI tools: support `myapp --version` printing something
like `v1.4.2 (commit a1b2c3d, built 2026-08-20)` — without hardcoding that string, and without
reading it from a file at runtime (which requires shipping an extra file alongside the binary).
Go's linker can inject this **directly into the compiled binary**, via `-ldflags -X`.

### The pattern: package-level vars with safe defaults

```go
var (
    version   = "dev"
    commit    = "none"
    buildDate = "unknown"
)
```

Plain `go run`/`go build`/`go test`, with no special flags, sees exactly these defaults — which
is deliberately what makes `go test` reliable in [the test file](#) below: it doesn't need any
special build invocation to get predictable values to assert against.

### Overriding at build time: `-ldflags "-X ..."`

```bash
go build -ldflags "-X main.version=v1.4.2 -X main.commit=a1b2c3d -X main.buildDate=2026-08-20" .
```

`-X importpath.name=value` tells the **linker** (not the compiler) to overwrite that specific
string variable's value as the final binary is assembled — no source code changes needed, and no
runtime cost, since this happens once, at build time.

### Wiring in real values automatically

In practice, these values usually come from your version control system and the build
environment itself, computed in a script or CI step:

```bash
VERSION=$(git describe --tags --always)
COMMIT=$(git rev-parse --short HEAD)
DATE=$(date -u +%Y-%m-%d)

go build -ldflags "-X main.version=$VERSION -X main.commit=$COMMIT -X main.buildDate=$DATE" .
```

This is the exact mechanism behind `--version` output in a huge number of real-world Go CLI
tools — `kubectl version`, `docker version`, and many others all use some variant of this pattern.

## 🔍 Code Walkthrough (`main.go` and the test file)

```go
var (
    version   = "dev"
    ...
)
```

Note these are **package-level `var` declarations, not `const`** — `-X` specifically requires a
`var` (a mutable, linker-patchable location), and it must hold a plain string; `-X` cannot target
a `const`, a struct field, or anything computed.

```go
func TestDefaultMetadata(t *testing.T) {
    ...
    {"version", version, "dev"},
```

This test deliberately asserts the **default** values — proving that ordinary `go test` (which
never passes `-ldflags`) reliably sees `"dev"`/`"none"`/`"unknown"`, which is exactly why those
particular defaults were chosen: safe, obviously-a-placeholder values, not empty strings that
might be mistaken for "successfully built but genuinely versionless."

## ▶️ How to Run

```bash
cd level-00-getting-started/55-build-metadata
go run main.go
go test -v ./...
go build -ldflags "-X main.version=v1.4.2 -X main.commit=a1b2c3d -X main.buildDate=2026-08-20" -o app .
./app
```

## ✅ Expected Output

Plain run (no injected metadata):

```
=== Build Metadata ===
----------------------------------
version   : dev
commit    : none
buildDate : unknown

See the README for the -ldflags -X command that overrides these at
build time — a real release build would wire this into `go build`.
```

With the `-ldflags -X` build above:

```
=== Build Metadata ===
----------------------------------
version   : v1.4.2
commit    : a1b2c3d
buildDate : 2026-08-20

See the README for the -ldflags -X command that overrides these at
build time — a real release build would wire this into `go build`.
```

## 🧠 Key Takeaways

- `-ldflags "-X pkg.Var=value"` overwrites a package-level **string `var`** at link time.
- This only works on `var`, never `const` — `-X` needs a patchable memory location.
- Safe, obviously-placeholder defaults (`"dev"`, `"none"`, `"unknown"`) make plain, un-flagged
  builds and tests behave predictably.
- Real projects typically compute these values from `git describe`/`git rev-parse`/`date` in a
  build script or CI step, not by hand.

## 🛠️ Try It Yourself

1. Run `go test -v ./...` with no special flags and confirm every subtest passes against the
   defaults.
2. Build with the full `-ldflags -X` command shown above and confirm the injected values appear
   exactly as given.
3. Try targeting a `const` instead of a `var` for one of these fields, and read the linker error
   `-X` gives you — confirming the "must be a var" rule for yourself.

## ⚠️ Common Mistakes

- Declaring these as `const` instead of `var` — `-X` will fail with a clear error, but it's an
  easy mistake to make since these values never change *within* a single run.
- Forgetting the full import path in `-X` for anything beyond `package main` — `-X` needs
  `full/import/path.VarName`, not just `VarName`, once the variable lives outside `main`.
