# 20 — Source Layout

## 🎯 Learning Objectives

- Recognize the simplest valid Go project layout.
- Know when a flat layout is the *right* choice, not just the lazy one.
- Understand what typically gets added first as a project grows.

## 📖 Concept

Unlike some ecosystems, Go does not mandate a folder structure. The simplest possible valid Go
project is:

```
my-project/
├── go.mod
└── main.go
```

That's it — one file, one module manifest, and it builds and runs. Every lesson in this
repository so far has essentially been this minimal shape (with a `README.md` added for
teaching purposes, and `go.mod` shared at the repo root instead of per-lesson).

### When flat is right

For small tools, scripts, and single-purpose programs, keep it flat. Splitting into
subpackages purely for the sake of "organization" before you have a real reason adds import
overhead with no benefit. A good rule of thumb: **don't create a new package until you have a
concrete reason to hide or reuse something.**

### What typically gets added first, as a project grows

| Addition | Purpose | Covered in |
|---|---|---|
| A second `.go` file, same package | Split one large file for readability | — |
| A subpackage (e.g. `internal/store`) | Group related code, hide implementation details | [lesson 22](../22-internal-packages) |
| `cmd/` | Support multiple independent binaries from one module | [lesson 23](../23-cmd-directory) |
| `pkg/` | Separate "library code meant for reuse" from application code | [lesson 24](../24-pkg-directory) |

This repository itself is a working example of layout evolving with purpose: it started as flat
lesson folders, and later lessons ([22](../22-internal-packages)–[26](../26-examples-directory))
introduce the structures real, larger Go projects reach for.

## 🔍 Code Walkthrough (`main.go`)

This lesson's code is intentionally unremarkable — the *folder* is the lesson. Notice there's no
subpackage, no `internal/`, nothing beyond what lesson 03 (Hello, World) already had. That's the
point: a flat layout stays valid indefinitely, for as long as it keeps making sense.

## ▶️ How to Run

```bash
cd level-00-getting-started/20-source-layout
go run main.go
```

## ✅ Expected Output

```
=== Source Layout ===
----------------------------------
This lesson folder IS the example: one go.mod (at the repo root),
plain .go files, no unnecessary subfolders. See the README for when
(and how) to grow beyond this.
```

## 🧠 Key Takeaways

- The minimum valid Go project is one file plus `go.mod`.
- A flat layout is a legitimate, permanent choice for small programs — not just a starting point.
- Structure should be added for a concrete reason (hiding details, multiple binaries, reuse), not
  as a default.
- This repository's own later lessons demonstrate the structures you reach for once a flat layout
  stops being enough.

## 🛠️ Try It Yourself

1. Look back at lessons 01–19 and confirm every one of them uses this same flat shape.
2. Sketch (on paper or in a text file) what you'd add first if this lesson's `main.go` grew to
   need a second, reusable helper function shared with another program.
3. Preview [lesson 22](../22-internal-packages) and [lesson 23](../23-cmd-directory) to see the
   two most common next steps.

## ⚠️ Common Mistakes

- Over-structuring a small tool with `internal/`, `pkg/`, and `cmd/` from day one, adding import
  ceremony with no actual benefit yet.
- Under-structuring a genuinely large, multi-binary project by cramming everything into one flat
  package, making it hard to tell what's meant to be reused versus what's private.
