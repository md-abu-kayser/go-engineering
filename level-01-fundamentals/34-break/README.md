# 34 — break

## 🎯 Learning Objectives

- Use `break` to exit a loop before its natural end.
- Recognize the classic gotcha: `break` inside a `switch` nested in a loop exits only the
  `switch`, not the loop.
- Know that `break` always targets its **nearest** enclosing `for`, `switch`, or `select`.

## 📖 Concept

[Lesson 30](../30-for-loop) already introduced `break` in its simplest form: inside a loop,
`break` ends that loop immediately.

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        break // exits the for loop right here
    }
}
```

### The gotcha: `break` targets its NEAREST enclosing construct

`break` doesn't mean "exit the loop" specifically — it means "exit the nearest enclosing `for`,
`switch`, or `select`," whichever one is closer. If a `switch` sits inside a loop, a `break`
written inside that `switch` exits the **switch**, not the loop around it:

```go
for i := 0; i < 5; i++ {
    switch {
    case i == 2:
        break // exits the SWITCH — the for loop is completely unaffected
    }
    // execution continues here, and the loop keeps running
}
```

This genuinely surprises people arriving from languages where `switch` doesn't even need an
explicit `break` to avoid fallthrough ([lesson 27](../27-switch-statements) already covered that
Go's `switch` doesn't fall through by default) — but the *reason* `break` exists at all inside a
`switch` is precisely to exit that switch early, before reaching the end of its case body, which
is a different thing from exiting an outer loop.

### The fix, when you actually need it: labeled break

If you genuinely need to break the **outer loop** from inside a nested `switch` (or a nested
loop), you need a **labeled break** — covered fully in [lesson 36](../36-labeled-break).

## 🔍 Code Walkthrough (`main.go`)

```go
for i := 0; i < 5; i++ {
    switch {
    case i == 2:
        fmt.Printf("  i=%d: breaking the SWITCH, not the loop\n", i)
        break
    default:
        fmt.Printf("  i=%d: default case\n", i)
    }
    fmt.Printf("  (loop continues after switch, i=%d)\n", i)
}
```

Notice the `"(loop continues after switch, i=%d)"` line prints on **every** iteration, including
`i == 2` — direct proof that `break` inside the `switch` case did not touch the surrounding `for`
loop at all.

## ▶️ How to Run

```bash
cd level-01-fundamentals/34-break
go run main.go
```

## ✅ Expected Output

```
=== break ===
----------------------------------
--- break in a plain loop ---
  i = 0
  i = 1
  i = 2

--- The gotcha: break inside switch-inside-loop ---
  i=0: default case
  (loop continues after switch, i=0)
  i=1: default case
  (loop continues after switch, i=1)
  i=2: breaking the SWITCH, not the loop
  (loop continues after switch, i=2)
  i=3: default case
  (loop continues after switch, i=3)
  i=4: default case
  (loop continues after switch, i=4)

See lesson 36 (labeled break) for how to ACTUALLY break the outer loop
from inside a nested switch or inner loop.
```

## 🧠 Key Takeaways

- `break` exits its **nearest** enclosing `for`, `switch`, or `select` — not necessarily "the loop."
- A `break` inside a `switch` nested in a loop only exits the switch; the loop keeps running.
- This is a genuine, common source of confusion — worth internalizing precisely, not just
  "knowing about it."
- [Lesson 36](../36-labeled-break) covers the actual fix: a labeled `break` targeting the loop
  specifically.

## 🛠️ Try It Yourself

1. Add a `case i == 4:` with its own `break`, and confirm the loop still runs to completion (all
   5 iterations) despite two separate switch-level breaks.
2. Remove the `switch` entirely and put the same `if i == 2 { break }` directly in the loop body
   instead — confirm this DOES stop the loop, contrasting directly with the switch-wrapped version.
3. Predict, before running it, what would happen if you nested a `for` loop inside a `switch`
   case, with a `break` inside that inner loop — which construct does it exit?

## ⚠️ Common Mistakes

- Assuming `break` inside a `switch`-inside-a-loop stops the loop — it doesn't; it only exits the
  switch, and the loop's next iteration proceeds normally.
- Adding an unnecessary `break` at the end of a `switch` case out of habit (from a language that
  needs it to prevent fallthrough) without realizing it does something different in Go — it's
  harmless there (Go's switch doesn't fall through anyway, [lesson 27](../27-switch-statements)),
  but genuinely changes behavior if placed inside a loop-wrapping switch where you actually
  intended to exit the loop.
