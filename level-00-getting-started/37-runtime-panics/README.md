# 37 — Runtime Panics

## 🎯 Learning Objectives

- Explain the difference between a returned `error` and a `panic`.
- Recognize the most common causes of a runtime panic.
- Use `defer` + `recover` to turn a panic into a normal, handleable error.
- Write a test that deliberately triggers and catches a panic.

## 📖 Concept

Go has **two** distinct ways for something to go wrong:

| Mechanism | Meaning | How to handle it |
|---|---|---|
| Returned `error` | An expected, "this can normally fail" condition | Check `if err != nil`, covered in [lesson 38](../38-reading-error-messages) |
| `panic` | An unexpected condition the program can't sensibly continue from | `recover()`, or let the program crash |

A panic **unwinds the call stack** — each function returns immediately, running its deferred
calls as it goes, until either a `recover()` catches it or it reaches `main` and crashes the
whole program, exiting with status code 2 (see [lesson 39](../39-exit-status)).

### The four panics you'll hit constantly

```go
var s []int
_ = s[0]                       // panic: index out of range [0] with length 0

var m map[string]int
m["key"] = 1                    // panic: assignment to entry in nil map

var p *int
_ = *p                          // panic: invalid memory address or nil pointer dereference

x := 10
_ = x / 0                       // panic: integer divide by zero (compile error for CONSTANT 0, but runtime panic for a variable that happens to be 0)

var i interface{} = "hello"
n := i.(int)                    // panic: interface conversion: interface {} is string, not int
```

Recognizing "index out of range", "nil pointer dereference", "nil map", and "interface
conversion" by their exact wording is the fastest way to jump straight to the real bug, rather
than reading the whole stack trace line by line every time.

### Recovering: `defer` + `recover`

```go
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    result = riskyDivide(a, b)
    return result, nil
}
```

Three rules that make this pattern work, and that are worth internalizing precisely:

1. `recover()` only has an effect when called **directly inside a deferred function** — calling
   it anywhere else (even one level of indirection away) just returns `nil` and does nothing.
2. The deferred function runs **regardless** of whether a panic happened — `recover()` returns
   `nil` in the normal, no-panic case, which is why the `if r != nil` check matters.
3. Because `result` and `err` are **named return values**, the deferred function can set `err`
   and have it actually take effect on the function's return — an anonymous return wouldn't allow
   this.

## 🔍 Code Walkthrough (`main.go` and the test file)

`riskyDivide` is a deliberately "unsafe" function — real code often calls into something like it
(your own code, or a library) without controlling whether it panics. `safeDivide` is the pattern
you'd wrap around it at a boundary where you want failures to be recoverable instead of fatal —
e.g. one request handler in a server shouldn't take the whole process down.

```go
func TestRiskyDivide_Panics(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Fatal("expected riskyDivide(10, 0) to panic, but it did not")
        }
    }()
    riskyDivide(10, 0)
    t.Fatal("unreachable: riskyDivide should have panicked before this line")
}
```

This test **deliberately** triggers a panic and recovers from it right there in the test, to
prove the panic really happens — the final `t.Fatal` after the call is intentionally unreachable
if `riskyDivide` behaves as expected, and exists as a safety net in case it doesn't.

## ▶️ How to Run

```bash
cd level-00-getting-started/37-runtime-panics
go run main.go
go test -v ./...
```

## ✅ Expected Output

```
10 / 2 = 5
Handled gracefully: recovered from panic: runtime error: integer divide by zero

See the README for the most common causes of a Go panic, and why
recover() only works when called directly inside a deferred function.
```

## 🧠 Key Takeaways

- `error` is for expected failure; `panic` is for "this should be impossible."
- The four classic panics: index out of range, nil pointer dereference, nil map write,
  and bad interface type assertion.
- `recover()` only works when called directly inside a deferred function — no exceptions.
- Named return values let a deferred `recover` block actually change what the function returns.

## 🛠️ Try It Yourself

1. Add a fifth "classic panic" function of your own, wrapping one of the other three examples
   above (nil map, nil pointer, or bad type assertion) the same way `safeDivide` wraps
   `riskyDivide`.
2. Remove the `defer func() { recover() ... }()` block from `safeDivide` entirely and run
   `go run main.go` — watch the program crash with an actual, unrecovered panic and stack trace.
3. Put the `defer`'s `recover()` call inside a **regular** (non-deferred) helper function instead,
   call that helper from within the `defer`, and confirm `recover()` no longer catches anything —
   proving rule #1 above for yourself.

## ⚠️ Common Mistakes

- Recovering from **every** panic everywhere, turning genuine bugs into silently-swallowed
  errors — recover at deliberate boundaries (a request handler, a worker goroutine), not
  reflexively around every function.
- Calling `recover()` outside of a `defer`, or inside a function called *by* the deferred
  function rather than directly inside it — in both cases, it silently does nothing, which is
  easy to misdiagnose as "recover doesn't work."
