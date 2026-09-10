# 57 — The main Function

## 🎯 Learning Objectives

- Recall `main`'s mandatory, unchangeable signature: no parameters, no return value.
- Understand `main` is called by the Go **runtime**, not by any code you write.
- Know that a program's entire lifetime is bounded by `main` — it starts when `main` starts, and
  ends the moment `main` returns.

## 📖 Concept

> **Related, different angle:** [level 00's lesson 04]
> (../../level-00-getting-started/04-package-main) covered `package main` — the *package-level*
> significance of the name `main`. This lesson is specifically about the **function** `main()`
> itself, once you're already inside that special package.

### The signature is fixed — no exceptions

```go
func main() {
    ...
}
```

This exact shape — no parameters, no return type — is **mandatory**. Go will not compile
`func main(args []string)`, `func main() int`, or any other variation, no matter how natural it
might seem coming from a language where `main` conventionally receives arguments or returns an
exit code. If you need command-line arguments, [level 00's lesson 44]
(../../level-00-getting-started/44-command-arguments) already covered the actual mechanism:
`os.Args`, read from inside `main` (or anywhere else), not passed as a parameter to it. If you
need to signal a specific exit status, that's `os.Exit(code)`
([level 00's lesson 39](../../level-00-getting-started/39-exit-status)), not a return value.

### Called by the runtime, not by your own code

Every other function in this repository, across every lesson, has been called by **some line of
code you wrote** — `main()` calling `setup()`, one function calling another. `main()` itself is
different: nothing in your source code calls it. The Go runtime calls it automatically, once,
when the program starts — it's the designated **entry point**, not just a function that happens
to be named specially.

### The program's lifetime is bounded by `main`

The program begins running the moment `main` starts, and — critically — **ends** the moment
`main` returns, whether that's a normal fall-through return or an explicit `return` statement
partway through. (`os.Exit` is the one way to end the program *without* `main` returning at all,
bypassing even deferred calls — [level 00's lesson 39]
(../../level-00-getting-started/39-exit-status) covers that distinction in depth.)

## 🔍 Code Walkthrough (`main.go`)

```go
func setup() {
    fmt.Println("setup() ran — called explicitly, by main() itself")
}

func main() {
    ...
    setup()
    ...
}
```

`setup` is an entirely ordinary function — nothing about it is special. The contrast is the
point: `setup()` is called by `main()`, a line of code you can see right here; `main()` itself has
no equivalent calling line anywhere in this file, or anywhere in the whole program.

## ▶️ How to Run

```bash
cd level-01-fundamentals/57-main-function
go run main.go
```

## ✅ Expected Output

```
=== The main Function ===
----------------------------------
main() itself was called by the Go RUNTIME, not by any
line of code in this program — it's the program's entry point.
setup() ran — called explicitly, by main() itself

os.Args (main() can't take parameters, so this is the workaround): [/tmp/go-buildXXXX/b001/exe/main]

Once main() returns here, the entire program ends.
```

(`os.Args`'s exact printed path depends on whether you used `go run` or a built binary — see
[level 00's lesson 07](../../level-00-getting-started/07-go-run).)

## 🧠 Key Takeaways

- `func main()` — no parameters, no return type — is a mandatory, unchangeable signature.
- `main` is called once, automatically, by the Go runtime — never by your own source code.
- `os.Args` and `os.Exit` are how `main` receives input and signals an exit status, since it
  can't have parameters or a return value itself.
- The program's entire lifetime is bounded by `main`'s execution — it ends the moment `main`
  returns (or `os.Exit` is called anywhere).

## 🛠️ Try It Yourself

1. Try changing `main`'s signature to accept a parameter (e.g. `func main(x int)`) and read the
   compiler's exact rejection.
2. Add a second helper function called from `main`, and trace through the file to confirm `main`
   is the only function in it with no corresponding call site.
3. Add a `return` statement partway through `main` (inside an `if`) and confirm any code written
   after it — including in `setup`, if called afterward — never runs, since the whole program
   ends right there.

## ⚠️ Common Mistakes

- Trying to give `main` command-line-style parameters out of habit from another language — Go's
  `main` never takes any; use `os.Args` instead, from inside the function body.
- Assuming returning early from `main` behaves like returning from any other function (just
  "going back to the caller") — there is no caller to return to; it ends the entire program.
