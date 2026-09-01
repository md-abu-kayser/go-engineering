# 27 — switch Statements

## 🎯 Learning Objectives

- Write a `switch` with multiple values in a single `case`.
- Know that Go's `switch` does **not** fall through by default — the single biggest difference
  from C/Java/JavaScript's `switch`.
- Use the explicit `fallthrough` keyword for the rare cases you genuinely want that behavior.

## 📖 Concept

```go
switch day {
case "Saturday", "Sunday":
    return "Weekend"
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    return "Weekday"
default:
    return "Unknown"
}
```

### Multiple values per case

```go
case "Saturday", "Sunday":
```

A single `case` can list several comma-separated values — matching **any** of them — without
needing to repeat the case body or fall through from a separate case.

### No fallthrough by default — this is the big one

In C, Java, and JavaScript, forgetting a `break` at the end of a `switch` case causes execution
to silently "fall through" into the next case — a famously common source of bugs. **Go's `switch`
does not do this.** Each case is complete on its own; once a matching case's body finishes, the
`switch` ends automatically:

```go
switch n {
case 2:
    fmt.Println("two")   // execution stops HERE
case 3:
    fmt.Println("three")  // never reached, even though there's no "break"
}
```

No `break` statement is needed (or even idiomatic) at the end of a Go `case` — the language
simply doesn't fall through, so there's nothing to prevent.

### `fallthrough`: opt-in, for the rare case you actually want it

```go
switch n {
case 2:
    fmt.Println("two")
    fallthrough    // explicitly continue into the NEXT case
case 3:
    fmt.Println("three")
}
```

`fallthrough` is a real keyword, used when you genuinely want C-style behavior for one specific
transition — but it's opt-in and explicit, and it only continues into the **immediately next**
case, unconditionally (it does **not** re-evaluate that next case's own condition).

## 🔍 Code Walkthrough (`main.go`)

```go
switch n {
case 2:
    fmt.Println("two")
    fallthrough
case 3:
    fmt.Println("three (reached via explicit fallthrough, not because n == 3)")
case 4:
    fmt.Println("four (NOT reached — fallthrough only continues ONE case)")
}
```

This deliberately prints `case 3`'s body even though `n` is `2`, not `3` — proof that
`fallthrough` really does force execution into the next case regardless of that case's own
condition — while `case 4` is correctly **not** reached, showing `fallthrough` only propagates
one single step, not indefinitely.

## ▶️ How to Run

```bash
cd level-01-fundamentals/27-switch-statements
go run main.go
```

## ✅ Expected Output

```
=== switch Statements ===
----------------------------------
Saturday   -> Weekend
Monday     -> Weekday
Someday    -> Unknown

--- No fallthrough by default ---
two

--- Explicit fallthrough (opt-in, rare) ---
two
three (reached via explicit fallthrough, not because n == 3)
```

## 🧠 Key Takeaways

- A single `case` can match multiple, comma-separated values.
- Go's `switch` does **not** fall through by default — no `break` needed, and none is idiomatic.
- `fallthrough` is an explicit, opt-in keyword for the rare case you want C-style behavior —
  and it only continues exactly one case forward, unconditionally.

## 🛠️ Try It Yourself

1. Add a `case 4:` to the first (no-fallthrough) switch and confirm reaching `case 2` never
   touches it — matching the "no fallthrough by default" behavior.
2. Chain two `fallthrough` statements in a row (case 1 falls to case 2, which falls to case 3) and
   confirm all three bodies run in sequence.
3. Try putting `fallthrough` as the very last case in a `switch` and read the compiler error —
   there's nothing left to fall through *into*.

## ⚠️ Common Mistakes

- Writing `break` at the end of every Go `case` out of habit from another language — it's
  harmless but completely unnecessary; Go never needed it in the first place.
- Assuming Go's `switch` behaves like C's and adding defensive `break`s "just in case" — the
  behavior is genuinely different (no fallthrough), not just differently spelled.
