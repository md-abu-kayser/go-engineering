# 12 — Unsigned Integers

## 🎯 Learning Objectives

- Use `uint8`/`uint16`/`uint32`/`uint64`/`uint`, and know their non-negative-only range.
- Understand exactly why `len()`/`cap()` deliberately return a **signed** `int`, not `uint`.
- Recognize and avoid the classic unsigned-underflow bug.

## 📖 Concept

Unsigned integers store **only non-negative values**, using every bit for magnitude instead of
reserving one for sign — which doubles the positive range compared to the same-sized signed type:

```go
int8:  -128 to 127
uint8:    0 to 255      // same 8 bits, but 0..255 instead of -128..127
```

### Why Go's `len()`/`cap()` deliberately return `int`, not `uint`

This is a genuinely deliberate, debated Go design decision worth understanding: even though a
length or capacity can never be negative, `len()` and `cap()` return a **signed** `int`. The
reasoning: signed arithmetic on lengths (subtracting, comparing against `-1` as a "not found"
sentinel, etc.) is common and safe with `int`, but becomes a minefield of underflow bugs with
`uint` — exactly the problem demonstrated below.

### The classic bug: unsigned underflow

```go
var count uint = 0
count--
// count is now NOT -1 — it WRAPS to uint's maximum value (a huge positive number)
```

Because `uint` cannot represent `-1`, decrementing past zero **wraps around** to the type's
maximum value instead — silently, with no error. This is arguably the single most common
integer-related bug across languages with unsigned types, and it's exactly why a loop like:

```go
for i := someUint; i >= 0; i-- {  // BUG: i >= 0 is ALWAYS true for an unsigned type!
    ...
}
```

either loops forever (conceptually) or, in Go's case, keeps looping until it wraps all the way
back around — either way, not what was intended. `i >= 0` is a **useless**, always-true condition
for any unsigned type, since it can never be negative in the first place.

## 🔍 Code Walkthrough (`main.go`)

```go
var count uint = 0
count--
fmt.Printf("uint(0) - 1  = %d (NOT -1 ...)\n", count)
```

This is the underflow bug made directly visible: `count` after this decrement is an enormous
positive number (`uint`'s maximum value), not `-1` — proof that this isn't a theoretical concern,
it's exactly what happens the moment an unsigned value would need to go negative.

```go
for i := 0; i < 5; i++ { // bounded manually to keep this demo finite
```

Note the outer loop here uses a **signed** `int` counter, specifically so this demonstration
itself terminates safely — the `break` inside simulates what a naive, buggy `uint`-based
"countdown" loop condition would have failed to do on its own.

## ▶️ How to Run

```bash
cd level-01-fundamentals/12-unsigned-integers
go run main.go
```

## ✅ Expected Output

```
=== Unsigned Integers ===
----------------------------------
uint8        : 200 (range: 0 to 255)
int8 max     : 127  <- uint8 goes almost twice as high, using the same 8 bits

len(s)       : 3 (type: int — a signed int, on purpose)

--- Unsigned underflow ---
uint(0) - 1  = 18446744073709551615 (NOT -1 — uint cannot represent negative numbers)

--- Why this matters in a loop ---
  itemsLeft = 3
  itemsLeft = 2
  itemsLeft = 1
  itemsLeft = 0
  (stopping the DEMO here — a naive `for itemsLeft >= 0` loop would NOT have stopped)
```

(The exact underflow value depends on `uint`'s size on your platform — `18446744073709551615` is
`uint64`'s maximum, i.e. `2^64 - 1`, on a typical modern 64-bit machine.)

## 🧠 Key Takeaways

- Unsigned types trade "can be negative" for double the positive range at the same bit width.
- `len()`/`cap()` deliberately return signed `int`, specifically to avoid underflow footguns.
- Decrementing an unsigned value below zero wraps to that type's **maximum** value, silently.
- `for i := someUint; i >= 0; i--` is a genuinely useless, always-true condition — a classic bug.

## 🛠️ Try It Yourself

1. Confirm `math.MaxUint64` (or whatever your platform's `uint` maps to) matches the underflow
   value this program actually prints.
2. Write a **correct** countdown loop over a `uint8` that stops safely at zero (hint: check
   `if i == 0 { break }` at the top of the loop body, before decrementing, rather than relying on
   a `>= 0` continuation condition).
3. Try subtracting a larger unsigned value from a smaller one (e.g. `uint(3) - uint(5)`) and
   predict the wrapped result before running it.

## ⚠️ Common Mistakes

- Writing `for i := x; i >= 0; i--` where `i` is any unsigned type — this condition can never be
  false, so the loop either never terminates as intended or wraps around and continues.
- Choosing `uint` "because a count can't be negative" without considering that **intermediate**
  calculations (subtracting two counts, for instance) can absolutely go negative conceptually,
  even if the final, correct answer wouldn't — `int` is usually the safer default even for
  values that happen to always be non-negative in practice.
