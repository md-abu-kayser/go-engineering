# 35 — Stack Traces

## 🎯 Learning Objectives

- Read a Go stack trace and identify the sequence of calls that led to a given point.
- Print a stack trace from running code with `runtime/debug.PrintStack()`.
- Get the same information interactively from Delve with the `stack` command.

## 📖 Concept

A **stack trace** lists every function call currently "in progress," from the most recent (where
execution is right now) down to where the program started. It answers the question: *"how did
we get here?"* — essential both while debugging and while reading a panic message
([lesson 37](../37-runtime-panics)).

### From inside running code: `runtime/debug.PrintStack()`

```go
import "runtime/debug"

debug.PrintStack()
```

This prints the current goroutine's call stack to stderr immediately, with no debugger attached
— useful for logging exactly where in the code some interesting event occurred, in production or
in a test, not just during an interactive debugging session.

### From Delve: the `stack` command

```bash
dlv debug .
```

```
(dlv) break main.printStackOnce
(dlv) continue
(dlv) stack
0  0x... in main.printStackOnce
1  0x... in main.factorial
2  0x... in main.factorial
3  0x... in main.factorial
4  0x... in main.factorial
5  0x... in main.factorial
6  0x... in main.main
7  0x... in runtime.main
```

Each numbered frame is one function call, deepest first. Notice `main.factorial` appears **five
times** — once per level of recursion — which is exactly the kind of thing a stack trace makes
immediately visible that a single breakpoint wouldn't.

### Reading the shape of a trace

- **Top of the trace** = where you are right now (or where the panic happened).
- **Bottom of the trace** = the program's entry point (`main.main`, then Go's own runtime
  startup).
- **Repeated entries** = recursion, or the same function called from multiple places.

## 🔍 Code Walkthrough (`main.go`)

```go
func factorial(n int) int {
    if n <= 1 {
        printStackOnce(n)
        return 1
    }
    return n * factorial(n-1)
}
```

Recursion is the clearest way to produce a stack trace worth reading — `factorial(5)` calls
`factorial(4)`, which calls `factorial(3)`, and so on, giving `debug.PrintStack()` (and Delve's
`stack` command) a real, multi-level call chain to display.

## ▶️ How to Run

```bash
cd level-00-getting-started/35-stack-traces
go run main.go
```

## ✅ Expected Output (shape)

The exact memory addresses, absolute file paths, and `runtime/debug` internal line numbers below
will differ on your machine — what matters is the **shape**: a `goroutine` header, a couple of
frames inside `runtime/debug` itself, then `main.printStackOnce`, then `main.factorial` repeated
once per level of recursion, ending in `main.main`.

```
=== Stack trace at the deepest recursive call ===
goroutine 1 [running]:
runtime/debug.Stack()
	.../runtime/debug/stack.go:24 +0x5e
runtime/debug.PrintStack()
	.../runtime/debug/stack.go:16 +0x13
main.printStackOnce(...)
	.../35-stack-traces/main.go:36 +0x6a
main.factorial(0x1)
	.../35-stack-traces/main.go:18 +0x19
main.factorial(0x2)
	.../35-stack-traces/main.go:21 +0x35
main.factorial(0x3)
	.../35-stack-traces/main.go:21 +0x35
main.factorial(0x4)
	.../35-stack-traces/main.go:21 +0x35
main.factorial(0x5)
	.../35-stack-traces/main.go:21 +0x35
main.main()
	.../35-stack-traces/main.go:40 +0x1c

5! = 120

See the README for reading this trace, and for the `dlv stack` equivalent.
```

Note that `debug.PrintStack()` writes to **stderr**, while the rest of this program's output goes
to stdout ([lesson 42](../42-standard-error) covers exactly this split) — in a normal terminal
both streams appear interleaved together as above, but redirecting stdout alone
(`go run main.go 2>/dev/null`) would show only the non-trace lines.

## 🧠 Key Takeaways

- A stack trace is an ordered list of "who called whom," deepest call first.
- `runtime/debug.PrintStack()` prints one from running code, no debugger required.
- Delve's `stack` command shows the equivalent, interactively, while paused.
- Repeated function names in a trace are your signal to look for recursion.

## 🛠️ Try It Yourself

1. Run `go run main.go` and count how many `main.factorial` frames appear — confirm it matches
   the argument you called `factorial` with (5).
2. Add a `debug.PrintStack()` call inside `main` itself (not inside `factorial`) and compare how
   much shorter that trace is.
3. Attach `dlv debug .`, break on `main.printStackOnce`, and run `stack` — compare it line-for-line
   against the `debug.PrintStack()` output from a normal run.

## ⚠️ Common Mistakes

- Reading a stack trace bottom-to-top out of habit from other contexts — Go's convention (and
  most languages') is top = most recent, bottom = program start.
- Ignoring repeated frames as "noise" — in real bugs, an unexpectedly deep or infinitely repeating
  stack (often ending in "goroutine stack exceeds …") is itself the diagnosis: uncontrolled
  recursion.
