# 24 — if Statements

## 🎯 Learning Objectives

- Write a basic `if` statement with Go's exact syntax: no parentheses, mandatory braces.
- Understand why Go requires braces even for a single-statement body.
- Confirm the condition must be a genuine `bool` expression, matching [lesson 07](../07-booleans).

## 📖 Concept

```go
if age >= 18 {
    fmt.Println("You are an adult.")
}
```

Two syntax details that differ from many C-family languages:

### No parentheses around the condition

```go
if age >= 18 {     // idiomatic
if (age >= 18) {    // ALSO legal, but never written this way in real Go code — gofmt won't
                      // remove the parens for you, but every style guide and reviewer will
                      // flag them as unnecessary
```

Go simply doesn't need the parentheses to know where the condition ends — the opening `{` marks
that boundary unambiguously.

### Braces are mandatory, always

```go
if age < 0 {
    fmt.Println("...")
}
```

Unlike C, Java, or JavaScript — where you can write `if (cond) doSomething();` with no braces for
a single statement — Go's grammar **requires** braces on every `if` body, no matter how short.
This is a deliberate design choice: brace-optional single-statement bodies are a well-known,
recurring source of real bugs (most famously, Apple's "goto fail" security bug came from exactly
this pattern in C). Go simply removes the option entirely.

### The condition must be a real `bool`

As [lesson 07](../07-booleans) established, Go has no implicit truthy/falsy conversion:

```go
if isMember { ... }               // fine — isMember is already a bool
if age { ... }                     // COMPILE ERROR — age is an int, not a bool
if age >= 65 || isMember { ... }    // fine — this whole expression evaluates to a bool
```

## 🔍 Code Walkthrough (`main.go`)

```go
hasDiscount := age >= 65 || isMember
if hasDiscount {
```

Computing the condition into a named variable first (`hasDiscount`), rather than inlining the
whole expression into the `if`, is a small but genuinely useful habit once a condition has more
than one or two parts — it gives the logic a name, and makes the `if` itself trivially readable.

## ▶️ How to Run

```bash
cd level-01-fundamentals/24-if-statements
go run main.go
```

## ✅ Expected Output

```
=== if Statements ===
----------------------------------
You are an adult.
Eligible for a discount.

See the README for exactly why Go REQUIRES braces, even for one-liners.
```

## 🧠 Key Takeaways

- `if` needs no parentheses around its condition — idiomatic Go never adds them.
- Braces are mandatory on every `if` body, even a single statement — no optional-brace footgun.
- The condition must be a genuine `bool` expression; there's no truthy/falsy fallback.
- Naming a complex condition (`hasDiscount := ...`) before the `if` often reads better than
  inlining it.

## 🛠️ Try It Yourself

1. Try writing an `if` body without braces (`if age >= 18 fmt.Println("...")`) and read the exact
   syntax error Go gives you.
2. Try `if age { }` (using the `int` directly as if it were a `bool`) and read that compiler
   error too.
3. Rewrite the `hasDiscount` example inlined directly into the `if` condition, and decide for
   yourself which version you find more readable.

## ⚠️ Common Mistakes

- Habitually adding parentheses around `if` conditions out of muscle memory from another language
   — not a compile error, but immediately flagged in any Go code review as unidiomatic.
- Expecting to omit braces for a "trivial" one-line `if` body — Go's grammar doesn't allow it,
   full stop.
