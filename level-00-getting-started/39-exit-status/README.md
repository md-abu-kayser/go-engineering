# 39 — Exit Status

## 🎯 Learning Objectives

- Understand what a process's exit status is and who reads it.
- Use `os.Exit` to end a program with a specific, non-zero status.
- Know how an unrecovered panic's exit status differs from a deliberate `os.Exit`.
- Check a program's exit status from the shell.

## 📖 Concept

Every process, when it finishes, reports a single integer back to whatever started it — its
**exit status**. By strong convention across virtually every operating system:

- **`0`** means success.
- **Any non-zero value** means some kind of failure — the specific number is up to the program.

This is how shell scripts, CI pipelines, and `Makefile`s decide whether a command "worked" —
they check the exit status, not the text it printed.

### Setting it explicitly with `os.Exit`

```go
if err := validateAge(age); err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
}
```

`os.Exit(n)` terminates the program **immediately** with status `n` — critically, it does **not**
run any pending deferred functions. This is an important difference from a normal `return` out of
`main`, and worth remembering if you rely on `defer` for cleanup (closing files, flushing
buffers) — code that needs that cleanup to run should return a value from `main` and let it exit
normally, rather than calling `os.Exit` deep inside a function.

### What "falling off the end of `main`" does

If `main` simply returns without calling `os.Exit`, Go exits with status **0** automatically — you
only need `os.Exit` when you want a **non-zero** status, or want to skip deferred cleanup on
purpose (rare, and usually only appropriate right at program startup, before anything has
registered a `defer`).

### An unrecovered panic's exit status

If a panic reaches `main` without being recovered ([lesson 37](../37-runtime-panics)), Go prints
the panic message and stack trace to stderr and exits with status **2** — distinct from both a
"clean" `os.Exit(1)` and a normal successful exit, which is occasionally useful for distinguishing
"the program deliberately reported failure" from "the program crashed unexpectedly."

### A `go run` gotcha: it always reports `1` on failure

`go run` compiles and runs your program in one step, but it is itself a **separate process**
wrapping the one you wrote. If your program exits non-zero, `go run` prints
`exit status N` (telling you the real code) but then **always exits with status `1` itself** —
regardless of whether your program actually exited with `1`, `2`, `3`, or anything else. So
`echo $?` right after a `go run` invocation tells you *whether* it failed, but not the *specific*
code your program used. To see the true exit status, build a real binary first and run that
directly:

```bash
go build -o app .
./app
echo $?          # this is your program's ACTUAL exit status
```

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Fprintln(os.Stderr, "Error:", err)
os.Exit(1)
```

Notice the error message goes to `os.Stderr`, not `os.Stdout` — covered in depth in
[lesson 42](../42-standard-error). Separating "the actual output" from "diagnostic/error text"
is exactly why both streams exist, and it's a habit worth having from your very first programs.

## ▶️ How to Run

```bash
cd level-00-getting-started/39-exit-status
go build -o app .
./app
echo "Exit status: $?"
```

## ✅ Expected Output

```
Age 15 is valid.

See the README for how to check this program's exit status from your shell,
and how it would differ if validateAge had failed instead.
Exit status: 0
```

## 🧠 Key Takeaways

- Exit status `0` means success; anything else means some form of failure, by convention.
- `os.Exit(n)` ends the program immediately with status `n`, **skipping** any pending `defer`s.
- Falling off the end of `main` normally is equivalent to `os.Exit(0)`.
- An unrecovered panic exits with status `2`, distinct from a deliberate `os.Exit(1)`.
- `go run` always exits `1` on any failure, no matter the underlying code — build a real binary
  (`go build`) and run it directly if you need the true exit status.
- `$?` in your shell (`%errorlevel%` on Windows `cmd`) shows the most recently finished command's
  exit status.

## 🛠️ Try It Yourself

1. Change `age` to `200` (invalid), run `go build -o app . && ./app`, and check `echo $?` —
   confirm it's `1`.
2. Add a `defer fmt.Println("cleanup")` at the top of `main`, trigger the `os.Exit(1)` path, and
   notice "cleanup" is **never printed** — proof `os.Exit` skips deferred calls.
3. Force an unrecovered panic instead (e.g. `var s []int; _ = s[5]`), then `go build -o app .` and
   `./app; echo $?` — confirm it reports `2`, not `1`. Then try `go run main.go; echo $?` on the
   same broken code and notice it reports `1` instead — the gotcha above, seen firsthand.

## ⚠️ Common Mistakes

- Calling `os.Exit` deep inside a function that other code expects to clean up after (closing a
  database connection, flushing a log) — deferred cleanup anywhere in the call stack above that
  point simply never runs.
- Assuming a program "worked" because it printed something that looked right — always check the
  actual exit status in scripts (`$?`, or `if command; then ...`), not just the presence of output.
