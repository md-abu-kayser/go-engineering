# 22 — internal/ Packages

## 🎯 Learning Objectives

- Use the `internal/` directory convention to restrict what can import a package.
- Understand that this restriction is **enforced by the compiler**, not just documentation.
- Know when to reach for `internal/` in a real project.

## 📖 Concept

Go recognizes one special directory name: `internal`. Any package inside a directory named
`internal` can **only** be imported by code that lives inside the tree rooted at the **parent**
of that `internal` directory.

For this lesson, the package lives at:

```
level-00-getting-started/22-internal-packages/internal/greeting
```

The parent of `internal` here is `22-internal-packages/`. That means:

- ✅ `level-00-getting-started/22-internal-packages/main.go` **can** import
  `.../22-internal-packages/internal/greeting`.
- ❌ Any other lesson folder — or anything outside `22-internal-packages/` entirely — **cannot**,
  even though this is all one module. The compiler rejects it with an explicit error, not a
  warning.

### Why this matters

`internal/` lets you structure a codebase with real, enforced boundaries: "this package is an
implementation detail of this part of the project — it is *not* part of anything meant to be
reused or depended on elsewhere." This is especially valuable for:

- Open-source libraries hiding implementation details from consumers, while still organizing
  internal code into multiple packages for the maintainers' own clarity.
- Large monorepos preventing one team's private helper package from silently becoming a
  cross-team dependency nobody intended to support long-term.

## 🔍 Code Walkthrough

```go
import "go-engineering/level-00-getting-started/22-internal-packages/internal/greeting"
```

This import works specifically **because** `main.go` sits inside `22-internal-packages/`, the
parent of the `internal/` directory. Move `main.go` up one level (to `level-00-getting-started/`
directly) and this same import would fail to compile.

## ▶️ How to Run

```bash
cd level-00-getting-started/22-internal-packages
go run main.go
```

## ✅ Expected Output

```
Hello from an internal package, Gopher!

Try importing 'internal/greeting' from a DIFFERENT lesson folder —
the compiler refuses. See the README for why.
```

## 🧠 Key Takeaways

- A directory named `internal` restricts imports to the tree rooted at its parent — enforced by
  the Go compiler.
- This lets you have multiple, well-organized packages that are still effectively "private" to a
  module or subtree.
- `internal/` is about **import boundaries**, unrelated to unexported (lowercase) identifiers,
  which restrict access at the *package* level instead.

## 🛠️ Try It Yourself

1. Try importing `.../22-internal-packages/internal/greeting` from
   `level-00-getting-started/23-cmd-directory/main.go` (temporarily) and read the exact compiler
   error — then remove the import again.
2. Move the `internal/greeting` folder up one level, to directly under
   `level-00-getting-started/`, and notice every lesson folder can now import it.
3. Move it back down afterward to restore the original restriction.

## ⚠️ Common Mistakes

- Assuming `internal/` prevents access **within** its own allowed tree — it doesn't; it only
  blocks imports from **outside** that tree.
- Confusing `internal/` (import boundary) with lowercase identifiers (export boundary) — a
  package can be fully exported (all uppercase names) and still be unreachable from outside
  simply because it's under `internal/`.
