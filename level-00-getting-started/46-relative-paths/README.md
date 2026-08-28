# 46 — Relative Paths

## 🎯 Learning Objectives

- Build paths portably with `filepath.Join`, instead of concatenating strings with `"/"`.
- Compute a relative path between two locations with `filepath.Rel`.
- Split a path into its directory, base name, and extension.

## 📖 Concept

[Lesson 45](../45-working-directory) established that relative paths are resolved against the
current working directory. This lesson is about **constructing** and **interpreting** paths
correctly — which matters more than it might seem, because path separators differ across
operating systems (`/` on Linux/macOS, `\` on Windows).

### Building paths: `filepath.Join`

```go
filepath.Join("data", "reports", "2026", "summary.csv")
// -> "data/reports/2026/summary.csv" on Linux/macOS
// -> "data\reports\2026\summary.csv" on Windows
```

`filepath.Join` uses the **correct separator for the current OS** automatically, and also
**cleans** the result — collapsing redundant slashes and resolving `.`/`..` segments:

```go
filepath.Join("data//reports/../reports", "./summary.csv")
// -> "data/reports/summary.csv" — cleaned up automatically
```

Never build a path with `"data" + "/" + "reports"` by hand — it's not portable, and it doesn't
get this cleanup for free.

### Computing a relative path: `filepath.Rel`

```go
rel, err := filepath.Rel("/home/gopher/project", "/home/gopher/project/data/reports")
// rel == "data/reports"
```

Given a base directory and a target path, `filepath.Rel` computes the relative path that gets you
from one to the other — the inverse operation of resolving a relative path against a base.

### Taking a path apart

```go
filepath.Dir("data/reports/summary.csv")   // -> "data/reports"
filepath.Base("data/reports/summary.csv")  // -> "summary.csv"
filepath.Ext("data/reports/summary.csv")   // -> ".csv"
```

These three cover the vast majority of "I have a path, I need one piece of it" needs, without
manual string splitting on separators (which, again, differ by OS).

## 🔍 Code Walkthrough (`main.go`)

```go
messy := filepath.Join("data//reports/../reports", "./summary.csv")
```

This input is deliberately "messy" — a doubled slash, a `..` that cancels out `reports`, and a
redundant `./` — specifically to demonstrate `filepath.Join`'s cleanup behavior isn't just about
joining segments with the right separator; it genuinely normalizes the result.

## ▶️ How to Run

```bash
cd level-00-getting-started/46-relative-paths
go run main.go
```

## ✅ Expected Output

```
=== Relative Paths ===
----------------------------------
filepath.Join(...)        : data/reports/2026/summary.csv
filepath.Join (messy in)  : data/reports/summary.csv
filepath.Rel(...)         : data/reports
filepath.Dir(joined)      : data/reports/2026
filepath.Base(joined)     : summary.csv
filepath.Ext(joined)      : .csv
```

(On Windows, the separators in the first two lines would be `\` instead of `/`.)

## 🧠 Key Takeaways

- `filepath.Join` is the portable, correct way to build a path — never hand-concatenate with `/`.
- `filepath.Join` also cleans the result: redundant slashes and `.`/`..` segments are resolved.
- `filepath.Rel` computes the relative path between two known locations.
- `filepath.Dir`/`Base`/`Ext` split a path into directory, final element, and extension.

## 🛠️ Try It Yourself

1. Try `filepath.Join("data", "..", "logs", "app.log")` and predict the cleaned result before
   running it.
2. Call `filepath.Rel` with two paths that **don't** share a common ancestor easily reachable via
   `..` segments, and read the error it returns.
3. On a Windows machine (or by reading Go's docs if you don't have one handy), compare
   `filepath.Join`'s output separator to what you'd see on Linux/macOS for the exact same inputs.

## ⚠️ Common Mistakes

- Concatenating paths manually (`dir + "/" + file`) instead of `filepath.Join` — this breaks on
  Windows and skips the automatic cleanup.
- Using the unrelated `path` package (not `path/filepath`) for filesystem paths — `path` is
  specifically for **URL-style, always-forward-slash** paths, not OS filesystem paths; mixing them
  up is a common, subtle source of cross-platform bugs.
