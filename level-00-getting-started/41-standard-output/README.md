# 41 — Standard Output

## 🎯 Learning Objectives

- Understand that `fmt.Println` writes to `os.Stdout` by default.
- Write to `os.Stdout` explicitly, three different ways, and see they're equivalent.
- Redirect a program's standard output from the shell.

## 📖 Concept

Every process has (at least) two separate output streams: **standard output** (`stdout`), for a
program's normal, intended output, and **standard error** (`stderr`,
[lesson 42](../42-standard-error)), for diagnostics and error messages. This lesson focuses on
`stdout`.

### Three equivalent ways to write to stdout

```go
fmt.Println("...")                          // implicit: fmt.Println always targets os.Stdout
fmt.Fprintln(os.Stdout, "...")                // explicit: same destination, spelled out
os.Stdout.WriteString("...\n")                 // most direct: os.Stdout is just an *os.File
```

`os.Stdout` is a regular `*os.File` value, open for writing, connected to whatever the operating
system considers this process's standard output — usually your terminal, unless redirected.

### Redirecting stdout from the shell

```bash
go run main.go > output.txt
```

The `>` operator tells your shell to connect this program's `os.Stdout` to a **file** instead of
your terminal — the Go program itself needs zero code changes to support this; it's purely a
shell-level redirection of the same stream.

```bash
go run main.go | wc -l
```

Piping (`|`) connects one program's `stdout` to another program's `stdin`
([lesson 40](../40-standard-input)) — this is exactly how Unix-style command chaining works, and
Go programs participate in it automatically, with no special code required.

## 🔍 Code Walkthrough (`main.go`)

The three writes in `main()` are functionally identical — same bytes, same destination — shown
side by side so you can see that `fmt.Println` isn't doing anything magical: it's a convenience
wrapper around exactly the same `os.Stdout` you can write to directly.

## ▶️ How to Run

```bash
cd level-00-getting-started/41-standard-output
go run main.go
go run main.go > /tmp/output.txt && cat /tmp/output.txt
```

## ✅ Expected Output

```
Written via fmt.Println (implicitly to stdout)
Written via fmt.Fprintln(os.Stdout, ...) (explicitly)
Written directly via os.Stdout.WriteString

See the README for redirecting this program's stdout to a file,
and for why that's different from redirecting stderr (lesson 42).
```

## 🧠 Key Takeaways

- `fmt.Println` always targets `os.Stdout` — there's no hidden behavior beyond that.
- `os.Stdout` is a plain `*os.File`; anything that accepts an `io.Writer` can target it.
- `>` redirects a program's stdout to a file; `|` connects it to another program's stdin.
- Go programs need no special code to support shell redirection — it's transparent.

## 🛠️ Try It Yourself

1. Run `go run main.go > /tmp/output.txt`, then `cat /tmp/output.txt` and confirm every line made
   it into the file, in order.
2. Try `go run main.go | wc -l` and confirm the line count matches what you'd expect from the
   program's output.
3. Write your own one-line program that writes a JSON string via `os.Stdout.WriteString`, and
   pipe it into a tool like `jq` (if available) to confirm the output is genuinely usable by
   another program.

## ⚠️ Common Mistakes

- Assuming redirecting stdout (`>`) also captures error messages — it doesn't; those typically go
  to stderr instead, which is exactly the subject of [lesson 42](../42-standard-error).
- Writing debug/diagnostic output with `fmt.Println` (stdout) instead of to stderr — this pollutes
  a program's "real" output, making it harder for other tools to consume via a pipe.
