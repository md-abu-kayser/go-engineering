# 01 — var Declarations

## 🎯 Learning Objectives

- Declare a variable with `var`, with and without an explicit type.
- Declare several variables together in a `var (...)` block.
- Declare a package-level variable, and know why `var` can do this when `:=` cannot.

## 📖 Concept

`var` is Go's general-purpose variable declaration keyword. It has more forms than you might
expect at first, each useful in a different situation.

### Explicit type

```go
var age int = 15
```

Spells out both the name and the type — useful when the type isn't obvious from the value, or
when you specifically want a **different** type than what would be inferred (e.g. `var x float64
= 5`, where `5` alone would otherwise default to `int`).

### Inferred type

```go
var city = "Dhaka"
```

Still `var`, but the type is inferred from the value — `city` is a `string`, exactly as if you'd
written `var city string = "Dhaka"`.

### No initial value

```go
var score int
```

Declares `score` with **no** initializer — it's automatically set to its type's **zero value**
(`0` for `int`). [Lesson 03](../03-zero-values) covers this in full.

### Grouped, with `var (...)`

```go
var (
    width  = 1920
    height = 1080
    title  = "Monitor"
)
```

The idiomatic way to declare several related variables together, without repeating `var` on every
line — `gofmt` aligns them automatically.

### Package-level `var`

```go
var appName string = "GO-ENGINEERING"
```

`var` can appear **outside** any function, at package level — this is a genuinely important
distinction from `:=` ([lesson 02](../02-short-declarations)), which can only be used **inside**
a function body.

## 🔍 Code Walkthrough (`main.go`)

Every form above appears in this lesson's `main.go`, deliberately side by side — `appName` at
package level, then all four in-function forms inside `main()` — so you can compare them
directly rather than reading about each in isolation.

## ▶️ How to Run

```bash
cd level-01-fundamentals/01-var-declarations
go run main.go
```

## ✅ Expected Output

```
=== var Declarations ===
----------------------------------
appName (package-level) : GO-ENGINEERING
age (explicit type)     : 15
city (inferred type)    : Dhaka
score (no initializer)  : 0
grouped: Monitor is 1920x1080
```

## 🧠 Key Takeaways

- `var name type = value` is the most explicit form; `var name = value` infers the type.
- `var name type` with no value gets that type's zero value automatically.
- `var (...)` groups multiple declarations idiomatically.
- `var` works at package level; `:=` ([lesson 02](../02-short-declarations)) does not.

## 🛠️ Try It Yourself

1. Add a `var pi float64 = 3.14159` and confirm `%T` (see [lesson 06 of level 00]
   (../../level-00-getting-started/06-fmt-printing)) reports its type as `float64`.
2. Remove the explicit type from `var age int = 15`, leaving `var age = 15`, and confirm the
   program behaves identically.
3. Try moving `var score int` to package level (outside `main`) and confirm it still compiles —
   `var` is equally at home in both places.

## ⚠️ Common Mistakes

- Writing `var x int = "hello"` — mixing an explicit type with an incompatible value is a compile
  error, not a silent conversion.
- Forgetting that a package-level `var` is visible to **every** function in the package, which can
  make code harder to reason about if overused — prefer function-local variables unless something
  genuinely needs to be shared package-wide.
