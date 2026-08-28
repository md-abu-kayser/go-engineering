# 43 — Environment Basics

## 🎯 Learning Objectives

- Read an environment variable with `os.Getenv`.
- Distinguish "unset" from "set to an empty string" with `os.LookupEnv`.
- Set an environment variable for the current process with `os.Setenv`.
- List every environment variable visible to a process with `os.Environ`.

## 📖 Concept

Environment variables are one of the most common ways to configure a program without changing
its code — API keys, feature flags, deployment mode (`dev`/`staging`/`prod`), and more are
routinely passed this way, especially in containerized deployments.

### Reading: `os.Getenv` vs `os.LookupEnv`

```go
home := os.Getenv("HOME")           // returns "" if HOME isn't set
value, ok := os.LookupEnv("APP_MODE") // ok is false if APP_MODE isn't set
```

`os.Getenv` is convenient but has a real ambiguity: it returns `""` both when a variable is
**genuinely unset** and when it's **set to an empty string**. `os.LookupEnv`'s second return
value (`ok`) resolves that ambiguity — use it whenever "was this actually set?" matters, not just
"what's a reasonable default if not?"

### Setting: `os.Setenv`

```go
os.Setenv("APP_MODE", "debug")
```

This sets the variable **only within the current process** (and any child processes it later
starts) — it has no effect on your shell, your terminal session, or any other running program.
It's mainly useful for tests that need a controlled environment, or for propagating a computed
value to a subprocess you're about to launch.

### Listing everything: `os.Environ`

```go
for _, kv := range os.Environ() {
    fmt.Println(kv)   // each entry looks like "KEY=value"
}
```

Returns every variable visible to the process as `"KEY=value"` strings — useful for debugging
what a program actually sees, especially when its environment was set up somewhere else
(a `Dockerfile`, a CI config, a systemd unit).

## 🔍 Code Walkthrough (`main.go`)

```go
value, ok := os.LookupEnv("APP_MODE")
if !ok {
    fmt.Println(`os.LookupEnv("APP_MODE")     : not set`)
}
```

`APP_MODE` almost certainly isn't set in your shell, which is exactly the point — this branch
demonstrates the "genuinely unset" case concretely, before the program sets it itself a few lines
later with `os.Setenv` and reads it back.

## ▶️ How to Run

```bash
cd level-00-getting-started/43-environment-basics
go run main.go
APP_MODE=production go run main.go   # now APP_MODE IS set before the program even starts
```

## ✅ Expected Output (first run, `APP_MODE` unset beforehand)

```
=== Environment Basics ===
----------------------------------
os.Getenv("HOME")           : "/root"
os.LookupEnv("APP_MODE")     : not set
after os.Setenv, LookupEnv   : "debug"

Total environment variables visible to this process: 12
```

(Your `HOME` value and the total variable count will differ based on your machine.)

## 🧠 Key Takeaways

- `os.Getenv` can't distinguish "unset" from "empty string" — use `os.LookupEnv` when that matters.
- `os.Setenv` only affects the current process and anything it spawns — never your shell.
- `os.Environ()` lists every variable visible to the process, as `"KEY=value"` strings.
- Environment variables are a process-local, inherited-from-parent configuration mechanism, not a
  persistent or global setting.

## 🛠️ Try It Yourself

1. Run the program with `APP_MODE=production go run main.go` and confirm `LookupEnv` reports it as
   set, with the value `"production"`, **before** the program's own `os.Setenv` call runs.
2. Print out the full `os.Environ()` list (loop over it and `fmt.Println` each entry) and skim
   what your shell actually exposes to Go programs.
3. Set an environment variable to an **empty string** explicitly
   (`APP_MODE= go run main.go` on macOS/Linux) and confirm `os.LookupEnv` reports `ok = true`
   with an empty value — genuinely different from not being set at all.

## ⚠️ Common Mistakes

- Using `os.Getenv` and treating `""` as "not configured," when it might actually mean
  "explicitly configured to be empty" — use `os.LookupEnv` if that distinction matters for your
  program's logic.
- Expecting `os.Setenv` to persist beyond the current process, or to change your shell's
  environment — it does neither.
