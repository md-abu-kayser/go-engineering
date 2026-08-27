# 11 — go fmt

## 🎯 Learning Objectives

- Understand what `gofmt` (and its wrapper, `go fmt`) actually changes.
- Run `gofmt` from the command line to check and fix formatting.
- Understand why Go deliberately has **no configuration options** for formatting.

## 📖 Concept

Almost every language community eventually argues about tabs vs. spaces, brace placement, and
line length. Go sidesteps the entire debate: `gofmt` is the **one, canonical formatter**, it
ships with every Go installation, and it has no configuration file. Every Go codebase in the
world — the standard library, this repo, any company's internal code — is formatted by the same
tool with the same rules. This means:

- Code review never wastes time on formatting nitpicks.
- Every Go file you'll ever read has the same visual shape.
- Editors can auto-format on save with zero setup decisions.

### `gofmt` vs `go fmt`

```bash
gofmt -l .      # list files that are NOT correctly formatted (empty output = all clean)
gofmt -w .      # rewrite files in place to fix formatting
gofmt -d .      # show a diff of what would change, without writing anything

go fmt ./...    # a thin wrapper around `gofmt -l -w` for the current module
```

`go fmt` is the convenient, module-aware version; `gofmt` is the underlying tool with more flags
(`-l`, `-d`, and others useful in scripts and CI).

### What `gofmt` actually does

- Converts indentation to **tabs**, consistently.
- Aligns struct fields, and consecutive assignments where appropriate.
- Normalizes spacing around operators.
- Sorts and groups imports.
- Removes unnecessary parentheses.

It does **not** rename variables, reorder functions, or make any decision that changes what your
code means — only how it's laid out.

## 🔍 Code Walkthrough (`main.go`)

The `Point` struct in this file is written exactly as `gofmt` would produce it. If you added
extra spaces before `int` on either field, `gofmt -l .` would immediately flag this file as
needing formatting — try it (see below).

## ▶️ How to Run

```bash
cd level-00-getting-started/11-go-fmt-command
go run main.go
gofmt -l .
```

## ✅ Expected Output

```
Point{X: 3, Y: 4}

This file is already gofmt-clean. See the README for how to prove it.
```

`gofmt -l .` should print **nothing** — no output means every file in this folder is already
correctly formatted.

## 🧠 Key Takeaways

- `gofmt` is Go's single, non-configurable code formatter.
- `gofmt -l` lists misformatted files; `gofmt -w` fixes them in place.
- `go fmt ./...` is the convenient module-wide shortcut.
- Consistent formatting across the entire Go ecosystem removes an entire category of debate.

## 🛠️ Try It Yourself

1. Open `main.go` and deliberately misalign the struct fields or add extra spaces around `:=`.
2. Run `gofmt -l .` and see this file listed as needing formatting.
3. Run `gofmt -w .` to auto-fix it, then `gofmt -l .` again to confirm it's clean.
4. In VS Code, enable "format on save" for Go (the Go extension does this by default) and watch
   it happen automatically every time you save.

## ⚠️ Common Mistakes

- Manually fighting with indentation instead of just running `gofmt -w .` — let the tool do it.
- Assuming `gofmt` also manages imports (adding/removing them) — that's `goimports`, a separate,
  related tool covered conceptually in [lesson 05](../05-imports) and referenced again in
  [lesson 29](../29-editor-workflow).
