# 41 — recover Basics

## 🎯 Learning Objectives

- Use `recover()` to convert a panic into a normal, checkable `error` instead of a crash.
- Know that `recover()` called with no panic in progress simply returns `nil` — always safe.
- Master the precise, easy-to-miss rule: `recover()` only catches a panic when called **directly**
  inside a deferred function — one level of indirection through a helper function defeats it.

## 📖 Concept

> **Related, but different lens:** [level 00's lesson on runtime panics]
> (../../level-00-getting-started/37-runtime-panics) already introduced this same `defer` +
> `recover` pattern with a `safeDivide` example. This lesson goes one step further into
> `recover`'s **precise placement rule** — a subtlety that lesson didn't need to cover — verified
> directly below rather than just asserted.

### The standard pattern: convert a panic into an error

```go
func safeCall(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    fn()
    return nil
}
```

This is the shape behind nearly every real-world use of `recover` in Go: wrap a risky call in a
deferred function, check `recover()`'s return value, and turn a non-nil result into an ordinary
`error` — using named return values (`err`) so the deferred function's assignment actually takes
effect on the surrounding function's return, exactly as covered in
[level 00's panic lesson](../../level-00-getting-started/37-runtime-panics).

### `recover()` with nothing panicking: always safe, always `nil`

```go
r := recover() // nothing is panicking right now
// r is nil — completely harmless to call
```

Calling `recover()` when there's no panic in progress is always safe and simply returns `nil` —
it's never an error to call `recover()`  "just in case," which is exactly what makes the pattern
above safe to wrap around any function unconditionally.

### THE precise rule: recover() must be called DIRECTLY inside a deferred function

This is the subtlety worth internalizing carefully:

```go
func recoverHelper() any {
    return recover() // recover() is called here, inside recoverHelper...
}

func demonstrateRecoverMustBeDirect() {
    defer func() {
        r := recoverHelper() // ...but recoverHelper is CALLED BY the deferred function,
        // it isn't the deferred function itself. r is ALWAYS nil here, even
        // during a genuine, in-progress panic.

        real := recover() // called DIRECTLY, right here — THIS one works.
    }()
    panic("...")
}
```

`recover()` only has an effect when the function **directly and literally** executing it is the
one Go deferred — one level of indirection through a normal function call (`recoverHelper()`)
completely defeats it, even though a panic is genuinely unwinding through that exact call stack
at that exact moment. This isn't a guess — it's empirically confirmed by this lesson's own output
below: the same panic first fails to be caught via the helper, then **is** caught by a second,
direct `recover()` call immediately afterward, in the very same deferred function invocation.

## 🔍 Code Walkthrough (`main.go`)

```go
err = safeCall(func() {
    panic("deliberate panic for this demo")
})
fmt.Printf("safeCall (panicking) -> err = %v\n", err)
```

`safeCall` returns a normal `error` here — the panic never escapes to `main` at all; from the
caller's perspective, this looks exactly like any other function that returned an error, with the
crash fully contained.

```go
r := recoverHelper()
fmt.Printf("recoverHelper() during a real panic: %v (nil — did NOT catch it)\n", r)

real := recover()
fmt.Printf("recover() called DIRECTLY in this deferred func: %v (this ONE works)\n", real)
```

Both lines run inside the **same** deferred function call, moments apart — the only difference is
*where* `recover()` was actually executed (inside `recoverHelper`, versus directly here). That
difference alone is what determines success or failure, which is exactly the point this lesson
exists to make concrete.

## ▶️ How to Run

```bash
cd level-01-fundamentals/41-recover-basics
go run main.go
```

## ✅ Expected Output

```
=== recover Basics ===
----------------------------------
  running safely, no panic here
safeCall (no panic)  -> err = <nil>
safeCall (panicking) -> err = recovered from panic: deliberate panic for this demo

--- recover() with no panic in progress ---
recover() called with nothing panicking: <nil>

--- recover() must be called DIRECTLY inside the deferred function ---
recoverHelper() during a real panic: <nil> (nil — did NOT catch it)
recover() called DIRECTLY in this deferred func: a panic used to demonstrate recover's placement rule (this ONE works)

Program continued normally — the direct recover() call caught the panic.
```

## 🧠 Key Takeaways

- `recover()` converts an in-progress panic into a normal return value — the standard pattern
  wraps it in a deferred closure and assigns to a named `error` return.
- `recover()` with no panic in progress is always safe, always returns `nil`.
- `recover()` only catches a panic when called **directly** inside the deferred function itself —
  routing through a helper function, even one line of indirection, defeats it completely.
- A failed indirect `recover()` attempt doesn't "use up" the panic — a subsequent direct call, in
  the same deferred function, can still catch it, as this lesson's own output demonstrates.

## 🛠️ Try It Yourself

1. Remove the second, direct `recover()` call from `demonstrateRecoverMustBeDirect` (in a scratch
   copy) and run the program — watch it crash for real, since nothing ultimately recovered the
   panic.
2. Write your own two-line "helper defeats recover" demonstration from scratch, without looking
   at this lesson's code, to confirm you've internalized the rule.
3. Use `safeCall` to wrap a function that triggers one of [lesson 40](../40-panic-basics)'s
   built-in panic causes (e.g. an out-of-range slice index) instead of an explicit `panic(...)`
   call, and confirm it's caught identically.

## ⚠️ Common Mistakes

- Extracting `recover()` into a shared helper function "to avoid repetition," not realizing this
  silently breaks it — `recover()` must be written directly inside each deferred function that
  needs it; it cannot be factored out.
- Assuming a failed recovery attempt (like the indirect one in this lesson) permanently
  "consumes" the panic — it doesn't; the panic is still fully in progress until something
  actually, successfully recovers it.
