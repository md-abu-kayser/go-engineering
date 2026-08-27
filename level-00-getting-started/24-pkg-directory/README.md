# 24 — pkg/ Directory

## 🎯 Learning Objectives

- Use the `pkg/` directory convention for exported, reusable code.
- Understand the (real, ongoing) community debate about whether `pkg/` is a good idea.
- Contrast `pkg/` with `internal/` ([lesson 22](../22-internal-packages)).

## 📖 Concept

Some large Go projects use a top-level `pkg/` directory to hold packages that are explicitly
**meant to be imported by other projects** — as opposed to `internal/`, which exists specifically
to **prevent** that.

```
my-project/
├── go.mod
├── cmd/
│   └── myapp/main.go        # the binary — imports pkg/ and internal/
├── pkg/
│   └── stringutil/            # reusable, exported, meant for outside use
└── internal/
    └── store/                  # private implementation detail
```

### The `pkg/` vs. flat-at-module-root debate

This convention is **genuinely contested** in the Go community, and it's worth knowing both
sides rather than treating it as settled:

- **In favor:** for a large project with many packages, `pkg/` gives a clear, greppable signal —
  "everything under here is public API," separate from `cmd/` (binaries) and `internal/`
  (private code). Several very large, well-known Go projects use it successfully.
- **Against:** critics (including some well-known voices in the Go community) argue `pkg/` is
  redundant — in Go, a package's exported identifiers (capitalized names) already signal what's
  public, and every package not under `internal/` is *already* importable by anyone. Adding a
  `pkg/` folder on top doesn't add any enforcement, unlike `internal/`, which the compiler
  actually checks.

**Practical takeaway:** for small-to-medium projects, most idiomatic Go code just puts reusable
packages directly at (or near) the module root, with no `pkg/` wrapper — reserve `pkg/` for
larger projects where the extra signal genuinely earns its keep, and treat it as a style choice,
not a rule.

## 🔍 Code Walkthrough

```
24-pkg-directory/
├── main.go
└── pkg/
    └── stringutil/
        └── stringutil.go
```

```go
import "go-engineering/level-00-getting-started/24-pkg-directory/pkg/stringutil"
```

Unlike `internal/` ([lesson 22](../22-internal-packages)), there is **no compiler-enforced
restriction** here — `pkg/stringutil` could be imported from anywhere, by anyone, in any module.
The `pkg/` folder name is purely a human-readable signal, not an enforced boundary.

## ▶️ How to Run

```bash
cd level-00-getting-started/24-pkg-directory
go run main.go
```

## ✅ Expected Output

```
=== pkg/ directory ===
----------------------------------
rehpoG
Hello From The Pkg Directory
```

## 🧠 Key Takeaways

- `pkg/` is a convention signaling "this is meant to be imported by others" — unlike `internal/`,
  it is **not enforced** by the compiler.
- The Go community is genuinely split on whether `pkg/` earns its complexity for smaller projects.
- Exported (capitalized) identifiers already communicate "this is public" at the package level,
  with or without a `pkg/` wrapper.
- Choose `pkg/` deliberately for larger projects where the extra signal helps; skip it otherwise.

## 🛠️ Try It Yourself

1. Add a third function to `stringutil` (e.g. `IsPalindrome`, reusing the logic from
   [lesson 10](../10-go-test-command)) and call it from `main.go`.
2. Write a small table-driven test for `stringutil.Reverse` (see [lesson 10](../10-go-test-command)
   for the pattern) and save it as `pkg/stringutil/stringutil_test.go`.
3. Read a couple of real-world opinions on `pkg/` online and decide which side you find more
   convincing for your own future projects.

## ⚠️ Common Mistakes

- Believing `pkg/` provides the same enforcement as `internal/` — it provides **none**; it's
  purely organizational.
- Adopting `pkg/` reflexively for a small, single-purpose project, adding directory depth with no
  real benefit.
