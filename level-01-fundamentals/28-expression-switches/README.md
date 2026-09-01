# 28 — Expression Switches

## 🎯 Learning Objectives

- Use a "tagless" `switch` (no expression after the keyword) as an alternative to a long
  `if`/`else if`/`else` chain.
- Recognize that a tagless `switch` is exactly equivalent to `switch true`.
- Use a `switch` with an init statement, the same way `if` supports one ([lesson 25](../25-if-with-init)).

## 📖 Concept

[Lesson 27](../27-switch-statements) covered `switch` with a **value** to match against
(`switch day { case "Saturday": ... }`). Go also supports a **tagless** form — no expression at
all after `switch` — where each `case` is a full boolean expression:

```go
switch {
case score >= 90:
    return "A"
case score >= 80:
    return "B"
default:
    return "F"
}
```

This is precisely equivalent to `switch true { case score >= 90: ... }` — Go evaluates each
`case`'s boolean expression in order, top to bottom, and runs the body of the **first** one that's
`true` (or `default`, if none match).

### Tagless switch vs. if/else-if chains: a genuine style choice

[Lesson 26](../26-else-branches)'s `grade` function and this lesson's `classify` function
implement the **exact same logic**, two different ways:

```go
// if/else-if/else (lesson 26)
if score >= 90 {
    return "A"
} else if score >= 80 {
    return "B"
} else {
    return "F"
}

// tagless switch (this lesson)
switch {
case score >= 90:
    return "A"
case score >= 80:
    return "B"
default:
    return "F"
}
```

Both are fully idiomatic Go — this is genuinely a matter of taste and readability, not a rule
with a single correct answer. Many developers reach for the tagless `switch` once there are more
than two or three conditions, since it avoids the visually nested `} else if {` chain.

### `switch` also supports an init statement

```go
switch length := len("Gopher"); {
case length > 10:
    fmt.Println("long")
case length > 3:
    fmt.Println("medium")
default:
    fmt.Println("short")
}
```

Exactly the same pattern as `if init; condition` from [lesson 25](../25-if-with-init) — `length`
here is scoped to the whole `switch`, visible in every `case`, and invisible outside it.

## 🔍 Code Walkthrough (`main.go`)

```go
switch {
case score >= 90:
```

The bare `switch {` (nothing between `switch` and `{`) is what makes this "tagless" — contrast
with [lesson 27](../27-switch-statements)'s `switch day {`, which has an explicit value to match
against.

## ▶️ How to Run

```bash
cd level-01-fundamentals/28-expression-switches
go run main.go
```

## ✅ Expected Output

```
=== Expression Switches ===
----------------------------------
score  95 -> grade A (via tagless switch)
score  82 -> grade B (via tagless switch)
score  71 -> grade C (via tagless switch)
score  40 -> grade F (via tagless switch)

--- Same logic, if/else-if/else style (lesson 26) vs switch (this lesson) ---
Both are equally valid Go — this is a genuine STYLE choice, not a rule.

--- switch with an init statement ---
medium
```

## 🧠 Key Takeaways

- A tagless `switch { case cond1: ... }` is equivalent to `switch true`, evaluating boolean
  expressions top to bottom.
- Choosing between a tagless `switch` and an `if`/`else if` chain is a genuine style decision —
  both are fully idiomatic.
- `switch` supports an init statement (`switch init; { ... }`), scoped identically to `if`'s.

## 🛠️ Try It Yourself

1. Reorder `classify`'s cases so a lower threshold comes before a higher one (e.g. `score >= 70`
   before `score >= 90`), and observe how a `95` now incorrectly grades as `"C"` — proof that case
   order matters, exactly like an `if`/`else if` chain.
2. Rewrite `validateAge` from [lesson 26](../26-else-branches) as a tagless `switch` instead of
   early returns, and compare readability.
3. Add a `case length == 6:` before `case length > 3:` in the init-statement example and confirm
   it correctly takes priority for `"Gopher"` (6 characters).

## ⚠️ Common Mistakes

- Assuming a tagless `switch`'s cases are somehow evaluated all at once or in some other order —
  they're checked strictly top to bottom, exactly like a chain of `if`/`else if`, and the first
  match wins.
- Treating "tagless switch vs. if/else-if" as a rule with one correct answer — it's a readability
  choice; use whichever makes a given piece of logic clearest to the next reader.
