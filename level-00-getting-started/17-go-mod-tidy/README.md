# 17 — go mod tidy

## 🎯 Learning Objectives

- Understand exactly what `go mod tidy` adds, removes, and verifies.
- Understand the relationship between `go.mod` and `go.sum`.
- Know when to run `go mod tidy` in a normal development workflow.

## 📖 Concept

As you write Go code, you `import` packages — some standard library, some third-party. Go does
**not** require you to manually edit `go.mod` every time you add or remove an import. Instead,
you write your code and then run:

```bash
go mod tidy
```

which:

1. **Adds** any dependency your code imports but `go.mod` doesn't yet list, resolving it to a
   compatible version.
2. **Removes** any dependency `go.mod` lists that your code no longer actually imports anywhere.
3. Updates `go.sum` with cryptographic checksums for every dependency, so future downloads of the
   exact same versions can be verified as unmodified.

### `go.mod` vs `go.sum`

| File | Purpose |
|---|---|
| `go.mod` | Declares your module's name, Go version, and **direct + indirect dependency versions**. |
| `go.sum` | Cryptographic checksums for every dependency (and its dependencies), used to verify integrity on download. Never edited by hand. |

Both files are meant to be committed to version control. `go.sum` in particular exists purely for
security and reproducibility — you'll rarely open it, but you should never delete or hand-edit it.

### A worked example (hypothetical)

Imagine you add this import to a file:

```go
import "github.com/google/uuid"
```

Before running `go mod tidy`, `go build` would fail with something like:

```
no required module provides package github.com/google/uuid
```

After running `go mod tidy`, Go automatically fetches the latest compatible version, adds a line
like `require github.com/google/uuid v1.6.0` to `go.mod`, and records its checksum in `go.sum`.
Delete the import later and run `go mod tidy` again, and that `require` line disappears.

## 🔍 Code Walkthrough (`main.go`)

This lesson uses only the standard library on purpose, to keep the whole repository free of
external dependencies — which is exactly why running `go mod tidy` at the repo root right now
would report no changes. The concept is demonstrated in prose above rather than in running code.

## ▶️ How to Run

```bash
cd level-00-getting-started/17-go-mod-tidy
go run main.go
```

From the repository root:

```bash
go mod tidy
git diff go.mod go.sum   # should show no changes — this repo is already tidy
```

## ✅ Expected Output

```
=== go mod tidy ===
----------------------------------
This module currently has zero external dependencies, so `go mod tidy`
has nothing to add or remove here. See the README for what it does on a
module that *does* depend on third-party packages.
```

## 🧠 Key Takeaways

- `go mod tidy` keeps `go.mod`'s dependency list in sync with what your code actually imports.
- It adds missing requirements and removes unused ones — you never hand-edit `require` lines.
- `go.sum` is the checksum ledger for dependencies; commit it, never hand-edit it.
- Run `go mod tidy` after adding or removing any import from an external module.

## 🛠️ Try It Yourself

1. In a scratch module (`go mod init example.com/scratch`), add
   `import "rsc.io/quote"` to a `main.go`, run `go mod tidy`, and inspect the resulting
   `go.mod`/`go.sum`.
2. Remove the import, run `go mod tidy` again, and watch the dependency disappear from `go.mod`.
3. Run `go mod tidy -v` to see verbose output about exactly what changed.

## ⚠️ Common Mistakes

- Hand-editing version numbers in `go.mod` instead of using `go get <module>@<version>` followed
  by `go mod tidy` — manual edits are easy to get subtly wrong.
- Committing code with an import that fails `go mod tidy` cleanly (e.g. an unreachable/private
  module) — always run `go mod tidy` and `go build ./...` before committing.
