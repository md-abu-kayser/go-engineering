# 53 — Reproducible Build Basics

## 🎯 Learning Objectives

- Explain what a "reproducible build" means, and why it's valuable.
- Find your own local file paths accidentally embedded in a compiled binary.
- Use `-trimpath` to remove that source of non-determinism.

## 📖 Concept

A build is **reproducible** if compiling the exact same source code, with the exact same
toolchain version, always produces a **byte-for-byte identical** binary — regardless of who runs
the build, or from what directory, or when. This matters for real reasons beyond pedantry:

- **Supply-chain security**: if a binary is reproducible, anyone can independently verify that a
  published binary actually corresponds to the published source — no hidden extra code slipped in
  during the build.
- **Build caching**: reproducible outputs mean a CI system (or `go build`'s own cache) can safely
  skip recompiling something it's already built identically before.

### The most common accidental source of non-determinism: file paths

By default, `go build` embeds the **absolute path to your source directory** into the compiled
binary, for use in stack traces and debugging info:

```bash
go build -o app .
strings app | grep "$(pwd)"   # your local path is very likely in there
```

Two developers building the exact same code from `/home/alice/project` and
`/Users/bob/dev/project` will get **different** binaries — not because the logic differs, but
purely because of where each of them happened to check out the repository.

### Fixing it: `-trimpath`

```bash
go build -trimpath -o app .
```

`-trimpath` strips local filesystem paths from the compiled binary, replacing them with paths
relative to the module root instead. Combined with a pinned Go version and `CGO_ENABLED=0`
([lesson 48](../48-cross-compilation-overview)), this gets you most of the way to a genuinely
reproducible build.

### What still needs care

Reproducibility isn't automatic even with `-trimpath` — a couple of other common culprits:

- **Embedding a timestamp** (e.g. `time.Now()` baked into a version string at build time) — by
  definition, this changes every single build. If you need a build timestamp, that's fine, just
  understand it means the binary is deliberately *not* reproducible in that one specific respect.
- **Map iteration order** — Go deliberately randomizes map iteration order at runtime (not at
  compile time, so this doesn't affect binary reproducibility itself, but it's a related
  "surprising non-determinism" trap worth knowing about).

## 🔍 Code Walkthrough (`main.go`)

This lesson's code is intentionally simple — the actual lesson content is in the **build
commands** you run against it, not the program's logic, which is identical either way.

## ▶️ How to Run

```bash
cd level-00-getting-started/53-reproducible-build-basics
go build -o app-with-path .
go build -trimpath -o app-trimmed .
strings app-with-path | grep 53-reproducible-build-basics
strings app-trimmed | grep 53-reproducible-build-basics
```

## ✅ Expected Output (of `go run main.go`)

```
=== Reproducible Build Basics ===
----------------------------------
This program's own logic doesn't change between builds — but by
default, Go embeds your LOCAL FILE PATH into the compiled binary.

See the README for how -trimpath removes that, and why it matters
for producing byte-identical binaries from identical source.
```

## 🧠 Key Takeaways

- A reproducible build produces byte-identical output from identical source, every time.
- By default, Go embeds your absolute local source path into the binary.
- `-trimpath` removes that, replacing it with a module-relative path instead.
- Timestamps and other build-time-computed values are a *separate*, deliberate source of
  non-reproducibility — not something `-trimpath` addresses.

## 🛠️ Try It Yourself

1. Run the `strings ... | grep` commands above and confirm your local path appears in the
   untrimmed binary but not in the trimmed one.
2. Build this lesson twice in a row with `-trimpath` and compare the two binaries with
   `diff <(strings app1) <(strings app2)` — they should be identical (or very close, depending on
   your Go version's exact reproducibility guarantees).
3. Look up `go help buildmode` and `go env GOFLAGS` — note you can set `-trimpath` as a persistent
   default with `go env -w GOFLAGS=-trimpath` instead of typing it on every build.

## ⚠️ Common Mistakes

- Assuming `go build` output is automatically reproducible — it isn't, by default, specifically
  because of the embedded absolute path.
- Adding a build timestamp for "traceability" and then being confused why builds aren't
  reproducible — that's an intentional trade-off, not a bug, and should be a deliberate choice.
