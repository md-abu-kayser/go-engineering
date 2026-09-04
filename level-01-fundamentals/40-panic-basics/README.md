# 40 — panic Basics

## 🎯 Learning Objectives

- Understand `panic` as a distinct control-flow mechanism, not just "an error that crashes."
- Panic with a **custom value type**, not only strings — and inspect that value later.
- Confirm that deferred calls still run **during** a panic's stack unwinding, even before
  anything actually recovers it.

## 📖 Concept

> **Related, but different lens:** [level 00's lesson on runtime panics]
> (../../level-00-getting-started/37-runtime-panics) already covered panic/recover from a
> "debugging and defensive coding" angle, with a `safeDivide` example. This lesson covers the
> same feature from a pure **language-mechanics** angle — what panic actually *is*, as a control
> flow construct — using different examples. Both are worth knowing; neither repeats the other.

### `panic` accepts ANY value, not just strings

```go
panic("a plain string")                     // fine
panic(fmt.Errorf("a formatted error"))       // fine
panic(lookupError{Index: 10, Length: 3})      // ALSO fine — any value at all
```

`panic`'s parameter type is `any` — literally anything can be a panic value. Using a **structured**
value (a custom struct, as this lesson does with `lookupError`) instead of a bare string lets
whatever eventually recovers the panic inspect real, structured fields — not just parse a
human-readable message back apart.

### Panicking stops normal execution and starts unwinding

```go
func riskyLookup(items []string, index int) string {
    if index < 0 || index >= len(items) {
        panic(lookupError{Index: index, Length: len(items)})
    }
    return items[index] // never reached if we panicked above
}
```

Once `panic` runs, the current function stops executing immediately — no more of its code runs,
including the `return` statement below the panic. Control immediately begins **unwinding**: the
current function exits (abnormally), then its caller, then that caller's caller, and so on — each
one exiting in turn — until either something **recovers** the panic, or it reaches the very top
of the program with nothing having recovered it, crashing the whole process.

### Deferred calls still run during unwinding

```go
func runWithCleanup() {
    defer fmt.Println("cleanup ran")
    fmt.Println("about to panic")
    panic("something went wrong")
}
```

Even though `runWithCleanup` never reaches a normal `return`, its `defer` **still runs** — `defer`
statements are guaranteed to execute during unwinding, precisely so cleanup code (closing a file,
releasing a lock) still happens even when a function exits abnormally via panic rather than a
normal return. This is one of the most important reasons `defer` and `panic` are designed to work
together.

## 🔍 Code Walkthrough (`main.go`)

```go
func() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("recovered a panic value of type %T: %v\n", r, r)
        }
    }()
    riskyLookup(items, 10)
}()
```

This recovers `lookupError{Index: 10, Length: 3}` and prints both its **type**
(`main.lookupError`) and its formatted message (via the `Error()` method it implements) — proof
the panic value survived all the way to the recover point as a real, structured value, not just
text.

```go
func runWithCleanup() {
    defer fmt.Println("  runWithCleanup: cleanup ran ...")
    fmt.Println("  runWithCleanup: about to panic")
    panic("something went wrong inside runWithCleanup")
}
```

Notice in the actual output, `"about to panic"` prints **before** `"cleanup ran"` — exactly the
order you'd expect: the function runs normally up to the `panic` call, and only **then** does the
deferred cleanup fire, as unwinding begins.

## ▶️ How to Run

```bash
cd level-01-fundamentals/40-panic-basics
go run main.go
```

## ✅ Expected Output

```
=== panic Basics ===
----------------------------------
riskyLookup(items, 1) = "b" (no panic — valid index)

--- panic with a CUSTOM value type, not just a string ---
recovered a panic value of type main.lookupError: index 10 out of range for length 3

--- Deferred calls run DURING unwinding, before anything recovers ---
  runWithCleanup: about to panic
  runWithCleanup: cleanup ran (deferred, even though we're about to panic)

See the README: an UNRECOVERED panic keeps unwinding all the way up and,
if nothing ever recovers it, crashes the whole program.
```

## 🧠 Key Takeaways

- `panic` accepts any value (`any`) — a custom struct works exactly as well as a string.
- Panicking immediately stops normal execution and begins unwinding the call stack.
- Deferred calls still run during unwinding, in every function panic passes through — this is
  precisely why `defer` is the standard tool for cleanup that must happen regardless of how a
  function exits.
- An unrecovered panic keeps unwinding all the way to the top and crashes the whole program —
  see [lesson 41](../41-recover-basics) for how to actually stop that.

## 🛠️ Try It Yourself

1. Add a second field to `lookupError` (e.g. `Items []string`) and print it too once recovered.
2. Add a second `defer` inside `runWithCleanup`, above the existing one, and confirm both run
   during unwinding, in LIFO order ([lesson 39](../39-defer-lifo-order)) — even though the
   function panicked instead of returning normally.
3. Remove the `recover()` call around `runWithCleanup()`'s invocation entirely, and — in a
   scratch copy only — run the program to see what an actually-unrecovered panic's crash output
   looks like.

## ⚠️ Common Mistakes

- Assuming `panic` must always be given a string or an `error` — any value works, and a
  structured custom type is often more useful to whatever eventually recovers it.
- Forgetting that `defer` still runs during a panic — code assuming cleanup **only** happens on a
  normal return path can leave resources open exactly when a panic occurs, which is the opposite
  of what you want.
