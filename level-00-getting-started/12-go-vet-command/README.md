# 12 — go vet

## 🎯 Learning Objectives

- Understand the difference between what the **compiler** checks and what `go vet` checks.
- Recognize the kinds of real bugs `go vet` catches.
- Run `go vet` as part of a normal development workflow.

## 📖 Concept

The Go compiler only rejects code that is **syntactically or type invalid**. But some code is
perfectly valid Go — it compiles and runs — while still almost certainly being a bug. `go vet`
is a static analysis tool that catches exactly this category of mistake.

### A classic example: mismatched Printf verbs

This code **compiles and runs**, but is wrong:

```go
name := "Gopher"
fmt.Printf("Hello, %d\n", name) // %d expects an integer, name is a string
```

Running `go vet` on a file containing this line reports:

```
./main.go:5:2: Printf format %d has arg name of wrong type string
```

The compiler doesn't catch this because `Printf`'s arguments are just `...any` — type-safe as
far as the compiler is concerned. `go vet` understands `fmt`'s format-string *conventions*
specifically, on top of normal type checking.

### Other things `go vet` catches

- Struct tags with invalid syntax (e.g. a malformed `json:"..."` tag).
- Copying a `sync.Mutex` by value (which breaks the lock).
- Unreachable code after a `return`.
- Suspicious comparisons, like comparing a pointer to itself.

## ▶️ How to Run

```bash
cd level-00-getting-started/12-go-vet-command
go vet ./...
go run main.go
```

`go vet` produces **no output** when everything is clean — silence means success, the same
convention as `gofmt -l`.

## ✅ Expected Output

```
Hello, Gopher is 15 years old

See the README for an example of the kind of bug `go vet` catches —
this file is intentionally correct so it builds and vets cleanly.
```

(and no output at all from `go vet ./...`)

## 🧠 Key Takeaways

- The compiler checks types and syntax; `go vet` checks for **likely bugs in valid code**.
- Mismatched `Printf`-style format verbs are the single most common thing `go vet` catches.
- `go vet` produces no output when everything is clean, just like `gofmt -l`.
- Many editors (including VS Code's Go extension) run `go vet` automatically in the background.

## 🛠️ Try It Yourself

1. In a scratch copy of `main.go`, change `%s` to `%d` in the `Printf` call (leaving `name` as a
   string) and run `go vet ./...` — read the exact error message it produces.
2. Fix it back and confirm `go vet ./...` is silent again.
3. Look up `go tool vet help` (or `go vet -h`) to see the full list of checks `go vet` runs by
   default.

## ⚠️ Common Mistakes

- Ignoring `go vet` warnings because "it still compiled" — that's precisely the category of bug
  it exists to catch.
- Assuming `go vet` catches *all* bugs — it only catches specific, well-understood suspicious
  patterns, not general logic errors.
