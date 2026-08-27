# 19 — Semantic Import Versioning

## 🎯 Learning Objectives

- Explain Go's "semantic import versioning" rule.
- Recognize why it exists: allowing two incompatible major versions of the same module to
  coexist in one build.
- Correctly import a `v2+` module.

## 📖 Concept

Go modules follow [Semantic Versioning](https://semver.org/) — `MAJOR.MINOR.PATCH` — and layers
one extra rule on top that most package managers don't have: **starting at major version 2, the
major version becomes part of the import path itself.**

| Version range | Import path |
|---|---|
| `v0.x.x` or `v1.x.x` | `github.com/example/project` |
| `v2.x.x` | `github.com/example/project/v2` |
| `v3.x.x` | `github.com/example/project/v3` |

So upgrading from `v1.4.0` to `v2.0.0` isn't just a version bump — it changes **every import
statement** that uses the module, from:

```go
import "github.com/example/project"
```

to:

```go
import "github.com/example/project/v2"
```

### Why this exists

Semantic versioning promises that a major version bump (`v1` → `v2`) can include breaking
changes. Go's module system takes that promise literally: because `v1` and `v2` have **different
import paths**, they are treated as **completely different packages** as far as the build is
concerned. This means:

- Two different dependencies in your project can each depend on a *different* major version of
  the same module, and both work simultaneously, with no conflict.
- There's no ambiguity about which major version a given `import` line refers to — the version is
  right there in the path.

This is a deliberate trade-off other ecosystems don't make, and it's one of the reasons large Go
dependency graphs tend to resolve cleanly rather than deadlocking on version conflicts.

## 🔍 Code Walkthrough (`main.go`)

The example table in `main.go` mirrors the rule directly: notice `v0` and `v1` **share** an
import path (no suffix), while `v2` and `v3` each get their own distinct suffix.

## ▶️ How to Run

```bash
cd level-00-getting-started/19-semantic-import-versioning
go run main.go
```

## ✅ Expected Output

```
=== Semantic Import Versioning ===
----------------------------------
v0.x.x or v1.x.x -> github.com/example/project
v2.x.x           -> github.com/example/project/v2
v3.x.x           -> github.com/example/project/v3

Notice: only v2 and above change the import path itself.
```

## 🧠 Key Takeaways

- `v0`/`v1` share one import path; `v2+` each get a distinct `/vN` suffix.
- This lets incompatible major versions of the same module coexist in one dependency graph.
- Upgrading a dependency across a major version boundary always means updating every import
  statement that uses it, not just a version number.
- This rule only applies to modules meant to be imported by others — it's invisible for
  application code no one else depends on (like this repo).

## 🛠️ Try It Yourself

1. Search GitHub for a popular Go module that has published a `v2` (e.g. search for a repository
   with a `/v2` subfolder or a `v2` branch) and look at how its `go.mod` declares its own module
   path with the `/v2` suffix.
2. Explain, in your own words, why `v2` needing a new import path avoids the "two different
   things with the same name" problem other ecosystems sometimes hit.

## ⚠️ Common Mistakes

- Bumping a `go.mod`'s `module` line to include `/v2` without actually meaning to publish a new
  major version — this is a real, binding declaration, not cosmetic.
- Forgetting to update the module's own internal imports (of itself) when adding the `/v2` suffix
  during a real major version release.
