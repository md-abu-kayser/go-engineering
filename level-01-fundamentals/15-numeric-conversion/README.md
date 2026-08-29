# 15 — Numeric Conversion

## 🎯 Learning Objectives

- Convert explicitly between numeric types — and know that Go has no implicit conversion, ever.
- Understand exactly how a `float` → `int` conversion truncates.
- Understand exactly how narrowing to a smaller integer type can silently lose data.

## 📖 Concept

Go has **no implicit numeric conversions** — not even between an `int` and an `int64`, or an
`int32` and an `int`. Every conversion between distinct numeric types must be written explicitly:

```go
var i int = 42
var i64 int64 = int64(i)      // required, even though this specific conversion is always safe
var f float64 = float64(i)
```

This is a deliberate Go design choice: implicit numeric conversions are a common, subtle source
of bugs in languages that allow them (silently promoting an `int` to a `float` mid-expression,
for instance) — Go makes every conversion visible in the source.

### `float` → `int`: truncation, not rounding

```go
var f float64 = 3.99
int(f)   // 3
```

Converting a float to an int **discards** the fractional part entirely — it never rounds to the
nearest whole number. If you need actual rounding, use `math.Round` (on the float) before
converting.

> **A related constant-conversion rule:** as in [lesson 05](../05-typed-constants), an untyped
> float **constant** with a genuine fractional part (like the literal `3.99`) cannot be converted
> to `int` at all — `int(3.99)` written directly is a **compile-time error**, not a truncation.
> The truncating behavior above only applies once the value is a real `float64`, held in a
> variable — which is exactly why this lesson's code assigns `3.99` to a variable first.

### Narrowing to a smaller type: silent data loss

```go
var big int32 = 300
int8(big)  // NOT 300 — int8 can only represent -128..127
```

When you convert a value that doesn't fit into the destination type's range, Go **does not
error, panic, or clamp the value** — it simply discards the high-order bits that don't fit,
producing whatever bit pattern remains, interpreted as the new (smaller) type. The resulting
number can look completely unrelated to the original value.

## 🔍 Code Walkthrough (`main.go` and the test file)

```go
var big int32 = 300
var narrow int8 = int8(big)
```

`300` in binary is `100101100` — nine bits. `int8` only has eight bits of storage, so the highest
bit is discarded, leaving `00101100` (`44`), interpreted as a **signed** 8-bit value. This is
exactly what the test file locks in as `44` — not a guess, but the actual, deterministic result of
truncating bits, verified by `go test`.

```go
func TestNarrowingConversionWraps(t *testing.T) {
    ...
    if got := int8(big); got != 44 {
```

This test exists specifically so that if this lesson's code (or your understanding of it) ever
drifts, `go test` catches the exact discrepancy immediately, rather than requiring you to
re-derive the expected wrapped value by hand every time.

## ▶️ How to Run

```bash
cd level-01-fundamentals/15-numeric-conversion
go run main.go
go test -v ./...
```

## ✅ Expected Output

```
=== Numeric Conversion ===
----------------------------------
int(42) -> int64: 42, -> float64: 42

--- float -> int truncates, never rounds ---
int(3.99)  = 3
int(3.01)  = 3
int(-3.99) = -3 (truncates TOWARD ZERO, not down)

--- Narrowing conversion can silently lose data ---
int32(300) -> int8 = 44 (NOT 300 — the high bits were simply discarded)
int64(40000) -> int16 = -25536 (wrapped, same silent truncation)
```

## 🧠 Key Takeaways

- Go requires explicit conversion between any two distinct numeric types, with no exceptions.
- `float` → `int` conversion truncates toward zero; it never rounds.
- Converting to a smaller type that can't hold the value silently discards high-order bits — no
  error, no panic, no clamping.
- The resulting "wrapped" value is fully deterministic (based on real bit truncation), not random
  garbage — which is exactly why it's testable, as this lesson's test file demonstrates.

## 🛠️ Try It Yourself

1. Work out `300` in binary by hand, confirm it's 9 bits, and verify why discarding the top bit
   gives you `44`.
2. Use `math.Round` (from the `math` package) before converting a float to int, and confirm you
   now get genuine rounding instead of truncation.
3. Run `go test -v ./...` and then deliberately change the test's expected value to something
   wrong, to see what a real test failure looks like for this kind of check.

## ⚠️ Common Mistakes

- Assuming a narrowing numeric conversion will either error or clamp to the type's max/min value —
  it does neither; it silently truncates bits, which can produce a value that looks nothing like
  the original.
- Using `int(f)` when you actually wanted rounding — `int()` always truncates; reach for
  `math.Round(f)` first if you need the nearest whole number instead of the one closer to zero.
