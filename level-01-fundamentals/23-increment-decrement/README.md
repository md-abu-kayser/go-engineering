# 23 — Increment & Decrement

## 🎯 Learning Objectives

- Use `++` and `--` correctly, as standalone statements.
- Understand the genuinely Go-specific rule: `++`/`--` are **statements**, never expressions.
- Know that Go has **no** pre-increment (`++x`) form at all — only post-increment-style syntax,
  and even that isn't usable as a value.

## 📖 Concept

```go
n++   // equivalent to n += 1, equivalent to n = n + 1
n--    // equivalent to n -= 1
```

At a glance, this looks identical to `++`/`--` in C, Java, JavaScript, and many other languages.
The genuinely important difference is **what kind of thing `++`/`--` are** in Go.

### `++`/`--` are STATEMENTS, not expressions — this is Go-specific

In C-family languages, `x++` is an **expression** — it has a value (`x`'s value before or after
incrementing, depending on prefix/postfix form), and you can use that value directly:

```c
// C, Java, JavaScript — this is valid there:
y = x++;          // assign x's old value to y, then increment x
foo(x++);          // pass x's old value to foo, then increment x
```

**None of this compiles in Go.** In Go, `x++` is a **statement** on its own — it performs the
increment and produces **no usable value whatsoever**. You cannot assign it, pass it as an
argument, or use it in any expression context at all:

```go
y := x++          // COMPILE ERROR in Go
fmt.Println(x++)   // COMPILE ERROR in Go
```

### Go has no pre-increment form, either

Some languages distinguish `x++` (post-increment) from `++x` (pre-increment), with different
values when used as expressions. **Go has neither concept as an expression** — `++x` isn't even
valid Go syntax at all (`x++` is the only form that exists, and it's always a statement).

### Why Go made this choice

This is a deliberate simplification: the pre/post-increment distinction is a notorious source of
subtle bugs and "undefined behavior" edge cases in C-family languages (e.g. combining multiple
increments of the same variable within one expression). Go sidesteps the entire category of bug
by simply not allowing `++`/`--` to participate in expressions at all — if you want the
pre-increment value, write it out: `x + 1`, without mutating `x` at all, or increment first as
its own statement and then use `x` normally on the next line.

## 🔍 Code Walkthrough (`main.go`)

```go
n--
n-- // two separate decrements
```

Two separate statements, on two separate lines (each ending its own statement) — there's no way
to write "decrement twice" in a single expression the way `n -= 2` already does more directly;
`--` only ever means "decrement by exactly one, as its own statement."

## ▶️ How to Run

```bash
cd level-01-fundamentals/23-increment-decrement
go run main.go
```

## ✅ Expected Output

```
=== Increment & Decrement ===
----------------------------------
n := 5    -> n = 5
n++       -> n = 6
n--; n--  -> n = 4

--- Counting up with ++ ---
  count = 0
  count = 1
  count = 2

--- What ++ and -- CANNOT do (see README) ---
The following would NOT compile if uncommented:
  x := 5
  y := x++       // ERROR: ++ is a STATEMENT, not an expression
  fmt.Println(x++) // ERROR: same reason — you can't use its "result"
  z := ++x        // ERROR: Go has NO pre-increment form at all
```

## 🧠 Key Takeaways

- `x++`/`x--` are Go **statements**, never expressions — they cannot be assigned, passed as
  arguments, or used inside any larger expression.
- Go has no pre-increment (`++x`) syntax at all — only the statement form `x++` exists.
- This is a deliberate simplification versus C-family languages, removing an entire class of
  order-of-evaluation bugs.
- Use `x += 1` (or plain `x = x + 1`) if you need the increment as part of a larger expression's
  logic — `++` itself simply can't participate.

## 🛠️ Try It Yourself

1. Uncomment (in a scratch copy) each of the three "would NOT compile" lines from `main.go`, one
   at a time, and read Go's exact compiler error for each.
2. Rewrite the counting loop to count **down** from 3 to 0 using `--` instead.
3. Compare Go's rule here to a language you already know that has C-style pre/post increment —
   write down one concrete example where that distinction has caused (or could cause) a subtle
   bug, and confirm Go's restriction genuinely prevents it.

## ⚠️ Common Mistakes

- Writing `y := x++` out of habit from another language and being confused by Go's compiler error
  — remember `++`/`--` are statement-only in Go, full stop.
- Looking for a `++x` pre-increment form in Go — it doesn't exist in any form, not even as a
  statement; only `x++` is valid syntax.
