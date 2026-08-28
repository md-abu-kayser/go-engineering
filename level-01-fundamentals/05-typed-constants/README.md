# 05 — Typed Constants

## 🎯 Learning Objectives

- Distinguish an untyped constant from a typed one.
- Understand why the same untyped constant can be used as `int`, `float64`, or `int64` with no
  conversion, while a typed constant cannot.
- Know when an untyped floating-point constant can still be used as an integer.

## 📖 Concept

[Lesson 04](../04-constants) introduced `const` broadly. This lesson focuses on one of Go's
genuinely distinctive features: **untyped constants**.

### Untyped: flexible until context demands a type

```go
const untypedNumber = 100   // no type given — this is UNTYPED
```

An untyped constant doesn't commit to a concrete type at its declaration. Instead, it takes on
whatever type is needed **at each place it's used**:

```go
var asInt int = untypedNumber        // used as int here
var asFloat float64 = untypedNumber  // used as float64 here — SAME constant, no conversion
```

No `int(...)` or `float64(...)` conversion is needed in either case — the untyped constant simply
becomes whichever type its context requires, checked at compile time for validity.

### Typed: pinned, like any other typed value

```go
const typedNumber float64 = 100   // explicitly typed — this IS a float64, permanently
```

Once a constant has an explicit type, it behaves exactly like any other value of that type —
assigning it somewhere expecting a different type requires an explicit conversion, the same as
you'd need for a `var` of that type.

```go
var x int = typedNumber        // COMPILE ERROR — typedNumber is float64, not int
var x int = int(typedNumber)   // fine — explicit conversion
```

### The one nuance: whole-number untyped floats can become ints

```go
const wholeFloat = 4.0
var x int = wholeFloat  // fine! 4.0 has no fractional part
```

An untyped **floating-point** constant can still be assigned to an integer variable, but **only**
if its value has no fractional part — `4.0` qualifies, `4.5` would not (that would be a compile
error, not silent truncation).

## 🔍 Code Walkthrough (`main.go`)

```go
var asInt int = untypedNumber
var asFloat float64 = untypedNumber
var asInt64 int64 = untypedNumber
```

Three completely different concrete types, all assigned directly from the **same** untyped
constant, with zero conversion syntax anywhere — this is the practical payoff of untyped
constants, and it's exactly why numeric literals throughout Go (`5`, `3.14`, etc.) work so
smoothly across different numeric contexts without littering your code in explicit casts.

## ▶️ How to Run

```bash
cd level-01-fundamentals/05-typed-constants
go run main.go
```

## ✅ Expected Output

```
=== Typed Constants ===
----------------------------------
untypedNumber as int     : 100 (int)
untypedNumber as float64 : 100 (float64)
untypedNumber as int64   : 100 (int64)
typedNumber as float64   : 100 (float64)
typedNumber as int       : 100 (int, needed int(...))
wholeFloat (4.0) as int  : 4
```

## 🧠 Key Takeaways

- An untyped constant adapts to whatever numeric type its context requires, with no conversion.
- A typed constant behaves exactly like a regular value of that type — conversions are required
  to use it as anything else.
- An untyped floating constant can be used as an integer only if it has no fractional part.
- This is exactly why `5` works seamlessly as `int`, `float64`, or `int64` throughout ordinary Go
  code — numeric literals are untyped constants by default.

## 🛠️ Try It Yourself

1. Change `wholeFloat` to `4.5` and observe the compile error when assigning it to an `int`.
2. Try assigning `typedNumber` directly to an `int` variable **without** `int(...)` and read the
   exact compiler error.
3. Declare your own untyped constant and use it as three different numeric types in three
   different variable declarations, confirming none of them need an explicit conversion.

## ⚠️ Common Mistakes

- Assuming all constants behave like untyped ones — adding an explicit type
  (`const x float64 = 5`) opts back into ordinary, strict typing rules.
- Being surprised that `const wholeFloat = 4.0` can become an `int` while an actual `var f
  float64 = 4.0` cannot be assigned to an `int` without an explicit `int(f)` conversion — the
  "no fractional part" leniency applies specifically to untyped **constants**, not to typed
  `float64` variables, even when their runtime value happens to be a whole number.
