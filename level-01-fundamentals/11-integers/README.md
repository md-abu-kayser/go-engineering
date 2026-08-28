# 11 — Integers

## 🎯 Learning Objectives

- Know that plain `int`'s size is platform-dependent (32 or 64 bits).
- Use the explicitly-sized integer types (`int8`, `int16`, `int32`, `int64`) when size matters.
- Understand integer division/modulo truncation, and silent overflow wraparound.

## 📖 Concept

### `int`: the everyday, platform-dependent type

```go
var i int = 42
```

Plain `int` is what you'll use for the vast majority of integer values in Go. Its size (32-bit or
64-bit) is determined by the target platform — in practice, essentially every modern platform Go
targets uses 64-bit `int`. You should default to `int` unless you have a **specific** reason to
need an exact size.

### The explicitly-sized types

```go
int8   // -128 to 127
int16  // -32,768 to 32,767
int32  // -2,147,483,648 to 2,147,483,647
int64  // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
```

Reach for these specifically when the exact size genuinely matters: matching a binary file
format, a network protocol's wire format, or interoperating with C code — not as a default
"more precise" choice for ordinary application logic.

### Integer division truncates toward zero

```go
7 / 2   // 3 (not 3.5 — integer division discards the fractional part entirely)
-7 / 2  // -3 (truncates TOWARD ZERO, not toward negative infinity — some languages differ here)
7 % 2   // 1  (the remainder)
```

This "truncate toward zero" rule for negative operands is worth memorizing specifically, since
some other languages (Python, for instance) truncate toward negative infinity instead, giving a
different result for negative operands.

### Overflow wraps around silently

```go
var small int8 = 127  // int8's maximum value
small++                 // wraps to int8's MINIMUM value: -128
```

Go does **not** panic, error, or warn on integer overflow by default — the value simply wraps
around using standard two's-complement arithmetic. This is a real, common source of subtle bugs
in code that doesn't carefully consider its integer types' ranges.

## 🔍 Code Walkthrough (`main.go`)

```go
var small int8 = 127
small++
```

`127` is `int8`'s maximum representable value. Incrementing it doesn't produce an error or a
larger type automatically — it silently becomes `-128`, `int8`'s minimum value, demonstrating
overflow wraparound directly rather than just describing it.

## ▶️ How to Run

```bash
cd level-01-fundamentals/11-integers
go run main.go
```

## ✅ Expected Output

```
=== Integers ===
----------------------------------
int          : 42 (size on this machine: 64 bits)
int8         : 127 (range: -128 to 127)
int16        : 32000 (range: -32768 to 32767)
int32        : 2000000000 (range: -2147483648 to 2147483647)
int64        : 9000000000000000000 (range: -9223372036854775808 to 9223372036854775807)

--- Integer division & modulo ---
7 / 2   = 3 (truncated, not rounded)
7 % 2   = 1 (remainder)
-7 / 2  = -3 (truncates TOWARD ZERO, not toward negative infinity)
-7 % 2  = -1

--- Overflow ---
int8(127) + 1 = -128 (wrapped around to the MINIMUM value, silently)
```

## 🧠 Key Takeaways

- Plain `int` is platform-dependent-sized (effectively always 64-bit on modern targets) and
  should be your default choice.
- Explicitly-sized types (`int8`/`int16`/`int32`/`int64`) matter for exact binary formats, not
  everyday logic.
- Integer division truncates toward zero, which affects negative operands differently than some
  other languages.
- Overflow wraps around silently — Go performs no automatic bounds checking on arithmetic.

## 🛠️ Try It Yourself

1. Read `main.go`'s use of `unsafe.Sizeof(i)` and confirm it matches what `go env GOARCH` implies
   for your machine (64-bit architectures like `amd64`/`arm64` give a 64-bit `int`).
2. Try `-7 % 2` by hand before running the program, and check your prediction against Go's actual
   "truncate toward zero" behavior.
3. Deliberately overflow an `int32` (not just `int8`) by adding `1` to `math.MaxInt32`, and
   confirm it wraps to `math.MinInt32`.

## ⚠️ Common Mistakes

- Assuming integer division rounds — it truncates; use explicit rounding logic (or work in
  `float64`) if you actually need rounding.
- Choosing a small integer type (`int8`, `int16`) "to save memory" for a value that could
  plausibly grow beyond its range — silent overflow wraparound is a much worse bug than the
  memory `int` or `int64` would have cost.
