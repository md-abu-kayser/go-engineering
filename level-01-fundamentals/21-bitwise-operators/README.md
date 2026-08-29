# 21 — Bitwise Operators

## 🎯 Learning Objectives

- Use `&`, `|`, `^` (binary XOR), and `^` (unary NOT) at the bit level.
- Use Go's own `&^` ("AND NOT" / bit clear) operator — not present as a dedicated symbol in most
  other C-family languages.
- Use `<<`/`>>` to shift bits, and connect that to multiplying/dividing by powers of two.

## 📖 Concept

Bitwise operators work on the individual **bits** of an integer's binary representation, rather
than treating it as a single numeric value.

```go
a & b    // AND: 1 in a result bit only where BOTH a and b have a 1 there
a | b     // OR:  1 in a result bit where EITHER a or b has a 1 there
a ^ b      // XOR: 1 in a result bit where a and b DIFFER (binary, two operands)
^a          // NOT: flips every bit (unary, one operand — note ^ is BOTH XOR and NOT depending on arity)
```

### Go's own operator: `&^` (AND NOT / bit clear)

```go
a &^ b   // "bit clear": take a, but force OFF any bit that's set in b
```

This is a genuinely Go-specific operator — most other C-family languages express "clear these
bits" as `a & ~b` (AND combined with a separate NOT), while Go gives it its **own** dedicated
symbol. `a &^ b` is exactly equivalent to `a & (^b)`, just spelled as one operator instead of two.
This is the natural tool for **removing** one specific flag from a combined set of flags
([lesson 06](../06-iota)'s bit-flag pattern), without needing a separate negation step.

### Shifting: `<<` and `>>`

```go
1 << 4   // 16  — shifting left by n is the same as multiplying by 2^n
16 >> 2   // 4   — shifting right by n is the same as dividing by 2^n (for non-negative values)
```

`<<`/`>>` are also central to the `1 << iota` bit-flag pattern from
[lesson 06](../06-iota) — each successive flag gets its own bit by shifting `1` further left each
time.

## 🔍 Code Walkthrough (`main.go`)

```go
const (
    FlagRead  = 1 << iota
    FlagWrite
    FlagExec
)
perms := FlagRead | FlagWrite | FlagExec
permsNoWrite := perms &^ FlagWrite
```

This directly connects three lessons together: `iota` ([lesson 06](../06-iota)) generates
distinct bit positions, `|` combines them into one value representing "all three flags," and
`&^` cleanly **removes** just one flag (`FlagWrite`) from that combined value — the exact
practical task `&^` exists for.

## ▶️ How to Run

```bash
cd level-01-fundamentals/21-bitwise-operators
go run main.go
```

## ✅ Expected Output

```
=== Bitwise Operators ===
----------------------------------
a = 1100 (12), b = 1010 (10)
a & b  = 1000 (8)  (AND: 1 where BOTH bits are 1)
a | b  = 1110 (14)  (OR: 1 where EITHER bit is 1)
a ^ b  = 0110 (6)  (XOR: 1 where the bits DIFFER)
^a     = -13  (unary NOT: flips every bit — sign matters for signed types)
a &^ b = 0100 (4)  (BIT CLEAR: a, with b's set bits forced OFF)

--- Shifting ---
1 << 4 = 16 (same as 1 * 2^4)
16 >> 2 = 4 (same as 16 / 2^2)

--- Practical use: removing one flag from a set ---
perms          = 111 (all three flags)
perms &^ Write = 101 (write flag cleanly removed)
```

## 🧠 Key Takeaways

- `&`/`|`/`^` (binary) work bit-by-bit, following AND/OR/XOR truth tables.
- `^` as a **unary** operator means bitwise NOT — the same symbol as XOR, disambiguated by arity.
- `&^` ("bit clear") is Go-specific — `a &^ b` clears exactly the bits set in `b` from `a`.
- `<<`/`>>` shift bits, equivalent to multiplying/dividing by powers of two.
- `&^` is the natural tool for removing one flag from a combined bit-flag value.

## 🛠️ Try It Yourself

1. Compute `a &^ b` by hand from their binary representations, and confirm it matches the
   program's output.
2. Remove a *different* flag (`FlagRead`, not `FlagWrite`) from `perms` using `&^`, and confirm
   the remaining two flags are still both set.
3. Confirm `a &^ b` really does equal `a & (^b)` by computing both expressions and comparing.

## ⚠️ Common Mistakes

- Writing `a & ^b` (two separate operators) instead of Go's single `a &^ b` — both may even
  produce the same result, but `&^` is the single, idiomatic, Go-specific way to express "bit
  clear" directly.
- Forgetting `^` means something different depending on whether it has one operand (NOT) or two
  (XOR) — reading unfamiliar code with `^` requires checking which form is actually being used.
