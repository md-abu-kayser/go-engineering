# 42 — Standard Error

## 🎯 Learning Objectives

- Understand why diagnostics and errors belong on `os.Stderr`, not `os.Stdout`.
- Write to `os.Stderr` with `fmt.Fprintln`/`fmt.Fprintf`.
- Redirect stdout and stderr **independently** from the shell.
- Know that Go's `log` package writes to stderr by default, and why that's the right default.

## 📖 Concept

[Lesson 41](../41-standard-output) covered `os.Stdout` — a program's real, intended output.
`os.Stderr` exists specifically for **diagnostics**: warnings, errors, progress messages —
anything that isn't the actual data/result the program produces. Keeping them separate means a
program's real output stays clean and pipeable, even while it's also reporting problems.

### Writing to stderr

```go
fmt.Fprintln(os.Stderr, "warning:", err)
fmt.Fprintf(os.Stderr, "error: %v\n", err)
```

Just like `os.Stdout`, `os.Stderr` is a plain `*os.File` — same `Fprintln`/`Fprintf` functions,
just a different destination.

### Why this separation matters in practice

Imagine a program that prints one price per line, but also warns about a couple of bad entries:

```bash
go run main.go
```
```
Keyboard   $49.99
warning: item "Mouse" has an invalid negative price: -100
Monitor    $159.99
```

Both streams land in your terminal together, interleaved — but they're independently
**redirectable**:

```bash
go run main.go > prices.txt          # only the real prices go to the file
go run main.go 2> warnings.txt        # only the warnings go to this file
go run main.go > prices.txt 2>&1       # merge both into prices.txt
go run main.go 2>/dev/null            # discard warnings entirely, keep only real output
```

If warnings had been mixed into stdout instead, `prices.txt` would contain a warning line mixed
in with actual prices — breaking anything downstream that expected clean, parseable data (a CSV
importer, another script piping this output onward).

### `log` writes to stderr by default

Go's standard `log` package (`log.Println`, `log.Printf`, `log.Fatal`) writes to `os.Stderr` by
default, not `os.Stdout` — a deliberate design choice, consistent with the same principle: log
messages are diagnostics, not a program's primary output.

## 🔍 Code Walkthrough (`main.go`)

```go
if err := processItem(name, price); err != nil {
    fmt.Fprintln(os.Stderr, "warning:", err)
    continue
}
fmt.Printf("%-10s $%.2f\n", name, float64(price)/100)
```

Notice the deliberate split: the **error path** writes to `os.Stderr` and `continue`s (skipping
that item); the **success path** uses ordinary `fmt.Printf`, which targets `os.Stdout`. This is
the shape you'll write constantly in real command-line tools.

## ▶️ How to Run

```bash
cd level-00-getting-started/42-standard-error
go run main.go
go run main.go 2>/dev/null
go run main.go 1>/dev/null
```

## ✅ Expected Output

Combined (normal run, both streams to your terminal):

```
Keyboard   $49.99
warning: item "Mouse" has an invalid negative price: -100
Monitor    $159.99
```

With `2>/dev/null` (stderr discarded — only real output remains):

```
Keyboard   $49.99
Monitor    $159.99
```

With `1>/dev/null` (stdout discarded — only the warning remains):

```
warning: item "Mouse" has an invalid negative price: -100
```

## 🧠 Key Takeaways

- `os.Stderr` is for diagnostics; `os.Stdout` is for a program's real, intended output.
- Both are independently redirectable (`>` for stdout, `2>` for stderr, `2>&1` to merge them).
- Go's `log` package targets stderr by default, matching this same convention.
- Keeping the split correct is what lets other tools reliably consume just your real output.

## 🛠️ Try It Yourself

1. Run all three commands in "How to Run" and confirm each one's output matches what's shown above.
2. Swap the `fmt.Fprintln(os.Stderr, ...)` call for a plain `fmt.Println(...)` (stdout) and rerun
   `go run main.go 2>/dev/null` — notice the warning **incorrectly** still shows up, since it's no
   longer actually going to stderr.
3. Replace the manual `fmt.Fprintln(os.Stderr, ...)` call with `log.Println(...)` instead, and
   confirm `2>/dev/null` still successfully hides it — proving `log` really does default to stderr.

## ⚠️ Common Mistakes

- Writing warnings/errors with `fmt.Println` out of habit — this silently defeats the entire
  purpose of having two separate streams.
- Assuming `2>&1` and `>` combined do what you want regardless of order — `> file.txt 2>&1` (stdout
  to the file, then stderr follows stdout) behaves differently from `2>&1 > file.txt` (stderr
  follows the *old* stdout, i.e. the terminal, then stdout moves to the file) — order matters in
  shell redirection.
