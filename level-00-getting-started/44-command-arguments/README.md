# 44 — Command Arguments

## 🎯 Learning Objectives

- Read raw command-line arguments with `os.Args`.
- Define and parse named flags with the standard `flag` package.
- Distinguish flags from positional arguments in the same invocation.

## 📖 Concept

### The raw layer: `os.Args`

```go
os.Args[0]   // the program's own path/name
os.Args[1:]  // everything the user actually typed after it
```

`os.Args` is a plain `[]string` — no parsing, no flag syntax, just exactly what was typed,
whitespace-separated by the shell. Every other command-line argument mechanism in Go (including
`flag`) is built on top of this.

### The common case: named flags with `flag`

Manually parsing `--name=Gopher --shout` out of a raw `[]string` is tedious and easy to get
subtly wrong. The standard `flag` package handles this for you:

```go
name := flag.String("name", "Gopher", "who to greet")
shout := flag.Bool("shout", false, "uppercase the greeting")
flag.Parse()   // must be called AFTER all flag.Xxx() declarations
```

- `flag.String("name", "Gopher", "...")` returns a `*string`, pre-populated with the **default**
  value (`"Gopher"`), which gets overwritten if `-name=...` was actually passed.
- The third argument to each `flag.Xxx` call is a description, shown automatically if the user
  passes `-h` or `-help`.
- After `flag.Parse()`, `flag.Args()` returns whatever's left over — the **positional**
  arguments that weren't consumed as flags.

### Flag syntax `flag` understands

```bash
go run main.go -name=Alice -shout
go run main.go -name Alice -shout
go run main.go -name=Alice extra positional args
```

Both `-name=Alice` and `-name Alice` work; `--name` (double dash) is also accepted as a synonym
for `-name` (single dash) — Go's `flag` package deliberately doesn't distinguish them, unlike some
other languages' conventions.

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Printf("os.Args (raw)      : %v\n", os.Args)
```

This line runs **before** `flag.Parse()` is even called, so it shows you the truly raw input —
useful for seeing exactly what `flag` is about to process on your behalf.

```go
flag.Parse()
```

Must come **after** every `flag.String`/`flag.Bool`/etc. declaration and **before** you read any
of their values — calling it too early means later-declared flags won't be recognized.

## ▶️ How to Run

```bash
cd level-00-getting-started/44-command-arguments
go run main.go -name=Alice -shout extra stuff
```

## ✅ Expected Output

```
=== Command Arguments ===
----------------------------------
os.Args (raw)      : [/tmp/go-buildXXXX/b001/exe/main -name=Alice -shout extra stuff]
program name       : /tmp/go-buildXXXX/b001/exe/main
positional args    : [-name=Alice -shout extra stuff]

HELLO, Alice!!!
remaining (non-flag) arguments: [extra stuff]
```

(`os.Args[0]`'s exact path depends on whether you used `go run` or a built binary — see
[lesson 07](../07-go-run).)

## 🧠 Key Takeaways

- `os.Args[0]` is the program path; `os.Args[1:]` is everything the user typed after it.
- `flag.Xxx(name, default, usage)` declares a named flag and returns a pointer to its value.
- `flag.Parse()` must run after every flag declaration and before you read any flag's value.
- `flag.Args()` returns whatever positional arguments were left after flags were consumed.

## 🛠️ Try It Yourself

1. Run the program with no arguments at all and confirm the defaults (`"Gopher"`, not shouting)
   are used.
2. Run it with `-h` and read the auto-generated usage text `flag` builds from your descriptions.
3. Add a third flag, `-times int`, defaulting to `1`, and print the greeting that many times in a
   loop.

## ⚠️ Common Mistakes

- Reading a flag's value (dereferencing its `*string`/`*bool`) **before** calling `flag.Parse()` —
  you'll only ever see the default, never what the user actually passed.
- Declaring a flag **after** calling `flag.Parse()` — that flag will never be recognized, since
  parsing already happened.
