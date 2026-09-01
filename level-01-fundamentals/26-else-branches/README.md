# 26 — else Branches

## 🎯 Learning Objectives

- Chain `if` / `else if` / `else` correctly.
- Know Go's mandatory brace-placement rule for `else` — and why it exists.
- Recognize the idiomatic "early return, no else" pattern and when it improves readability.

## 📖 Concept

```go
if score >= 90 {
    return "A"
} else if score >= 80 {
    return "B"
} else {
    return "F"
}
```

Straightforward chaining, same as most C-family languages — with one Go-specific syntax rule
worth knowing precisely.

### `else` must start on the same line as the previous `}`

```go
if cond {
    // ...
} else {          // "else" MUST be on the SAME LINE as the closing brace before it
    // ...
}
```

```go
if cond {
    // ...
}
else {              // COMPILE ERROR — "else" on its own line is NOT valid Go
    // ...
}
```

This isn't a style preference — it's enforced by Go's automatic semicolon insertion rules. A `}`
at the end of a line gets an implicit semicolon inserted after it unless something continues the
statement on the **same** line, so `else` on its own line effectively terminates the `if`
statement early, making the following `else` a syntax error with no matching `if`. `gofmt`
formats this correctly for you automatically, which is exactly why you'll rarely encounter this
error in practice — but it's worth understanding *why* the rule exists.

### The idiomatic alternative: early returns, no `else` at all

```go
func validateAge(age int) (ok bool, reason string) {
    if age < 0 {
        return false, "age cannot be negative"
    }
    if age > 150 {
        return false, "age is not plausible"
    }
    return true, ""
}
```

Once a branch **returns**, there's nothing left to be "else" about — the code after it only runs
if that branch didn't trigger, which is exactly what an `else` would have meant anyway, just with
less nesting. This "guard clause" style — handle the exceptional/invalid cases early and return,
leaving the main logic un-nested at the end — is extremely common and generally preferred in
idiomatic Go, especially as the number of conditions grows.

## 🔍 Code Walkthrough (`main.go`)

```go
func grade(score int) string {
    if score >= 90 {
        return "A"
    } else if score >= 80 {
    ...
```

`grade` uses a genuine `if`/`else if`/`else` chain because every branch is doing the **same
kind** of thing (returning a letter grade) — there's no natural "early return, different
concern" structure here, which is exactly when `else` chains remain the clearest choice.

```go
func validateAge(age int) (ok bool, reason string) {
    if age < 0 {
        return false, "age cannot be negative"
    }
    if age > 150 {
        return false, "age is not plausible"
    }
    return true, ""
}
```

By contrast, `validateAge` is checking for **distinct, unrelated** failure conditions, each of
which should immediately end the function — exactly the shape where early returns read more
clearly than nested `else if` chains would.

## ▶️ How to Run

```bash
cd level-01-fundamentals/26-else-branches
go run main.go
```

## ✅ Expected Output

```
=== else Branches ===
----------------------------------
score  95 -> grade A
score  82 -> grade B
score  71 -> grade C
score  40 -> grade F

--- Early return instead of else ---
age   -5: invalid (age cannot be negative)
age   30: valid
age  200: invalid (age is not plausible)
```

## 🧠 Key Takeaways

- `else` must begin on the same line as the preceding `}` — enforced by Go's semicolon insertion
  rules, and handled automatically by `gofmt`.
- `if`/`else if`/`else` chains work well when every branch is doing the same *kind* of thing.
- Early returns with no `else` at all are often clearer for handling distinct failure conditions —
  once a branch returns, everything after it implicitly only runs in the "else" case.

## 🛠️ Try It Yourself

1. Deliberately put `else` on its own line in a scratch copy and read the exact compiler error —
   then run `gofmt` on the file and watch it get fixed automatically.
2. Rewrite `validateAge` using an `if`/`else if`/`else` chain instead of early returns, and decide
   which version reads more clearly to you.
3. Add a third validation rule to `validateAge` (e.g. reject exactly `0`) using the early-return
   style, and confirm it fits naturally into the existing pattern.

## ⚠️ Common Mistakes

- Manually formatting `else` on its own line out of habit from another language's style — let
  `gofmt` handle brace placement; don't fight it by hand.
- Defaulting to `else` chains even when every branch is really handling an unrelated, independent
  concern — consider whether early returns would flatten the logic and make it easier to follow.
