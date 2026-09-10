# 52 — Shadowing

## 🎯 Learning Objectives

- Go beyond [lesson 02](../02-short-declarations)'s introduction to shadowing with a realistic,
  genuinely dangerous example.
- Recognize the classic "shadowed `err` in a nested if-block" bug pattern.
- Know that even Go's built-ins (`len`, `true`, `int`, ...) are ordinary, shadowable identifiers —
  not reserved keywords.

## 📖 Concept

[Lesson 02](../02-short-declarations) introduced shadowing briefly: a `:=` inside a nested block,
using a name that already exists in an outer scope, creates a genuinely **separate** variable.
This lesson focuses on **why that matters** in real code.

### The classic, genuinely dangerous bug: shadowed `err`

```go
func process(input string) error {
    value, err := strconv.Atoi(input)
    if err != nil {
        return err
    }

    if value < 0 {
        doubled, err := computeDoubled(value) // BUG: := creates a NEW err here
        if err != nil {
            // this branch DOES see the real error...
        }
    }

    return err // ...but THIS returns the OUTER err — still nil, even if the inner one wasn't!
}
```

Because `doubled, err := ...` uses `:=` inside the `if value < 0` block, and **both** `doubled`
and (crucially) `err` are new to *that specific block's scope*, Go creates a brand-new `err`
local to the `if` block — completely separate from the `err` declared several lines above. Any
error from `computeDoubled` is checked correctly *inside* that block, but is then silently lost:
the function's final `return err` refers to the **outer** `err`, which was never touched by the
inner assignment.

**The fix** is to use plain `=` instead of `:=` when you mean to update an existing variable in an
outer scope, not declare a new one:

```go
var doubled int
doubled, err = computeDoubled(value) // = , not := — reuses the OUTER err
```

This bug is extremely common in real Go code specifically because `:=` is so convenient and
idiomatic — it's easy to reach for it out of habit even when a variable of that name already
exists one scope up.

### Even built-in identifiers can be shadowed

```go
len := "surprise!"
fmt.Println(len) // prints "surprise!" — len is now a plain string variable in this scope
```

Names like `len`, `cap`, `true`, `false`, `nil`, `int`, `string`, and others are **not** reserved
keywords in Go — they're ordinary, pre-declared identifiers living in what's called the
**universe scope** (the outermost scope, surrounding every package). Because they're ordinary
identifiers, they can be shadowed by a more local declaration exactly like any other name — Go's
compiler won't stop you, even though doing this deliberately is almost always a bad idea in real
code.

## 🔍 Code Walkthrough (`main.go`)

```go
if value < 0 {
    doubled, err := computeDoubled(value)
    if err != nil {
        fmt.Printf("  (bug demonstrated: inner err was %v, but caller never sees it)\n", err)
    }
    _ = doubled
}
return err
```

The `fmt.Printf` inside the `if` block **does** print the real, non-nil inner error — proving the
bug isn't that the error goes undetected entirely, just that it never makes it back out through
the function's actual return value, which is the part that matters to `process`'s caller.

## ▶️ How to Run

```bash
cd level-01-fundamentals/52-shadowing
go run main.go
```

## ✅ Expected Output

```
=== Shadowing ===
----------------------------------
--- The classic err-shadowing bug ---
  (bug demonstrated: inner err was value -2000 is too extreme to double safely, but caller never sees it)
process("-2000") returned err = <nil> (BUG: should have been non-nil!)

--- Even built-ins can be shadowed (rarely a good idea!) ---
  len is now a variable holding: "surprise!"
  outside the block, len is the built-in function again: 5
```

## 🧠 Key Takeaways

- Using `:=` inside a nested block with a name that already exists outward creates a **new**,
  shadowed variable — it does not update the outer one.
- This is a genuinely common, real-world source of silently-swallowed errors — always check
  whether a variable you're assigning to with `:=` already exists in an outer scope.
- The fix is to declare the outer variable once and use plain `=` for subsequent updates within
  nested blocks.
- Go's built-ins (`len`, `true`, `int`, etc.) are ordinary, shadowable identifiers, not reserved
  keywords — though shadowing them deliberately is rarely a good idea.

## 🛠️ Try It Yourself

1. Fix `process`'s bug using the `var` + `=` approach shown above, and confirm
   `process("-2000")` now correctly returns a non-nil error.
2. Try shadowing `true` inside a block (`true := false`) and see what happens if you then try to
   use `true` as a boolean literal inside that same block.
3. Many editors and linters (including `go vet` with certain settings, and tools like
   `shadow` from `golang.org/x/tools`) can detect shadowed variables automatically — look up
   the `shadow` analyzer and consider whether it'd be worth enabling for a real project.

## ⚠️ Common Mistakes

- Using `:=` inside a nested `if`/`for`/block when you actually meant to update an outer
  variable — this is precisely the bug this lesson demonstrates, and it compiles cleanly with no
  warning by default.
- Shadowing a built-in identifier by accident (naming a local variable `len` or `min` without
  realizing it) and then being confused why calling it as a function no longer works later in that
  same scope.
