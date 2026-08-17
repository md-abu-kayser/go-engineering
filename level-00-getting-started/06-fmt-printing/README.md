# 06 — fmt & Printing

## 🎯 Learning Objectives

- Know the difference between `Print`, `Println`, and `Printf`.
- Use the most common formatting verbs confidently.
- Build a string with `Sprintf` instead of printing it immediately.
- Understand `Fprintf` and the idea of writing to a "destination" other than the console.

## 📖 Concept

The `fmt` package's functions come in families, each with a consistent naming pattern:

| Function                   | Destination                                        | Adds spaces between args?         | Adds trailing newline? |
| -------------------------- | -------------------------------------------------- | --------------------------------- | ---------------------- |
| `fmt.Print(...)`           | stdout                                             | only between non-string operands  | ❌                     |
| `fmt.Println(...)`         | stdout                                             | ✅ always                         | ✅                     |
| `fmt.Printf(fmt, ...)`     | stdout                                             | only what your format string adds | ❌ (add `\n` yourself) |
| `fmt.Sprintf(fmt, ...)`    | returns a `string`                                 | only what your format string adds | ❌                     |
| `fmt.Fprintf(w, fmt, ...)` | any `io.Writer` (a file, a buffer, `os.Stdout`, …) | only what your format string adds | ❌                     |

The pattern is consistent across the whole standard library: an `S` prefix means "return a
string instead of printing," and an `F` prefix means "write to the writer I give you instead of
stdout." You'll see this same `Print` / `Sprint` / `Fprint` pattern again with `log`, `errors`,
and elsewhere.

### Common formatting verbs

| Verb   | Meaning                                                  | Example    |
| ------ | -------------------------------------------------------- | ---------- |
| `%s`   | string                                                   | `"Gopher"` |
| `%d`   | base-10 integer                                          | `15`       |
| `%f`   | floating point (default precision)                       | `1.750000` |
| `%.2f` | floating point, 2 decimal places                         | `1.75`     |
| `%t`   | boolean                                                  | `true`     |
| `%v`   | the "default" format for any value — great for debugging | `Gopher`   |
| `%T`   | the **Go type** of the value                             | `int`      |
| `%q`   | a double-quoted, escaped string                          | `"Gopher"` |

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Print("No newline here...")
fmt.Print(" ...still the same line.\n")
```

`Print` does **not** add a newline for you — the second call continues on the same line as the
first, which is why we add `\n` manually.

```go
fmt.Printf("%%d (integer) : %d\n", age)
```

Notice `%%d` — a literal percent sign is written as `%%`, since `%` alone always starts a verb.

```go
summary := fmt.Sprintf("%s is %d years old and %.2fm tall.", name, age, height)
```

`Sprintf` builds and **returns** the formatted string instead of printing it — useful whenever
you want to construct a message to log, return from a function, or use somewhere other than the
console immediately.

## ▶️ How to Run

```bash
cd level-00-getting-started/06-fmt-printing
go run main.go
```

## ✅ Expected Output

```
=== Print vs Println ===
No newline here... ...still the same line.
Println adds a newline automatically.

=== Printf verbs ===
%s (string)              : Gopher
%d (integer)             : 15
%f (float, default)      : 1.750000
%.2f (float, 2 decimals) : 1.75
%t (boolean)             : true
%v (default format)      : Gopher
%T (type of the value)   : int
%q (quoted string)       : "Gopher"

=== Sprintf: build a string instead of printing it ===
Gopher is 15 years old and 1.75m tall.

=== Fprintf: write to a specific destination ===
Written explicitly to standard output.
```

## 🧠 Key Takeaways

- `Println` adds spaces and a newline automatically; `Print` does neither reliably.
- `Printf` gives you full control via format verbs — learn `%s %d %f %v %T %q` first.
- `Sprintf` returns a string; `Fprintf` writes to any `io.Writer`, not just the console.
- `%v` is your best friend for quickly debugging a value of any type.

## 🛠️ Try It Yourself

1. Add a variable of your own (any type) and print it with `%v` and `%T` to see both its value
   and its type.
2. Try `%5d` and `%-5d` with a small number and observe how width/alignment verbs work.
3. Use `fmt.Fprintf` to write formatted text into a `strings.Builder` instead of `os.Stdout`.

## ⚠️ Common Mistakes

- Forgetting the `\n` after `Printf` — unlike `Println`, it never adds one for you.
- Mismatching the number of `%` verbs and arguments — Go will print `%!d(MISSING)` or similar
  instead of crashing, so watch the output carefully.
- Using `%d` on a float or `%s` on a non-string without `%v` — Go will print an error marker
  like `%!d(float64=1.75)` rather than silently doing the wrong thing.
