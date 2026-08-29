# 18 — Arithmetic Operators

## 🎯 Learning Objectives

- Use all five arithmetic operators: `+` `-` `*` `/` `%`.
- Know that `+` is Go's **only** arithmetic-looking operator that also works on strings.
- Distinguish integer division (truncating) from float division (true) using the same operator.

## 📖 Concept

Go's five arithmetic operators are unsurprising if you've used any C-family language:

```go
a + b   // addition
a - b   // subtraction
a * b   // multiplication
a / b   // division (see below — behavior differs by operand type)
a % b   // remainder (integers only)
-a       // unary negation
```

### `/` behaves differently depending on the operand types

[Lesson 11](../11-integers) already covered this for integers: `/` **truncates** when both
operands are integers. When operands are floats instead, `/` performs genuine, non-truncating
division:

```go
17 / 5     // 3   (integer division — truncated)
17.0 / 5.0  // 3.4 (float division — the real quotient)
```

Same operator, same symbol, genuinely different behavior depending on operand type — worth
remembering explicitly rather than assuming "division" always means one specific thing in Go.

### `+` is the one operator that also works on strings

```go
"Hello, " + "Gopher" + "!"   // string concatenation
```

`+` is Go's **only** arithmetic-symbol operator with a defined meaning for `string` operands —
`-`, `*`, `/`, and `%` are simply invalid on strings; there's no "subtract a substring" or
"multiply a string" operation in Go. This overload exists specifically for concatenation and
nothing else.

### `%` (remainder) is integers-only

```go
17 % 5   // 2 — fine
17.0 % 5.0  // COMPILE ERROR — % doesn't work on floats at all
```

If you need a "remainder" for floats, reach for `math.Mod` instead — `%` itself is defined only
for integer types in Go.

## 🔍 Code Walkthrough (`main.go`)

```go
greeting := "Hello, " + "Gopher" + "!"
```

Placed directly after the numeric examples, specifically to contrast: this is the **same**
symbol (`+`) used moments earlier for numeric addition, now doing something conceptually
different (concatenation) — Go allows this one overload and no others among the arithmetic
operators.

## ▶️ How to Run

```bash
cd level-01-fundamentals/18-arithmetic-operators
go run main.go
```

## ✅ Expected Output

```
=== Arithmetic Operators ===
----------------------------------
17 + 5 = 22
17 - 5 = 12
17 * 5 = 85
17 / 5 = 3 (integer division truncates — lesson 11)
17 % 5 = 2 (remainder)

-17 = -17 (unary minus)

--- + also means string concatenation ---
"Hello, " + "Gopher" + "!" = "Hello, Gopher!"

--- Float arithmetic ---
17.0 / 5.0 = 3.4 (true division, not truncated)
```

## 🧠 Key Takeaways

- Go's five arithmetic operators: `+ - * / %`, plus unary `-` for negation.
- `/` truncates for integer operands, but performs true division for float operands.
- `+` is the only arithmetic-symbol operator that also works on strings (concatenation).
- `%` (remainder) only works on integers — use `math.Mod` for a float remainder.

## 🛠️ Try It Yourself

1. Try `17.0 % 5.0` directly and read the compiler's exact error confirming `%` is integer-only.
2. Use `math.Mod(17.0, 5.0)` instead and confirm it gives you the float remainder you'd expect.
3. Try subtracting one string from another (`"hello" - "h"`) and read the compiler error
   confirming `-` has no meaning for strings at all.

## ⚠️ Common Mistakes

- Assuming `%` works on floats the same way it does on ints — it doesn't; that's a compile error,
  not a runtime surprise.
- Forgetting that `/`'s behavior depends entirely on whether the operands are integer or float
  types — the exact same-looking expression can mean truncating or true division depending on
  what's on either side.
