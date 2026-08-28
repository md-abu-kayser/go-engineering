# 47 — Absolute Paths

## 🎯 Learning Objectives

- Check whether a path is absolute with `filepath.IsAbs`.
- Convert any path to an absolute one with `filepath.Abs`.
- Know when a program should prefer absolute paths over relative ones.

## 📖 Concept

An **absolute path** fully specifies a location from the filesystem root (`/` on Linux/macOS,
`C:\` etc. on Windows) — it means the same thing no matter what the current working directory
happens to be. A **relative path** ([lesson 46](../46-relative-paths)) only means something once
you know the working directory it's resolved against ([lesson 45](../45-working-directory)).

### Checking: `filepath.IsAbs`

```go
filepath.IsAbs("README.md")   // false
filepath.IsAbs("/etc/hosts")  // true
```

A quick, portable check — no need to manually inspect the first character yourself (which would
also need to account for Windows drive letters like `C:\`).

### Converting: `filepath.Abs`

```go
abs, err := filepath.Abs("README.md")
```

Resolves `"README.md"` against the current working directory and returns a clean, absolute
result — effectively `filepath.Join(currentWorkingDirectory, "README.md")`, cleaned up the same
way [lesson 46](../46-relative-paths) described. If the input is already absolute,
`filepath.Abs` returns it unchanged (after cleaning any redundant segments).

### When to prefer absolute paths

Prefer resolving to an absolute path **early**, once, near where a path first enters your
program (a config value, a command-line flag, a discovered file) — rather than passing a
relative path deep into your program and hoping every consumer of it runs from the same working
directory. This matters especially for:

- **Long-running services** (a daemon, a server) — their working directory might be set once at
  startup by whatever launched them (systemd, Docker, a supervisor process) and never match what
  a developer expects while testing locally.
- **Programs that call `os.Chdir`** partway through — any relative path computed *before* the
  `Chdir` and used *after* it silently means something different than intended.

## 🔍 Code Walkthrough (`main.go`)

```go
abs, err := filepath.Abs("README.md")
```

This resolves against whatever `os.Getwd()` currently reports — run this program from different
directories (as [lesson 45](../45-working-directory) demonstrated) and the printed absolute path
changes accordingly, even though the source code is identical.

## ▶️ How to Run

```bash
cd level-00-getting-started/47-absolute-paths
go run main.go
```

## ✅ Expected Output (shape)

```
=== Absolute Paths ===
----------------------------------
filepath.IsAbs("README.md")     : false
filepath.IsAbs("/etc/hosts")     : true
filepath.Abs("README.md")       : /home/claude/GO-ENGINEERING/level-00-getting-started/47-absolute-paths/README.md
filepath.Abs("/etc/hosts")       : /etc/hosts
```

(The exact absolute path in the third line depends on where this repository lives on your
machine.)

## 🧠 Key Takeaways

- An absolute path means the same thing regardless of working directory; a relative one doesn't.
- `filepath.IsAbs` checks which kind you have; `filepath.Abs` converts to absolute.
- `filepath.Abs` on an already-absolute path just cleans it, leaving it unchanged in meaning.
- Resolve to absolute early for anything that might outlive or ignore the working directory it
  started with — long-running services especially.

## 🛠️ Try It Yourself

1. Run this program from a couple of different working directories (as in
   [lesson 45](../45-working-directory)'s exercises) and confirm the absolute path changes to
   match.
2. Combine this with [lesson 43](../43-environment-basics): read a directory from an environment
   variable (with a sensible default), and resolve it to absolute with `filepath.Abs` before using
   it.
3. Look up `filepath.EvalSymlinks` and consider when it might matter beyond what `filepath.Abs`
   alone gives you (hint: symbolic links).

## ⚠️ Common Mistakes

- Assuming a relative path a developer tested locally will resolve the same way once deployed —
  `filepath.Abs` it as early as practical, especially for anything read from configuration.
- Forgetting `filepath.Abs` can still fail (rare, but possible if `os.Getwd()` itself fails) —
  always check its error, the same as any other path- or filesystem-related call.
