# 25 — if with Init

## 🎯 Learning Objectives

- Use the `if init; condition { }` form to declare a variable scoped only to the `if`/`else` chain.
- Recognize this as the idiomatic pattern for functions returning `(value, error)` or `(value, ok)`.
- Understand precisely where the init statement's variables are (and aren't) visible.

## 📖 Concept

Go's `if` supports an optional **init statement**, separated from the condition by a semicolon:

```go
if n, err := strconv.Atoi("123"); err == nil {
    // use n here
} else {
    // err (and n) are ALSO visible here
}
```

This is extremely common in idiomatic Go, precisely because so many functions return a
`(value, error)` pair ([lesson 38 of level 00](../../level-00-getting-started/38-reading-error-messages))
— running the call and checking its result in one compact statement avoids a separate,
easy-to-forget-to-check variable declared several lines earlier.

### The scope rule

Variables declared in the init statement are visible for the **entire** `if`/`else` chain — every
branch — but **nowhere else**:

```go
if n, err := strconv.Atoi("123"); err == nil {
    fmt.Println(n)
} else {
    fmt.Println(err) // fine — err is in scope here too
}

fmt.Println(n) // COMPILE ERROR — n is undefined out here
```

This scoping is deliberate: `n` and `err` exist **specifically** to support this one decision —
once the `if`/`else` chain is done, they've served their purpose and shouldn't linger in the
enclosing function's namespace where they could be confused with something else.

## 🔍 Code Walkthrough (`main.go`)

```go
if n, err := strconv.Atoi("not a number"); err == nil {
    fmt.Printf("Parsed successfully: %d\n", n)
} else {
    fmt.Printf("Failed to parse %q: %v\n", "not a number", err)
}
```

This deliberately uses input that **fails** to parse, so you see the `else` branch actually
execute — and notice `err` (the genuinely useful value here) is available in that branch exactly
because it was declared in the shared init statement, not just inside the `if`'s own body.

```go
if length := len("Gopher"); length > 3 {
```

A simpler example with no error involved at all — `if init; condition` isn't exclusively for
error handling; it's useful anytime you want a value scoped to just one decision.

## ▶️ How to Run

```bash
cd level-01-fundamentals/25-if-with-init
go run main.go
```

## ✅ Expected Output

```
=== if with Init ===
----------------------------------
Parsed successfully: 123
Failed to parse "not a number": strconv.Atoi: parsing "not a number": invalid syntax

See the README: `n` and `err` above are NOT visible out here.

"Gopher" has 6 characters — that's more than 3.
```

## 🧠 Key Takeaways

- `if init; condition { }` runs `init` first, then evaluates `condition` — both share one scope.
- This is the idiomatic pattern for checking a `(value, error)` result immediately, in one place.
- Variables from the init statement are visible across the whole `if`/`else` chain, but nowhere
  outside it.
- Use this pattern anytime a value is only needed to make one decision, to avoid unnecessary
  variables lingering in the surrounding function.

## 🛠️ Try It Yourself

1. Add a line right after the first `if`/`else` block attempting to print `n` directly, and read
   the exact "undefined: n" compiler error.
2. Rewrite the `strconv.Atoi` examples **without** the init-statement form (declare `n, err`
   separately, above the `if`) and compare readability — and notice `n`/`err` now linger in the
   function's scope afterward.
3. Write your own `if init; condition` using a map lookup's `value, ok := m[key]` pattern instead
   of an error.

## ⚠️ Common Mistakes

- Trying to use the init statement's variables after the `if`/`else` chain ends — they're
  deliberately out of scope there; declare them normally beforehand if you need that.
- Forgetting the semicolon between the init statement and the condition — `if n, err :=
  strconv.Atoi("123") err == nil` (no `;`) is a syntax error, not a working shortcut.
