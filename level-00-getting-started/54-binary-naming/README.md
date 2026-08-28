# 54 — Binary Naming

## 🎯 Learning Objectives

- Know exactly where `go build`'s default output name comes from.
- Override it explicitly with `-o`.
- Follow the common `<name>-<os>-<arch>` convention for distributing cross-compiled binaries.

## 📖 Concept

### The default name

```bash
cd level-00-getting-started/54-binary-naming
go build .
ls
```

With no `-o` flag, `go build` names the output binary after the **directory** the `main` package
lives in — here, that produces a binary literally named `54-binary-naming` (or
`54-binary-naming.exe` on Windows — see below). This is a convenience default, not something you
declare anywhere in code.

### Overriding it: `-o`

```bash
go build -o bin/myapp .
```

`-o` lets you choose any name and location instead — extremely common in real projects, since
"whatever the folder happens to be called" is rarely the name you actually want to ship.

### The `.exe` suffix, automatically, on Windows — but only for the DEFAULT name

```bash
GOOS=windows GOARCH=amd64 go build .
ls
# -> app.exe   (the folder-derived default name, with .exe appended automatically)
```

When cross-compiling for `GOOS=windows` (see [lesson 48](../48-cross-compilation-overview)) and
letting `go build` choose the output name itself (no `-o`), Go **automatically** appends `.exe` —
Windows requires it to treat a file as executable, so Go handles this for you rather than
requiring you to remember it.

This auto-appending **only** applies to that default, directory-derived name. If you pass an
explicit `-o` yourself, Go uses **exactly** what you typed, `.exe` or not:

```bash
GOOS=windows GOARCH=amd64 go build -o demo .
ls
# -> demo   (NOT demo.exe — Go respects an explicit -o exactly as given)
```

The resulting `demo` file is still a completely valid Windows executable — it's only the
*filename* that lacks `.exe`, which matters in practice because Windows itself relies on that
extension to recognize a file as runnable. If you're cross-compiling for Windows and specifying
`-o` yourself, add `.exe` deliberately.

### The release-binary naming convention

When shipping the same tool built for multiple platforms, the near-universal community
convention is:

```
<name>-<os>-<arch>[.exe]
```

For example:

```bash
GOOS=linux   GOARCH=amd64 go build -o myapp-linux-amd64     .
GOOS=linux   GOARCH=arm64 go build -o myapp-linux-arm64      .
GOOS=darwin  GOARCH=arm64 go build -o myapp-darwin-arm64      .
GOOS=windows GOARCH=amd64 go build -o myapp-windows-amd64.exe .
```

This is exactly the pattern you'll see in the "Releases" section of virtually any Go CLI tool on
GitHub — one clearly-labeled binary per platform, all built from the same source in one script or
CI job.

## 🔍 Code Walkthrough (`main.go`)

This lesson's code is intentionally simple — the point is entirely in the folder name and the
`go build` commands run against it, demonstrated below.

## ▶️ How to Run

```bash
cd level-00-getting-started/54-binary-naming
go build .
ls 54-binary-naming*        # the default-named binary
go build -o bin/myapp .
./bin/myapp                  # the explicitly-named binary
```

## ✅ Expected Output (of `go run main.go`)

```
=== Binary Naming ===
----------------------------------
This folder is named '54-binary-naming' — that's exactly what
`go build` (with no -o flag) will call the resulting binary.

See the README for overriding it with -o, and for the
<name>-<os>-<arch> convention used for released binaries.
```

## 🧠 Key Takeaways

- With no `-o`, `go build` names the binary after the containing directory.
- `-o <path>` overrides the name and/or location explicitly — used exactly as given, with no
  automatic changes.
- Cross-compiling for `GOOS=windows` with the **default** name automatically appends `.exe`; an
  explicit `-o` does not get this treatment, so add `.exe` yourself when you specify one.
- `<name>-<os>-<arch>[.exe]` is the standard convention for distributing multi-platform binaries.

## 🛠️ Try It Yourself

1. Run `go build .` in this folder (no `-o`) and confirm the resulting binary's name matches the
   folder name exactly.
2. Cross-compile for Windows with **no** `-o` flag and confirm Go automatically appends `.exe` to
   the default, folder-derived name — then try it again **with** `-o demo` (no `.exe` in what you
   type) and confirm this time it does *not* get added automatically.
3. Build this lesson for three different platforms using the naming convention above, and list
   all three resulting files together to see what a real multi-platform release looks like.

## ⚠️ Common Mistakes

- Forgetting the automatic `.exe` suffix applies **only** to the default (no `-o`) name, and then
  scripting a release build with an explicit `-o app` for Windows that ends up missing `.exe`
  entirely — a later step expecting `app.exe` won't find it.
- Shipping binaries all named identically (`app`, `app`, `app`) across platforms in one release —
  without the `-os-arch` suffix, users can't tell which download is meant for their machine.
