# 54 — Exported Identifiers

## 🎯 Learning Objectives

- Define an exported identifier: any name starting with an **uppercase** letter.
- Use an exported function and an exported variable from a genuinely separate package.
- Confirm this is Go's entire visibility mechanism — no `public` keyword, just capitalization.

## 📖 Concept

Go has no `public`/`private` keywords. Instead, visibility across package boundaries is
determined entirely by a name's **first letter**:

```go
func Average(nums []int) float64 { ... } // EXPORTED — starts with uppercase A
var Threshold = 50.0                       // EXPORTED — starts with uppercase T
```

Any identifier — a function, a variable, a type, a struct field, a constant — starting with an
**uppercase** letter is **exported**: reachable from any other package that imports this one.
[Lesson 55](../55-unexported-identifiers) covers the opposite case.

### A genuinely separate package, not just a function in the same file

This lesson includes a real second package, `stats/`, specifically so the example is authentic —
`main.go` **imports** it, exactly as it would import any third-party or standard-library package:

```go
import "go-engineering/level-01-fundamentals/54-exported-identifiers/stats"

avg := stats.Average(nums)     // reachable — Average is exported
threshold := stats.Threshold    // reachable — Threshold is exported
```

Note the access pattern: `stats.Average`, not just `Average` — exported identifiers from another
package are always accessed through the package's name (or its import alias), the same qualified
form you've used with `fmt.Println`, `strconv.Atoi`, and every standard-library call throughout
this repository so far.

## 🔍 Code Walkthrough (`stats/stats.go` and `main.go`)

```go
// package stats
func Average(nums []int) float64 { ... }
var Threshold = 50.0
```

Both `Average` and `Threshold` are capitalized specifically so `main.go`, in a different package,
can use them. If either were lowercase, this exact same `main.go` code would fail to compile —
[lesson 55](../55-unexported-identifiers) demonstrates precisely that failure.

## ▶️ How to Run

```bash
cd level-01-fundamentals/54-exported-identifiers
go run main.go
```

## ✅ Expected Output

```
=== Exported Identifiers ===
----------------------------------
stats.Average([60 75 40 90]) = 66.2
stats.Threshold   = 50.0
Average meets the exported threshold.
```

## 🧠 Key Takeaways

- Go has no `public`/`private` keywords — export status is determined entirely by a name's first
  letter.
- An uppercase first letter makes any identifier (function, variable, type, struct field, constant)
  exported and reachable from other packages.
- Exported identifiers are always accessed as `packageName.Identifier` from outside their package.
- This one simple rule is Go's **entire** cross-package visibility mechanism.

## 🛠️ Try It Yourself

1. Add a second exported function to `stats`, `Max(nums []int) int`, and call it from `main.go`.
2. Add an exported struct type to `stats` (e.g. `type Summary struct { Min, Max int }`) with
   exported fields, and construct one from `main.go`.
3. Run `go doc ./stats` (from this lesson's folder) and confirm it lists exactly the exported
   identifiers you'd expect — matching [level 00's `go doc` lesson]
   (../../level-00-getting-started/13-go-doc-command).

## ⚠️ Common Mistakes

- Forgetting the package qualifier when calling an imported package's exported function
  (`Average(nums)` instead of `stats.Average(nums)`) — Go always requires the qualifier for
  identifiers from another package, even exported ones.
- Assuming exporting a function also automatically documents it — capitalization controls
  *visibility*; documentation is a separate, deliberate act (doc comments,
  [lesson 56](../56-doc-comments)).
