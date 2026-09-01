# 29 — Type Switches (Preview)

## 🎯 Learning Objectives

- Recognize the type-switch syntax: `switch v := v.(type) { case SomeType: ... }`.
- Understand that each `case` in a type switch matches a **type**, not a value.
- Know this lesson is a **preview** — full interface concepts arrive in a later level.

## 📖 Concept

> **A note on scope:** this lesson is intentionally a *preview*. Interfaces — what `any` really
> is, how types satisfy them, and the full type system behind this — are a substantial topic of
> their own, covered properly in a later level. Here, the goal is just to recognize the **syntax**
> when you see it, using Go's built-in `any` type as a stand-in for "a value of any kind."

### The syntax

```go
switch v := v.(type) {
case int:
    // v has type int HERE, inside this case
case string:
    // v has type string HERE, inside this case
default:
    // v keeps its original type here
}
```

`v.(type)` is special syntax that can **only** appear directly inside a `switch` written exactly
this way — it's not valid anywhere else. Each `case` names a **concrete type**, and Go runs the
body of whichever case matches the value's actual, dynamic type at runtime.

### `any`: Go's built-in "value of any type"

```go
func describe(v any) string {
```

`any` is a built-in alias for the empty interface, `interface{}` — a type that every other Go
type automatically satisfies, meaning a parameter of type `any` can accept a value of **any**
type at all. This lesson uses it purely so `describe` has something to type-switch on; a full
explanation of interfaces (what makes `any` special, how custom interfaces work) is intentionally
saved for later.

### Matching `nil` specifically

```go
case nil:
    return "a nil value"
```

A type switch can include a `case nil:` to specifically handle the case where the value itself
is `nil` — this is a distinct case from any concrete type.

## 🔍 Code Walkthrough (`main.go`)

```go
values := []any{42, "Gopher", true, 3.14, nil}
for _, v := range values {
    fmt.Println(describe(v))
}
```

This slice deliberately mixes five completely different types — `int`, `string`, `bool`,
`float64`, and `nil` — specifically to exercise every branch of `describe`'s type switch,
including the `default` case (for `float64`, which has no explicit `case` of its own here).

## ▶️ How to Run

```bash
cd level-01-fundamentals/29-type-switches-preview
go run main.go
```

## ✅ Expected Output

```
=== Type Switches (Preview) ===
----------------------------------
an int: 42 (doubled: 84)
a string: "Gopher" (length: 6)
a bool: true
something else entirely: 3.14 (float64)
a nil value
```

## 🧠 Key Takeaways

- `switch v := v.(type) { case SomeType: ... }` is the type-switch syntax — `v.(type)` only works
  directly inside a switch written this way.
- Each `case` matches a concrete **type**, and `v` takes on that specific type inside its branch.
- `any` (an alias for `interface{}`) means "a value of any type" — the full story of interfaces
  is deliberately deferred to a later level; this lesson is scoped to syntax recognition only.
- `case nil:` specifically handles a `nil` value, separate from any concrete type case.

## 🛠️ Try It Yourself

1. Add a `case float64:` to `describe` and confirm `3.14` now hits that specific branch instead
   of falling through to `default`.
2. Add a new value to the `values` slice of a type not yet handled (e.g. a `[]int{1,2,3}`) and
   confirm it correctly falls into `default`.
3. Inside the `case int:` branch, try using `v` as a `string` — confirm the compiler correctly
   knows `v` is an `int` specifically inside that branch, and rejects the mismatch.

## ⚠️ Common Mistakes

- Trying to use `x.(type)` outside of this exact `switch` form — it's a syntax error anywhere
  else; regular type assertions (a different, related feature) use different syntax entirely.
- Assuming `any` means "no type safety at all" — inside each `case`, `v` is genuinely,
  statically typed as that specific case's type; the type switch itself is what recovers that
  safety after accepting a value of unknown type.
