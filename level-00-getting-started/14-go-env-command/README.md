# 14 — go env

## 🎯 Learning Objectives

- Understand what `go env` shows and why it matters.
- Distinguish a **raw shell environment variable** from Go's **effective configuration value**.
- Know the most important Go environment variables: `GOROOT`, `GOPATH`, `GOOS`, `GOARCH`,
  `GOMODCACHE`, `GOBIN`.

## 📖 Concept

`go env` prints Go's **effective configuration** — the values Go tooling actually uses, after
applying defaults. This is different from just checking your shell's environment variables,
because many of these values have sensible defaults that apply even when nothing is explicitly
set.

```bash
go env                # print everything
go env GOPATH          # print just one value
go env GOOS GOARCH      # print several
go env -w GOFLAGS="-mod=mod"   # persistently set a value (writes to Go's env config file)
```

### The variables you'll actually use

| Variable | Meaning |
|---|---|
| `GOROOT` | Where the Go installation itself lives (standard library, compiler, tools). |
| `GOPATH` | The legacy Go workspace root; today mainly used for the module cache and `go install` targets. |
| `GOBIN` | Where `go install` places binaries (defaults to `$GOPATH/bin`). |
| `GOOS` / `GOARCH` | The target operating system / architecture for the **next** build (see [lesson 08](../08-go-build) on cross-compilation). |
| `GOMODCACHE` | Where downloaded module source code is cached on disk. |
| `GOTOOLCHAIN` | Controls whether Go auto-downloads a newer toolchain when `go.mod` requires one (see [lesson 02](../02-go-version-and-toolchain)). |

### Why not just read `os.Getenv`?

```go
os.Getenv("GOPATH")
```

This reads the **raw** environment variable — which is frequently **empty**, because most
developers never set `GOPATH` explicitly; Go just falls back to a sensible default
(`~/go`) internally. `go env GOPATH`, by contrast, always shows the value Go is **actually
using**, default or not. This is a subtle but important distinction: don't infer Go's
configuration by reading shell environment variables yourself — ask `go env`.

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Printf("runtime.GOROOT() : %s\n", runtime.GOROOT())
```

`runtime.GOROOT()` is one of the few Go-environment values directly queryable **from inside a
running Go program** — the others (`GOPATH`, `GOBIN`, etc.) are tooling-level concepts, not
runtime ones, which is why the rest of this lesson points you at the `go env` command instead.

## ▶️ How to Run

```bash
cd level-00-getting-started/14-go-env-command
go run main.go
go env GOROOT GOPATH GOOS GOARCH
```

## ✅ Expected Output (shape)

```
=== From inside the running program ===
----------------------------------
runtime.GOROOT() : /usr/lib/go-1.22
os.Getenv("GOPATH") : "" (raw shell value — may be empty even if Go has a default)
os.Getenv("GOOS")   : "" (raw shell value — usually empty; GOOS is normally implicit)

=== Compare against the command line ===
Run these in your terminal and compare:
  go env GOROOT
  go env GOPATH
  go env GOOS GOARCH
```

## 🧠 Key Takeaways

- `go env` shows Go's **effective** configuration, including computed defaults.
- Raw `os.Getenv("GOPATH")` may be empty even when Go has a real, working default value.
- `runtime.GOROOT()` is queryable from code; most other Go env settings are tooling-level only.
- `go env -w` persistently changes a setting without exporting a shell variable every session.

## 🛠️ Try It Yourself

1. Run `go env` with no arguments and read through the full list once, start to finish.
2. Temporarily set `GOOS=windows` in your shell (`export GOOS=windows` / `set GOOS=windows`) and
   re-run `go env GOOS` — see it reflect your override, then unset it.
3. Find `GOMODCACHE` on your machine with `go env GOMODCACHE` and look inside — that's where
   every module you've ever downloaded lives on disk.

## ⚠️ Common Mistakes

- Reading `os.Getenv("GOPATH")` from Go code and assuming an empty string means "no module
  cache" — it almost always just means "using the default."
- Manually exporting `GOROOT` in your shell "just in case" — modern Go installations detect this
  automatically and manual overrides are a common source of confusing breakage.
