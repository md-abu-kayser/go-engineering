# 56 — Module Cache Overview

## 🎯 Learning Objectives

- Locate Go's module cache on your machine.
- Understand its on-disk layout: how modules and versions are organized.
- Clear it safely with `go clean -modcache`, and know when that's actually useful.

## 📖 Concept

Every time `go build`/`go mod tidy`/`go get` needs a dependency, Go downloads it **once** and
stores it in a local **module cache** — every subsequent project on your machine that needs the
same module version reuses that same cached copy instead of re-downloading it.

### Finding it

```bash
go env GOMODCACHE
```

Commonly something like `~/go/pkg/mod` on Linux/macOS. Every module your Go installation has
ever downloaded lives somewhere under this one directory.

### The on-disk layout

```
$GOMODCACHE/
├── github.com/
│   └── google/
│       └── uuid@v1.6.0/       <- exact source for this exact version
├── golang.org/
│   └── x/
│       └── tools@v0.21.0/
└── cache/
    └── download/                 <- raw downloaded module zips + metadata, before extraction
```

Notice the pattern: `<host>/<path>@<version>/` — every distinct version of every module gets its
**own** directory, which is precisely what makes it safe for two different projects on your
machine to depend on two different versions of the same module simultaneously, with no conflict.

### Files in the cache are read-only

Go deliberately makes cached module source files **read-only** on disk
([lesson 50](../50-file-permissions) covers the permission mechanics) — this prevents you (or a
build tool) from accidentally editing a "dependency" that's actually meant to be an immutable,
verified copy of published code.

### Clearing it: `go clean -modcache`

```bash
go clean -modcache
```

Deletes the **entire** module cache. Since Go re-downloads (and re-verifies, see
[lesson 57](../57-checksum-database-overview)) anything it needs again automatically, this is
always safe — just potentially slow the next time you build anything, since everything gets
re-fetched. Useful for reclaiming disk space, or ruling out a corrupted cache entry while
debugging a strange build issue.

## 🔍 Code Walkthrough (`main.go`)

```go
out, err := exec.Command("go", "env", "GOMODCACHE").Output()
```

Rather than trying to compute `GOMODCACHE`'s value independently, this shells out to `go env`
itself — the same principle from [lesson 14](../14-go-env-command): always trust the toolchain's
own answer for its effective configuration, rather than re-deriving defaults yourself.

## ▶️ How to Run

```bash
cd level-00-getting-started/56-module-cache-overview
go run main.go
```

## ✅ Expected Output (shape)

```
=== Module Cache Overview ===
----------------------------------
GOMODCACHE: /root/go/pkg/mod

This repository currently depends on zero external modules, so this
lesson's own module cache entry (if any) is just the standard library,
which isn't cached the same way third-party modules are.

See the README for the cache's on-disk layout and how to clear it.
```

(Your `GOMODCACHE` path will differ based on your machine and OS.)

## 🧠 Key Takeaways

- `GOMODCACHE` (found via `go env GOMODCACHE`) is where every downloaded module version lives.
- The layout is `<host>/<path>@<version>/` — every version gets its own directory, enabling
  multiple projects to depend on different versions of the same module with no conflict.
- Cached module files are read-only by design, protecting against accidental modification.
- `go clean -modcache` safely clears everything — Go re-downloads and re-verifies as needed.

## 🛠️ Try It Yourself

1. Run `go env GOMODCACHE` yourself and `ls` the directory it points to (if it exists yet on your
   machine).
2. In a scratch module, add a real external dependency (e.g. `rsc.io/quote`), run
   `go mod tidy`, then find that exact module+version's folder under your `GOMODCACHE`.
3. Try to edit a file inside that cached module folder directly, and observe the permission
   error — direct proof of the read-only protection described above.

## ⚠️ Common Mistakes

- Manually deleting individual files inside the module cache instead of using
  `go clean -modcache` — this can leave the cache in an inconsistent state; let the tooling manage
  it.
- Assuming the module cache is project-specific — it's **global**, shared machine-wide across
  every Go project you build.
