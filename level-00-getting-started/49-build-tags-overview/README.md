# 49 — Build Tags Overview

## 🎯 Learning Objectives

- Write a `//go:build` constraint and see it actually control which file compiles.
- Understand automatic OS/architecture constraints versus custom, hand-defined ones.
- Build with a custom tag using `go build -tags`.

## 📖 Concept

A **build constraint** (informally, a "build tag") tells the Go compiler *"only include this
file when the following condition is true."* Go evaluates these **before** compiling anything —
excluded files aren't just skipped at runtime, they're never even parsed for that build.

### The syntax: `//go:build`

```go
//go:build windows

package main
```

This comment **must** be the very first thing in the file (only preceded by blank lines/other
comments) and **must** be followed by a blank line before `package`. It supports `&&`, `||`, and
`!`:

```go
//go:build linux && amd64        // both conditions
//go:build windows || darwin      // either condition
//go:build !windows                // negation
```

### Automatic constraints: OS and architecture in the filename

Go recognizes certain filename suffixes automatically, with **no `//go:build` comment needed**:
a file named `foo_windows.go` is automatically Windows-only; `foo_linux_arm64.go` is automatically
Linux-on-ARM64-only. This lesson uses the explicit `//go:build` comment form instead (in
`platform_windows.go` / `platform_other.go`) specifically because it's more visible for learning
— but recognize both forms when you see them in real projects.

### Custom build tags

Beyond OS/architecture, you can define your **own** tags for anything — feature flags, build
variants, expensive integration tests you don't want running by default:

```go
//go:build debug

const debugEnabled = true
```

```bash
go run .                   # debug tag NOT set -> debug_off.go is used -> debugEnabled = false
go run -tags debug .        # debug tag IS set  -> debug_on.go is used  -> debugEnabled = true
```

## 🔍 Code Walkthrough

```
49-build-tags-overview/
├── main.go
├── platform_windows.go   // //go:build windows
├── platform_other.go     // //go:build !windows
├── debug_on.go            // //go:build debug
└── debug_off.go            // //go:build !debug
```

Exactly **one** of `platform_windows.go` / `platform_other.go` is compiled into any given build —
never both (their constraints are exact opposites), and never neither (together they cover every
possibility). The same is true for `debug_on.go` / `debug_off.go`. This is the standard pattern
for build-tag-based conditional compilation: a pair (or small set) of files whose constraints are
mutually exclusive and collectively exhaustive.

## ▶️ How to Run

```bash
cd level-00-getting-started/49-build-tags-overview
go run .
go run -tags debug .
```

## ✅ Expected Output

Without the tag:

```
=== Build Tags Overview ===
----------------------------------
platformName() : Not Windows (compiled in via //go:build !windows)
debugEnabled   : false

Try: go run -tags debug . -- and compare debugEnabled above.
```

With `-tags debug`:

```
=== Build Tags Overview ===
----------------------------------
platformName() : Not Windows (compiled in via //go:build !windows)
debugEnabled   : true

Try: go run -tags debug . -- and compare debugEnabled above.
```

(`platformName()` would read "Windows (compiled in via //go:build windows)" only if this were
actually built with `GOOS=windows`.)

## 🧠 Key Takeaways

- `//go:build <expr>` excludes a file from compilation entirely when the expression is false —
  not just at runtime, but before compilation even considers the file.
- Filename suffixes like `_windows.go` or `_arm64.go` apply OS/architecture constraints
  automatically, with no comment needed.
- Custom tags (`-tags <name>`) let you build conditional variants for anything — debug builds,
  optional features, excluded slow tests.
- A constraint pair should be mutually exclusive and collectively exhaustive, so exactly one
  version of a symbol is ever compiled in.

## 🛠️ Try It Yourself

1. Run both commands in "How to Run" and confirm `debugEnabled` really does flip between them.
2. Cross-compile with `GOOS=windows go build .` (no need to run it — it won't execute on Linux)
   and confirm it compiles successfully, pulling in `platform_windows.go` instead.
3. Add a third custom tag, `verbose`, following the same on/off file pattern as `debug`, and wire
   it into `main.go`.

## ⚠️ Common Mistakes

- Forgetting the **required blank line** between the `//go:build` comment and `package main` —
  without it, Go treats the line as an ordinary comment, not a build constraint, and silently
  compiles the file unconditionally.
- Writing constraints that overlap or leave gaps (e.g. both files say `//go:build linux`, or
  neither covers `darwin`) — leading to either a duplicate-symbol build error or an undefined
  symbol on the uncovered platform.
