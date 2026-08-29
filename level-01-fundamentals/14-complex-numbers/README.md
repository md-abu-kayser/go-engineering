# 14 — Complex Numbers

## 🎯 Learning Objectives

- Create complex number values using literals and the `complex()` built-in.
- Extract the real and imaginary parts with `real()`/`imag()`.
- Perform arithmetic on complex numbers, and use `math/cmplx` for complex-aware math functions.

## 📖 Concept

Go has **built-in** support for complex numbers — a genuinely unusual feature among mainstream,
systems-adjacent languages, included specifically because Go was designed with scientific and
numerical computing as one of its target use cases from the start.

### Creating complex values

```go
c1 := 3 + 4i              // a complex LITERAL — `i` marks the imaginary part
c2 := complex(1.5, -2.5)  // built explicitly from a real part and an imaginary part
```

`complex128` (built from two `float64`s) is the default complex type, just as `float64` is the
default real-number type; `complex64` (two `float32`s) exists for the same "specific memory
constraint" reasons `float32` does.

### Extracting the parts

```go
real(c1)  // the real component, as a plain float64
imag(c1)   // the imaginary component, as a plain float64
```

`real` and `imag` are **built-in functions** (like `len` or `make`), not methods — they work on
any complex value directly.

### Arithmetic just works

```go
c1 + c2
c1 * c2
```

Go's ordinary `+`, `-`, `*`, `/` operators work directly on complex numbers, following the
standard mathematical rules for complex arithmetic (e.g. multiplication combines both real and
imaginary parts according to `i² = -1`) — no special method calls needed.

### `math/cmplx`: complex-aware math functions

```go
cmplx.Abs(c)  // the MAGNITUDE of c — its distance from the origin in the complex plane
```

For `3+4i`, `cmplx.Abs` computes `√(3² + 4²) = √25 = 5` — the classic 3-4-5 right triangle, using
the Pythagorean theorem, which is exactly what "magnitude of a complex number" geometrically
means.

## 🔍 Code Walkthrough (`main.go`)

```go
c1 := 3 + 4i
fmt.Printf("cmplx.Abs(3+4i) = %v ...\n", cmplx.Abs(c1))
```

`3+4i` is chosen deliberately — it's the simplest possible example where the magnitude comes out
to a clean whole number (`5`), making the Pythagorean connection immediately visible rather than
buried in decimal noise.

## ▶️ How to Run

```bash
cd level-01-fundamentals/14-complex-numbers
go run main.go
```

## ✅ Expected Output

```
=== Complex Numbers ===
----------------------------------
c1 = (3+4i) (type complex128)
c2 = (1.5-2.5i)
real(c1) = 3, imag(c1) = 4

--- Arithmetic ---
c1 + c2 = (4.5+1.5i)
c1 * c2 = (14.5-1.5i)

--- math/cmplx ---
cmplx.Abs(3+4i) = 5 (the classic 3-4-5 right triangle)

complex64 example: (1+2i) (type complex64)
```

## 🧠 Key Takeaways

- Go has built-in `complex64`/`complex128` types, with `i`-suffixed literals.
- `complex(real, imag)` builds a value explicitly; `real()`/`imag()` extract the parts back out.
- Ordinary arithmetic operators work directly on complex values, following standard complex math.
- `math/cmplx` provides complex-aware equivalents of common `math` package functions.

## 🛠️ Try It Yourself

1. Compute `c1 * c2` by hand using the `(a+bi)(c+di) = (ac-bd) + (ad+bc)i` formula, and confirm it
   matches the program's printed result.
2. Try `cmplx.Sqrt(-1)` and confirm it returns `i` (or something equal to it) — the entire reason
   complex numbers exist mathematically: giving negative numbers a square root.
3. Build a complex number representing a point on the unit circle (magnitude `1`) and confirm
   `cmplx.Abs` on it returns `1`.

## ⚠️ Common Mistakes

- Forgetting the `i` suffix on an imaginary literal (`4` instead of `4i`) — this is silently just
  the real number `4`, not part of a complex value, unless combined via `complex(...)`.
- Reaching for complex numbers when ordinary two-value math (e.g. a 2D point struct) would be
  clearer — complex numbers are the right tool specifically when you're doing genuine complex-
  plane math (signal processing, certain geometric transforms), not as a general-purpose pair type.
