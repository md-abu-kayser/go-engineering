# 13 — Floating-Point

## 🎯 Learning Objectives

- Use `float64` (the default) and `float32`, and know when each is appropriate.
- Understand why `0.1 + 0.2 != 0.3` in binary floating point — and a Go-specific subtlety in
  *when* that surprise actually shows up.
- Recognize `+Inf`, `-Inf`, and `NaN`, and know `NaN` is never equal to anything — even itself.

## 📖 Concept

Go implements IEEE 754 binary floating-point arithmetic — the same standard virtually every
mainstream language uses, with the same well-known surprises. But Go has one extra wrinkle worth
understanding specifically, covered below.

### `float64` vs `float32`

```go
var f64 float64 = 3.14159265358979 // Go's default float type — use this unless you have a reason not to
var f32 float32 = 3.14159265358979  // half the storage, noticeably less precision
```

A bare floating-point literal (`3.14`) defaults to `float64` unless context requires otherwise —
`float32` exists mainly for memory-constrained scenarios (huge arrays of floats) or matching an
external format that specifically requires 32-bit floats.

### The classic surprise — and a Go-specific subtlety

Most decimal fractions — including `0.1`, `0.2`, and `0.3` — **cannot** be represented exactly in
binary floating point; each is stored as the *closest representable* binary approximation, and
those tiny errors don't necessarily cancel out when added. In most languages, `0.1 + 0.2 == 0.3`
is simply `false`. **In Go, it depends on exactly how you write it:**

```go
constSum := 0.1 + 0.2   // 0.1 and 0.2 here are UNTYPED CONSTANTS
constSum == 0.3          // -> often TRUE in Go!
```

This is genuinely Go-specific: as [lesson 05](../05-typed-constants) covered, untyped constant
expressions are evaluated using **arbitrary precision**, not `float64` rounding, until the moment
they're actually assigned to a typed value. So `0.1 + 0.2` here is computed *exactly*, in
unlimited precision, and only rounded to the nearest `float64` **once**, at the very end — which
can happen to land on precisely the same `float64` value that `0.3` alone would round to.

To see the **real**, reliable version of the classic surprise, the values need to already be
`float64` **variables** at the moment of addition, so each is independently rounded **before**
being summed:

```go
var a float64 = 0.1
var b float64 = 0.2
var c float64 = 0.3
a + b == c   // -> false — THIS is the surprise everyone expects
```

Here, `a` and `b` are each rounded to their nearest `float64` representation *individually*, and
adding those two already-rounded values accumulates a small error that `0.3` alone never picked
up — so the comparison correctly (if confusingly) comes out `false`.

### Comparing floats correctly: use an epsilon

```go
const epsilon = 1e-9
closeEnough := math.Abs(a-b) < epsilon
```

Regardless of which of the two cases above you're in, the robust, general rule is the same:
**never** compare floating-point results with `==` once any arithmetic is involved — check that
the difference is smaller than some acceptably tiny threshold instead.

### Special values: `Inf` and `NaN`

Unlike integer division by zero ([lesson 11](../11-integers), which panics), **floating-point**
division by zero produces special IEEE 754 values instead — but only when computed at **runtime**
through variables; as literal constants, `1.0 / 0.0` is actually a **compile-time error** in Go
(constant division by zero is caught before your program ever runs):

```go
var zero float64 = 0.0
1.0 / zero    // +Inf
-1.0 / zero   // -Inf
math.NaN()     // "Not a Number" — e.g. the result of 0.0/0.0
```

`NaN` has one famously strange property: it is **never equal to anything**, including itself —
`NaN == NaN` is `false`. Use `math.IsNaN(x)` to actually check for it.

## 🔍 Code Walkthrough (`main.go`)

```go
constSum := 0.1 + 0.2
fmt.Printf("  == 0.3 ? %t\n", constSum == 0.3)
```

This prints `true` — which looks like the "surprise" never happened. It did; it's just hidden by
Go's arbitrary-precision constant folding, exactly as the concept section explains.

```go
var a float64 = 0.1
var b float64 = 0.2
runtimeSum := a + b
```

Forcing `a` and `b` to be genuine `float64` **variables** before adding them reproduces the
surprise everyone expects — each is rounded independently first, so their sum picks up an error
`0.3` alone doesn't have.

## ▶️ How to Run

```bash
cd level-01-fundamentals/13-floating-point
go run main.go
```

## ✅ Expected Output

```
=== Floating-Point ===
----------------------------------
float64 : 3.14159265358979 (full precision)
float32 : 3.14159274101257 (visibly LESS precise — fewer bits to work with)

--- The classic surprise (and a Go-specific subtlety) ---
0.1 + 0.2 (const-folded)   = 0.29999999999999998890
  == 0.3 ?                   true (misleading! see below for why)
a + b (real float64 runtime arithmetic) = 0.30000000000000004441
  == c (float64 0.3) ?                     false (the REAL, reliable surprise)
  within epsilon of 0.3 ?                   true (this is the reliable check)

--- Special values ---
1.0 / 0.0  = +Inf
-1.0 / 0.0 = -Inf
NaN        = NaN
NaN == NaN ? false (NaN is NEVER equal to anything, including itself)
math.IsNaN(NaN) ? true (use this instead of == to check for NaN)
```

## 🧠 Key Takeaways

- `float64` is Go's default and usual choice; `float32` trades precision for half the memory.
- Untyped constant arithmetic in Go uses arbitrary precision, rounding only once at the end —
  which can mask the classic floating-point surprise if you test it with bare literals.
- The surprise reliably reappears once real `float64` **variables** are involved in the addition.
- Always compare floats with an epsilon threshold once genuine runtime arithmetic is involved —
  never rely on `==`.
- `1.0 / 0.0` as a literal constant expression is a **compile-time error** in Go; it only produces
  `+Inf` at runtime, through variables.
- `NaN` compares unequal to everything, including itself — use `math.IsNaN` to detect it.

## 🛠️ Try It Yourself

1. Print `constSum` and `runtimeSum` both to 20 decimal places and compare the two — confirm
   they're genuinely different `float64` values, even though both came from "0.1 + 0.2."
2. Try a few other decimal fractions as bare constant literals (e.g. `0.7 + 0.1 == 0.8`) and see
   whether Go's constant folding happens to mask the surprise there too, or not — it depends on
   the specific bit patterns involved.
3. Write your own `floatEquals(a, b, epsilon float64) bool` helper and use it in place of every
   direct float comparison in this program.

## ⚠️ Common Mistakes

- Testing "does Go have the classic float bug?" with bare literals (`0.1 + 0.2 == 0.3`) and
  concluding it doesn't — you may just be seeing constant folding's arbitrary-precision rounding,
  not evidence that runtime `float64` arithmetic is somehow exact.
- Checking `if x == math.NaN()` to detect NaN — this is **always false**, no matter what; use
  `math.IsNaN(x)`.
