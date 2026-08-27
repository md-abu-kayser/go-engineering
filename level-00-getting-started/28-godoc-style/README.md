# 28 — godoc Style

## 🎯 Learning Objectives

- Structure a package's documentation consistently across every exported symbol.
- Know when a dedicated `doc.go` file is the right place for package-level documentation.
- Read an entire package's documentation the way a real consumer would, using `go doc`.

## 📖 Concept

[Lesson 13](../13-go-doc-command) and [lesson 27](../27-documentation-comments) covered the
mechanics — a doc comment above a declaration, and the formatting syntax available inside one.
This lesson is about applying that **consistently across a whole package**, which is what
actually makes a package pleasant to consume.

### `doc.go` — a dedicated home for the package comment

A package-level doc comment can technically go above the `package` clause in **any** file in the
package. But when there's no single file that's the "obvious" place for it (a package with
several equally-important files, or a package split across many small files), the convention is
a dedicated `doc.go` containing **only**:

```go
// Package kvstore provides a minimal, in-memory key-value store.
// ...
package kvstore
```

This lesson's `kvstore/doc.go` follows exactly that pattern — its only job is documentation.

### Consistency across every exported symbol

Compare `kvstore/kvstore.go`'s comments:

```go
// Store is a minimal in-memory key-value store.
//
// The zero value is NOT ready to use — construct one with [New].
type Store struct { ... }

// New returns a ready-to-use, empty Store.
func New() *Store { ... }

// Set stores value under key, overwriting any existing value.
func (s *Store) Set(key, value string) { ... }
```

Notice the pattern repeated for **every** exported symbol: start with the name, one clear
sentence describing behavior, and — critically — the `Store` doc comment explicitly calls out
that its zero value is *not* usable, pointing readers at `New` instead. Documenting zero-value
behavior (usable or not) is one of the highest-value habits in idiomatic Go docs, precisely
because Go doesn't have constructors to force initialization.

## 🔍 Code Walkthrough (`main.go`)

`main.go` here is a straightforward **consumer** of `kvstore` — exactly the kind of code the
documentation in `kvstore/` was written to support. If the docs are good, using the package from
here should require no guessing.

## ▶️ How to Run

```bash
cd level-00-getting-started/28-godoc-style
go run main.go
go doc ./kvstore
go doc -all ./kvstore
```

## ✅ Expected Output

```
=== godoc Style ===
----------------------------------
name = Gopher
stored keys: 2
stored keys after delete: 1

Run `go doc ./kvstore` and `go doc -all ./kvstore` to read this
package exactly the way a new consumer of it would.
```

## 🧠 Key Takeaways

- A dedicated `doc.go` is the idiomatic home for package-level docs when no single file is the
  obvious place for them.
- Document zero-value behavior explicitly — "ready to use" or "construct with `New`" — since Go
  has no constructors to enforce initialization.
- Consistency (same shape of comment for every exported symbol) matters as much as any single
  comment's content.
- `go doc -all ./kvstore` is exactly how a new team member would first explore this package.

## 🛠️ Try It Yourself

1. Run `go doc -all ./kvstore` and read the whole package's rendered documentation top to bottom,
   as if you'd never seen the source.
2. Add a `Keys() []string` method to `Store`, with a doc comment matching the existing style.
3. Move the package doc comment out of `doc.go` and onto `kvstore.go` instead — confirm
   `go doc ./kvstore` renders identically either way (the file location doesn't matter, only
   that exactly one file has it).

## ⚠️ Common Mistakes

- Documenting some exported symbols thoroughly and leaving others with no comment at all —
  `go vet`'s stricter cousins (like `golint`/`staticcheck`) flag this, and it makes a package
  feel unfinished.
- Putting the package doc comment on **more than one** file — only one should have it; duplicates
  are confusing and tooling behavior around them is inconsistent.
