# 45 — Named Results

## 🎯 Learning Objectives

- Name a function's return values in its signature, for documentation and convenience.
- Use a "naked return" — a bare `return` with no values — enabled by named results.
- Understand the genuine gotcha: a deferred function can still modify a named result **after** a
  `return` statement has already run.

## 📖 Concept

Return values can be **named**, right in the function signature:

```go
func divide(a, b int) (result int, err error) {
```

This does two things at once: it gives each return value a documented **name** (immediately
useful — `go doc`, editor tooltips, and anyone reading the signature see `result` and `err`
instead of two anonymous `int`/`error`s), and it pre-declares `result`/`err` as ordinary local
variables, usable throughout the function body.

### The naked return

Because named results are already declared variables, a bare `return` (no values listed) sends
back whatever they currently hold:

```go
func divide(a, b int) (result int, err error) {
    if b == 0 {
        err = fmt.Errorf("cannot divide %d by zero", a)
        return // naked return — sends back result (still 0) and err (just set)
    }
    result = a / b
    return // naked return — result is set, err is still nil
}
```

Naked returns are a genuine style choice, and opinions differ — some find them concise for short
functions; many style guides discourage them in longer functions, where it's not obvious at a
glance what's actually being returned without scrolling back up to check the named variables'
current values. An **explicit** `return result, err` always still works too, named results or not.

### The gotcha: `defer` can still modify a named result after `return`

This is the subtle, genuinely important part: a `return` statement doesn't immediately end the
function — it sets the return values, and **then** any deferred calls run, and **only after
those** does the function actually exit. If a deferred function assigns to a named result, that
assignment **overwrites** whatever `return` had just set:

```go
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            result = -1 // overwrites result AFTER `return` already ran
        }
    }()
    result = a / b // panics if b == 0
    return
}
```

This exact mechanism is precisely what makes the `defer` + `recover` error-handling pattern
(covered in [lesson 41](../41-recover-basics) and
[level 00's panic lesson](../../level-00-getting-started/37-runtime-panics)) actually work — the
deferred `recover` needs to be able to change what the function ultimately returns, and named
results are what makes that possible.

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Printf("safeDivide(10, 0) = %d (deferred recover OVERWROTE the named result)\n", safeDivide(10, 0))
```

Dividing by zero here panics on `result = a / b`, before the `return` statement is ever reached
directly — but the deferred function still runs during unwinding
([lesson 40](../40-panic-basics)), recovers the panic, and sets `result = -1`, which is what the
function ultimately returns instead of crashing.

## ▶️ How to Run

```bash
cd level-01-fundamentals/45-named-results
go run main.go
```

## ✅ Expected Output

```
=== Named Results ===
----------------------------------
divide(10, 2) = 5, err = <nil>
divide(10, 0) = 0, err = cannot divide 10 by zero

--- Named results as documentation ---
rectangleStats(4, 6) = area 24, perimeter 20

--- The defer + named result gotcha ---
safeDivide(10, 2) = 5 (normal case)
safeDivide(10, 0) = -1 (deferred recover OVERWROTE the named result)
```

## 🧠 Key Takeaways

- Named return values document what each result means, directly in the function signature.
- A naked `return` sends back whatever the named results currently hold — a style choice, not a
  requirement.
- A `return` statement sets the return values, but the function doesn't truly exit until every
  deferred call has also run — and those calls can still modify named results afterward.
- This defer-can-still-modify-the-result mechanism is exactly what makes `defer` + `recover` work
  as an error-handling pattern.

## 🛠️ Try It Yourself

1. Change `safeDivide`'s deferred function to print `result`'s value **before** setting it to
   `-1`, and confirm it shows `0` — proof the panic happened before `result = a / b` could run.
2. Rewrite `divide` using explicit `return result, err` statements instead of naked returns, and
   confirm the behavior is identical.
3. Write your own function with a named result and a deferred function that increments it by 1
   right before returning, and confirm the caller sees the incremented value.

## ⚠️ Common Mistakes

- Assuming `return` immediately ends a function — deferred calls still run afterward, and (with
  named results) can still change what's actually returned.
- Overusing naked returns in long functions, where a reader has to scroll back to the signature
  to remember what's actually being returned — reserve them for short, simple functions where the
  named results are still fresh in view.
