# 07 — Booleans

## 🎯 Learning Objectives

- Use `bool` and Go's three logical operators: `&&`, `||`, `!`.
- Observe short-circuit evaluation directly, not just as a claimed rule.
- Know that Go has no implicit conversion between `bool` and any other type.

## 📖 Concept

`bool` has exactly two values, `true` and `false`, and three operators:

| Operator | Meaning | 
|---|---|
| `&&` | logical AND — true only if both sides are true |
| `\|\|` | logical OR — true if either side is true |
| `!` | logical NOT — inverts a boolean |

### Short-circuit evaluation

```go
false && hasSideEffect(...)
```

With `&&`, if the **left** side is already `false`, Go never even evaluates the right side —
the overall result can only be `false` regardless, so there's no point. Similarly, with `||`, if
the left side is already `true`, the right side is never evaluated either.

This isn't just an optimization detail — it's a **guarantee** you can rely on, and it's commonly
used deliberately:

```go
if user != nil && user.IsActive() {
    // safe: IsActive() is only called if user != nil, avoiding a nil-pointer panic
}
```

If Go evaluated both sides unconditionally, this extremely common nil-check-then-use pattern
would panic whenever `user` was `nil` — short-circuiting is what makes it safe.

### No implicit conversions

Unlike some languages, Go's `bool` doesn't implicitly convert to or from any other type — there's
no "truthy" `0`/`1`, no non-empty-string-is-true. An `if` condition must be an actual `bool`
expression; `if x` where `x` is an `int` is simply a compile error, not "does `x` truthy-check."

## 🔍 Code Walkthrough (`main.go`)

```go
result := false && hasSideEffect("right side of &&", true)
```

`hasSideEffect` prints a line every time it's actually **called**. Because the left side is
`false`, `&&` never calls it at all — you can see this directly in the program's own output: the
"(evaluating ...)" line for that call simply never appears, proving short-circuiting rather than
just asserting it.

## ▶️ How to Run

```bash
cd level-01-fundamentals/07-booleans
go run main.go
```

## ✅ Expected Output

```
=== Booleans ===
----------------------------------
isRaining              : true
haveUmbrella           : false
!isRaining             : false
isRaining && haveUmbrella : false
isRaining || haveUmbrella : true

--- Short-circuit evaluation ---
false && hasSideEffect(...):
  result = false (right side was NEVER evaluated — see above, no line printed)
true || hasSideEffect(...):
  result = true (right side was NEVER evaluated — see above, no line printed)
true && hasSideEffect(...):
  (evaluating right side of && (this time it DOES run))
  result = true
```

## 🧠 Key Takeaways

- `bool` has exactly two values; `&&`/`||`/`!` are the only logical operators.
- `&&` skips evaluating its right side if the left side is already `false`; `||` skips it if the
  left side is already `true`.
- Short-circuiting is a language guarantee, routinely relied on for safe nil-checks.
- Go has no implicit truthy/falsy conversion — `if` requires a genuine `bool` expression.

## 🛠️ Try It Yourself

1. Swap the order in the nil-check-style example (`hasSideEffect(...) && false`) and confirm the
   side effect **does** run this time, since it's now on the left.
2. Write your own two-condition `if` using `&&` where the left condition prevents a would-be panic
   on the right (e.g. checking a slice's length before indexing into it).
3. Try `if 1 {}` and read the compiler's exact error, confirming there's no implicit int-to-bool
   conversion.

## ⚠️ Common Mistakes

- Relying on evaluation order with `||` the way you rely on it with `&&` but getting the logic
  backwards — remember `||` short-circuits when the **left** side is already `true`, not `false`.
- Writing a condition that assumes some other language's truthy/falsy rules — Go has none; every
  condition must be an explicit `bool` expression.
