# 30 — gopls Overview

## 🎯 Learning Objectives

- Explain what `gopls` is and why almost every modern Go editor setup depends on it.
- Know the core categories of features it provides.
- Deliberately exercise several of them on real code, so they stop feeling like "editor magic".

## 📖 Concept

`gopls` (pronounced "go please") is the **official Go language server** — a single background
process, maintained by the Go team, that understands your entire module's code and answers
questions from your editor over a standard protocol (LSP, the Language Server Protocol). Every
editor feature demonstrated in [lesson 29](../29-editor-workflow) — hover docs, go-to-definition,
inline diagnostics, rename — is `gopls` running underneath, not a VS-Code-specific feature. The
exact same `gopls` process works behind Vim, Neovim, Emacs, GoLand, and any other LSP-aware
editor.

### What gopls actually provides

| Category | What it means in practice |
|---|---|
| **Diagnostics** | Live compile errors and `go vet`-style warnings, as you type — no manual build needed. |
| **Completion** | Context-aware autocomplete, including unexported struct fields and method sets. |
| **Navigation** | Go-to-definition, find-all-references, go-to-implementation. |
| **Hover** | Shows a symbol's doc comment and signature without leaving your cursor position. |
| **Refactoring** | Safe rename across the whole module, "extract function"/"extract variable" code actions. |
| **Formatting** | Runs `gofmt`/`goimports` on save, module-aware. |
| **Signature help** | Shows a function's parameter names/types while you're typing a call. |

### Why a separate process, instead of editor plugins doing this directly?

Understanding Go well enough to answer these questions requires essentially the **same analysis
the compiler does** — type-checking, import resolution, and so on. Rather than every editor
reimplementing that independently (and inevitably drifting from the real compiler's behavior),
`gopls` centralizes it once, using the actual `go/*` standard library packages the compiler
itself is built on. Every editor gets the same, compiler-accurate answers "for free".

## 🔍 Code Walkthrough (`main.go`)

This file exists purely as a practice surface — `invoice`, `lineItem`, and `totalCents` are
deliberately simple so you can focus entirely on the editor interactions in the checklist below,
not on understanding the logic.

## ▶️ How to Run

```bash
cd level-00-getting-started/30-gopls-overview
go run main.go
```

Then, in VS Code, work through this checklist on this file:

1. **Hover** over `lineItem` inside `invoice`'s field — see its doc/type info in a tooltip.
2. **Go to definition** (`F12`) on `fmt.Printf` — jump straight into the standard library source.
3. **Find references** (`Shift+F12`) on `cents` — see every place that field is used.
4. **Rename** (`F2`) `totalCents` to `sumCents` — watch both the method definition and its call
   site update together.
5. Introduce a deliberate type mismatch (e.g. `total += item.description`) and watch the
   diagnostic appear inline, instantly — then undo it.

## ✅ Expected Output

```
Total: $6.25

See the README for specific gopls features to try on this file.
```

## 🧠 Key Takeaways

- `gopls` is the single, official language server behind Go tooling in every modern editor.
- It reuses the same `go/*` packages the compiler itself uses — its answers are compiler-accurate.
- Diagnostics, completion, navigation, hover, rename, and formatting are all `gopls` features,
  not editor-specific reimplementations.
- Understanding `gopls` explains *why* the workflow in [lesson 29](../29-editor-workflow) works
  the same way across very different editors.

## 🛠️ Try It Yourself

1. Work through the six-step checklist above on this lesson's `main.go`.
2. Open a terminal and run `gopls version` to see it's a real, standalone binary Go installed on
   your machine.
3. Look up which editors besides VS Code you could use `gopls` with, and note that switching
   editors in the future won't mean re-learning Go tooling from scratch.

## ⚠️ Common Mistakes

- Assuming Go tooling features are tied to VS Code specifically — they're tied to `gopls`, which
  is editor-agnostic.
- Not restarting `gopls` after a large structural change (e.g. adding a new module) — VS Code's
  Go extension has a "Restart Language Server" command in the command palette for exactly this.
