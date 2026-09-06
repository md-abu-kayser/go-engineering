# 48 — Anonymous Functions

## 🎯 Learning Objectives

- Write a function literal — a function with no name, defined inline.
- Pass a function literal directly as an argument, with no intermediate variable.
- Write and call an IIFE (Immediately Invoked Function Expression).

## 📖 Concept

> **Scope note:** this lesson covers the **literal syntax** itself — writing a function with no
> name, wherever one is needed. Capturing variables from the *surrounding* scope (what actually
> makes a function a **closure**) is [lesson 49](../49-closures)'s dedicated topic. Every example
> here could, in principle, be rewritten as an equivalent named, top-level function — the value of
> anonymity here is purely about *not needing* a name for something used once, in one place.

### The basic form

```go
square := func(n int) int {
    return n * n
}
```

Same `func` syntax as [lesson 42](../42-functions), just with the name omitted — a **function
literal**. Assigning it to a variable makes it callable exactly like any named function
([lesson 47](../47-function-values) already covered function values generally; this is simply
where the value itself came from).

### Passed directly as an argument

```go
mapInts(numbers, func(n int) int {
    return n * 2
})
```

This is one of the most common uses: a small, one-off piece of logic that's only ever needed at
this exact call site doesn't need a separate, named top-level function cluttering the package —
define it right where it's used.

### IIFE: define and call in one statement

```go
result := func(a, b int) int {
    sum := a + b
    return sum * sum
}(3, 4)
```

The `(3, 4)` immediately after the closing `}` calls the function literal **right away** — useful
for scoping a small piece of setup logic (temporary variables like `sum` here) without leaking
them into the surrounding function, or without needing a separate named helper for something used
exactly once.

## 🔍 Code Walkthrough (`main.go`)

```go
func mapInts(nums []int, fn func(int) int) []int {
    result := make([]int, len(nums))
    for i, n := range nums {
        result[i] = fn(n)
    }
    return result
}
```

`mapInts` itself is an ordinary, named function — it's the **argument passed to it** (`func(n
int) int { return n * 2 }`) that's anonymous. This mirrors a genuinely common real pattern:
reusable, named "do something to every element" helpers, paired with small, throwaway,
call-site-specific logic passed in anonymously.

## ▶️ How to Run

```bash
cd level-01-fundamentals/48-anonymous-functions
go run main.go
```

## ✅ Expected Output

```
=== Anonymous Functions ===
----------------------------------
square(5) = 25

--- Passed directly as an argument ---
mapInts([1 2 3 4 5], double) = [2 4 6 8 10]

--- Immediately invoked (IIFE) ---
IIFE result = 49 ((3+4)^2)
```

## 🧠 Key Takeaways

- A function literal is `func(...) ... { ... }` with no name — usable anywhere a value is expected.
- Passing one directly as an argument avoids cluttering a package with narrowly-used named helpers.
- An IIFE defines and calls a function literal in one statement — useful for scoping setup logic.
- Every anonymous function here could be rewritten as a named one — anonymity is a convenience,
  not a different capability (that distinction belongs to closures, [lesson 49](../49-closures)).

## 🛠️ Try It Yourself

1. Rewrite the `mapInts` call's anonymous function as a separate, named `double` function instead,
   and pass that by name — confirm the output is identical.
2. Write your own `filterInts(nums []int, fn func(int) bool) []int` and call it with an anonymous
   function checking for even numbers.
3. Write an IIFE that takes no arguments at all and simply returns a computed constant — the
   simplest possible "run this once, right here" pattern.

## ⚠️ Common Mistakes

- Forgetting the calling parentheses on an IIFE (`func(a, b int) int { ... }` with nothing after
  it) — without them, you've only **defined** the function literal as an unused value, never
  actually called it.
- Overusing anonymous functions for logic that's genuinely reused in multiple places — if the same
  anonymous function body starts appearing more than once, that's a signal it deserves to become
  a real, named function instead.
