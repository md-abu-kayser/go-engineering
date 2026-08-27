# 40 — Standard Input

## 🎯 Learning Objectives

- Read input line-by-line with `bufio.Scanner`.
- Understand `os.Stdin` as just another `io.Reader`.
- Correctly detect EOF (end of input) instead of treating it as an error.
- Feed test input to a program non-interactively, from the command line.

## 📖 Concept

`os.Stdin` is Go's handle to the process's standard input stream — and importantly, it's just a
regular `*os.File`, which means it satisfies `io.Reader` like almost every other source of bytes
in Go (files, network connections, in-memory buffers). Anything written to work with an
`io.Reader` works with `os.Stdin` for free.

### The idiomatic way to read lines: `bufio.Scanner`

```go
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
    line := scanner.Text()
    // use line
}
if err := scanner.Err(); err != nil {
    // a REAL error occurred (rare for stdin)
}
```

- `scanner.Scan()` returns `true` once per available line, `false` when input is exhausted.
- `scanner.Text()` returns the current line, **without** its trailing newline.
- `scanner.Err()` reports a genuine read error — reaching the end of input normally is **not**
  one; `Scan()` simply returns `false` and `Err()` returns `nil`.

### Feeding input without typing interactively

Since this program reads from stdin, you can pipe input into it instead of typing by hand — the
standard way to test or script an interactive program:

```bash
printf "hello\nworld\n" | go run main.go
echo "single line" | go run main.go
go run main.go < somefile.txt
```

All three redirect **something else** to be `os.Stdin`, exactly as if you'd typed it and pressed
`Ctrl+D` (EOF) at the end.

## 🔍 Code Walkthrough (`main.go`)

```go
scanner := bufio.NewScanner(os.Stdin)
lineNum := 0
for scanner.Scan() {
    lineNum++
    line := scanner.Text()
    fmt.Printf("line %d (%d chars, uppercased): %s\n", lineNum, len(line), strings.ToUpper(line))
}
```

The loop naturally ends when input runs out (`Scan()` returns `false`) — there's no explicit "if
this is the last line" check needed; that's precisely what makes `bufio.Scanner` pleasant to use
compared to lower-level reading.

## ▶️ How to Run

```bash
cd level-00-getting-started/40-standard-input
printf "hello\nworld\n" | go run main.go
```

## ✅ Expected Output

```
=== Standard Input ===
----------------------------------
Reading lines from stdin until EOF. Try:
  printf "hello\nworld\n" | go run main.go

line 1 (5 chars, uppercased): HELLO
line 2 (5 chars, uppercased): WORLD

Read 2 line(s) total.
```

## 🧠 Key Takeaways

- `os.Stdin` is a regular `io.Reader` — everything you know about reading files applies.
- `bufio.Scanner` is the idiomatic way to read line-by-line input.
- Reaching EOF is normal termination, not an error — check `scanner.Err()` for genuine failures.
- Piping (`|`) or redirecting (`<`) lets you feed input without typing interactively.

## 🛠️ Try It Yourself

1. Run the program with no piped input at all (just `go run main.go`, then type a line and press
   `Ctrl+D`) and confirm it behaves the same as piped input.
2. Change `bufio.NewScanner` to use `bufio.ScanWords` instead of the default line-splitting
   (`scanner.Split(bufio.ScanWords)`) and observe it now processes one **word** at a time instead
   of one line.
3. Feed it a file with `go run main.go < somefile.txt` instead of a pipe, and confirm identical
   behavior — both are just different ways of connecting something else to `os.Stdin`.

## ⚠️ Common Mistakes

- Treating the end of input as an error condition — check `scanner.Err()` specifically; a `nil`
  error after the loop ends means everything worked normally.
- Forgetting a program reading from `os.Stdin` with no piped input will simply **hang**, waiting
  for you to type something — that's expected behavior, not a bug, when run interactively.
