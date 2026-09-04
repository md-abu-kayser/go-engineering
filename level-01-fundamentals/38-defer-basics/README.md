# 38 — defer Basics

## 🎯 Learning Objectives

- Use `defer` to schedule a function call for when the surrounding function returns.
- Understand the classic gotcha: a deferred call's **arguments** are evaluated immediately, at
  the `defer` statement itself — only the call is postponed.
- Use a closure to defer something that should read a variable's **latest** value instead.

## 📖 Concept

`defer` schedules a function call to run **later** — specifically, right before the enclosing
function actually returns, regardless of how many lines of code come after the `defer` statement:

```go
fmt.Println("1. first")
defer fmt.Println("3. last")
fmt.Println("2. second")
// prints: 1, 2, 3 — the deferred call runs only once main() is about to return
```

This is commonly used for cleanup: closing a file, unlocking a mutex, closing a database
connection — code that needs to run **no matter how** the function exits (normal return, or even
a panic, previewed in [lesson 40](../40-panic-basics)).

### The classic gotcha: arguments are evaluated immediately

```go
n := 1
defer fmt.Printf("n was %d\n", n) // n's CURRENT VALUE (1) is captured RIGHT NOW
n = 100
// when the deferred call finally runs, it still prints "n was 1" — not 100!
```

This is genuinely surprising the first time you hit it: only the **function call itself** is
postponed — its **arguments** are evaluated at the moment `defer` is written, exactly like any
other function call's arguments would be. If `n` changes afterward, the deferred call has no way
to know; it already captured `n`'s value when `defer` executed.

### The fix, when you need the latest value: wrap it in a closure

```go
m := 1
defer func() {
    fmt.Printf("m is %d\n", m) // m is read WHEN THE CLOSURE RUNS, not when defer was written
}()
m = 100
// prints "m is 100" — because the closure reads m at the time it actually executes
```

A closure with no arguments defers the **entire function body**, which means the variable
reference inside it (`m`) is resolved only when that body actually runs — capturing whatever
`m`'s value is by then, not whatever it was back at the `defer` line.

## 🔍 Code Walkthrough (`main.go`)

```go
n := 1
defer fmt.Printf("deferred: n was %d AT THE TIME OF defer (not later)\n", n)
n = 100
```

`fmt.Printf`'s arguments — including `n`'s value, `1` — are evaluated **immediately**, right here,
even though the actual call to `Printf` doesn't happen until `main` returns. Changing `n`
afterward has zero effect on what gets printed.

```go
m := 1
defer func() {
    fmt.Printf("deferred (closure): m is %d ...\n", m)
}()
m = 100
```

Here, nothing is evaluated immediately except the closure's own creation — `m` is looked up
**inside** the closure body, which only runs later, by which point `m` has already changed to
`100`.

## ▶️ How to Run

```bash
cd level-01-fundamentals/38-defer-basics
go run main.go
```

## ✅ Expected Output

```
=== defer Basics ===
----------------------------------
1. This prints first.
2. This prints second.

--- Arguments are evaluated immediately, at defer time ---
n is now 100 (changed AFTER the defer statement ran)

--- Using a closure to capture the LATEST value instead ---
m is now 100
deferred (closure): m is 100 — the value AT THE TIME defer actually RUNS
deferred: n was 1 AT THE TIME OF defer (not later)
3. This prints LAST — scheduled by defer, runs when main() returns.
```

(Every deferred call prints only at the very end, in reverse of the order they were deferred —
that reverse ordering is the specific subject of [lesson 39](../39-defer-lifo-order).)

## 🧠 Key Takeaways

- `defer` schedules a call to run when the surrounding function returns, not immediately.
- A deferred call's **arguments** are evaluated right away, at the `defer` statement — only the
  call itself is postponed.
- Wrap a deferred call in a no-argument closure if you need it to read a variable's value **at
  the time it actually runs**, rather than at the time `defer` was written.
- `defer` is the standard tool for cleanup that must run regardless of how a function exits.

## 🛠️ Try It Yourself

1. Add a second `defer fmt.Println(n)` right after changing `n` to `100`, and confirm it captures
   `100` — since it's a **separate** `defer` statement, evaluated at that later point.
2. Rewrite the `m` closure example without a closure (plain `defer fmt.Printf(..., m)`), and
   confirm it now prints `1`, matching the original gotcha rather than the closure's fix.
3. Predict what order all four deferred statements in `main.go` will actually print in, before
   running it — then check your answer.

## ⚠️ Common Mistakes

- Assuming a deferred call's arguments reflect the state of the program when the call actually
  *runs* — they don't; they're captured immediately, at the `defer` line itself.
- Forgetting that a closure changes this — wrapping the same logic in `func() { ... }()` shifts
  variable evaluation to run-time instead of defer-time, which is easy to forget mid-refactor.
