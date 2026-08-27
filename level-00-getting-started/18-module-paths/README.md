# 18 — Module Paths

## 🎯 Learning Objectives

- Explain how a package's import path is derived from its module path and folder location.
- Read any Go import path and correctly identify its module.
- Understand why module paths commonly look like URLs.

## 📖 Concept

A **module path** is the root identifier declared in `go.mod` (`module go-engineering`, for this
repo). Every package **inside** that module has an import path equal to:

```
<module path> + "/" + <folder path relative to the module root>
```

For example, this very lesson's package has the import path:

```
go-engineering/level-00-getting-started/18-module-paths
```

You never declare this explicitly anywhere — it falls directly out of two things: the `module`
line in `go.mod`, and where you physically put the `.go` files.

### Why module paths often look like URLs

Real, publishable modules use paths like:

```
github.com/gin-gonic/gin
golang.org/x/tools
```

This isn't a coincidence or a style preference — Go's tooling uses the module path as a literal
**fetch location**. `go get github.com/gin-gonic/gin` works because Go can turn that path
directly into a `git clone` (or equivalent) of that exact repository. There's no separate package
registry to publish to, unlike npm or PyPI — the module path *is* the source of truth.

### Reading an import path

Given an import like:

```go
import "github.com/spf13/cobra/doc"
```

You can read it as: *"the module `github.com/spf13/cobra`, specifically its `doc` subpackage."*
Recognizing the module boundary inside a longer import path is a skill that becomes automatic
with practice — it's always everything up to (and typically including) the repository name.

## 🔍 Code Walkthrough (`main.go`)

```go
fullImportPath := modulePath + "/" + packageDir
```

This line makes the concept concrete: string concatenation is literally how Go derives every
package's import path from the module path plus its location on disk.

## ▶️ How to Run

```bash
cd level-00-getting-started/18-module-paths
go run main.go
```

From the repo root, confirm it against reality:

```bash
go list ./level-00-getting-started/18-module-paths
```

## ✅ Expected Output

```
=== Module Paths ===
----------------------------------
Module path            : go-engineering
Package's folder       : level-00-getting-started/18-module-paths
Package's import path  : go-engineering/level-00-getting-started/18-module-paths

Every package's import path = <module path> + <its folder path>.
```

## 🧠 Key Takeaways

- A package's import path = its module's path + its folder path within the module.
- Module paths for publishable code conventionally match a real repository URL.
- There's no separate package registry — `go get` fetches directly from the module path.
- `go list <import-path>` is the ground truth for confirming any package's import path.

## 🛠️ Try It Yourself

1. Run `go list ./...` from the repo root and note how every listed path starts with
   `go-engineering/`.
2. Pick any standard library import you've used (e.g. `"strings"`) and note that it has **no**
   host prefix — the standard library is special-cased and needs no module path.
3. Look at a real open-source Go project's `go.mod` on GitHub and identify its module path.

## ⚠️ Common Mistakes

- Assuming a package's import path is just its folder name — it's always the **full path from
  the module root**, module path included.
- Trying to `import` a sibling package by a short local name instead of its full import path —
  Go always requires the complete path, even for packages in the same repository.
