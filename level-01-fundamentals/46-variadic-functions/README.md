# 46 — Variadic Functions

## 🎯 Learning Objectives

- Declare a function that accepts a variable number of arguments with `...T`.
- Call a variadic function with individual arguments, or by spreading an existing slice.
- Know that a variadic parameter must be the function's last parameter.

## 📖 Concept

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

`...int` means "accept any number of `int` arguments — zero, one, or many." Inside the function
body, `nums` is simply an ordinary `[]int` — the variadic syntax only changes how **callers** can
invoke the function, not how the parameter behaves once inside it.

### Calling a variadic function

```go
sum()             // zero arguments — nums is an empty (nil) slice inside
sum(5)              // one argument
sum(1, 2, 3, 4)      // several — list them individually, comma-separated
```

### Mixing ordinary and variadic parameters

```go
func joinWithPrefix(prefix string, items ...string) string {
```

Ordinary parameters can come before a variadic one — but the variadic parameter **must be last**;
Go's grammar doesn't allow anything after `...T` in a parameter list.

### Spreading an existing slice: `slice...`

```go
nums := []int{10, 20, 30, 40}
sum(nums...) // passes the slice's elements AS IF they were individual arguments
```

If you already have a `[]int` and want to pass its contents to a variadic function, you don't
need to unpack it manually — the `...` suffix **after** a slice at a call site "spreads" it,
passing each element as if you'd written them out individually. This is the same `...` symbol as
the declaration, but it means something different depending on which side it appears on.

### `fmt.Println` is variadic too — this isn't a special case

```go
fmt.Println("this", "call", "has", "five", "separate", "arguments")
```

Every `fmt.Println`/`fmt.Printf` call you've made throughout this entire repository has been
calling a variadic function — `fmt.Println`'s actual signature is
`func Println(a ...any) (n int, err error)`. Nothing about variadic functions is exclusive to
user-defined code.

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Printf("sum(nums...) = %d (spread an existing []int)\n", sum(nums...))
```

This line demonstrates both directions of `...` in the same program: `sum`'s own declaration uses
`...int` to **accept** a variable number of arguments, while `nums...` at this call site
**spreads** an already-built slice into that same call.

## ▶️ How to Run

```bash
cd level-01-fundamentals/46-variadic-functions
go run main.go
```

## ✅ Expected Output

```
=== Variadic Functions ===
----------------------------------
sum()        = 0
sum(5)       = 5
sum(1,2,3,4) = 10

Items: apple banana cherry

--- Spreading a slice with ... ---
sum(nums...) = 100 (spread an existing []int)

--- fmt.Println is variadic too ---
this call has five separate arguments
```

## 🧠 Key Takeaways

- `...T` in a parameter list accepts any number of `T` arguments, exposed as a `[]T` inside.
- A variadic parameter must be the function's last parameter.
- `slice...` at a **call site** spreads an existing slice's elements into individual arguments.
- `fmt.Println`/`fmt.Printf` are themselves variadic functions — nothing special about them.

## 🛠️ Try It Yourself

1. Try calling `sum` with both a spread slice **and** an extra individual argument
   (`sum(nums..., 5)`) and read the compiler's exact error — spreading and individual arguments
   can't be mixed in one call.
2. Write a variadic function that returns the **maximum** of any number of `int` arguments,
   handling the zero-arguments case sensibly.
3. Try declaring a function with a variadic parameter that ISN'T last (e.g.
   `func bad(nums ...int, label string)`) and read the compiler's error confirming the "must be
   last" rule.

## ⚠️ Common Mistakes

- Trying to mix spread and individual arguments in one call — `fn(slice..., extra)` is not valid
  Go; a call either spreads one slice, or lists individual arguments, never both together.
- Forgetting a variadic function receives a `nil` slice (not a slice of zero-valued elements) when
  called with zero arguments — matching [lesson 03](../03-zero-values)'s zero-value coverage for
  slices, and safe to range over just like any other `nil` slice.
