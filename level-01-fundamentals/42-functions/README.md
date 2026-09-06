# 42 — Functions

## 🎯 Learning Objectives

- Declare a function with Go's exact syntax: name, typed parameters, and a return type.
- Use the shorthand for consecutive parameters that share a type.
- Declare a function with no parameters, and one with no return value.

## 📖 Concept

```go
func add(a int, b int) int {
    return a + b
}
```

The shape is always: `func`, the function's name, parentheses containing its parameters (each
with an explicit type — Go never infers a parameter's type from how it's called), then the return
type, then the body in braces.

### Shared-type parameter shorthand

```go
func multiply(a, b int) int {
```

When consecutive parameters share the same type, you can write the type once, after the last name
in that group — `a, b int` means exactly the same thing as `a int, b int`. This is purely a
notational shortcut; it changes nothing about how the function behaves.

### No return value: just omit the return type

```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

Go has no `void` keyword — a function that doesn't return anything simply has nothing written
after its parameter list, before the opening `{`.

### No parameters: empty parentheses

```go
func currentVersion() string {
    return "v1.0.0"
}
```

Straightforward — an empty parameter list is just `()`, exactly as you'd expect.

## 🔍 Code Walkthrough (`main.go`)

Every variation covered above — full explicit types, the shared-type shorthand, no return value,
no parameters — appears once in this lesson's `main.go`, specifically so you see all four shapes
side by side rather than reading about each in isolation.

## ▶️ How to Run

```bash
cd level-01-fundamentals/42-functions
go run main.go
```

## ✅ Expected Output

```
=== Functions ===
----------------------------------
add(3, 4)      = 7
multiply(3, 4) = 12
Hello, Gopher!
currentVersion() = v1.0.0
```

## 🧠 Key Takeaways

- Every parameter needs an explicit type — Go never infers a parameter's type from a call site.
- Consecutive same-typed parameters can share one type declaration: `a, b int`.
- No `void` keyword — omit the return type entirely for a function that returns nothing.
- Empty parentheses `()` mean "no parameters," exactly as expected.

## 🛠️ Try It Yourself

1. Add a `subtract(a, b int) int` function using the shared-type shorthand, and call it from `main`.
2. Write a function with three parameters of two different types (e.g. `name string, age, id int`)
   and figure out where the shorthand can and can't apply.
3. Write a function that takes no parameters and returns nothing at all — what's the smallest
   possible valid function signature in Go?

## ⚠️ Common Mistakes

- Forgetting a parameter's type when it "obviously" matches a neighboring parameter that already
  has one written after it, instead of before — the shorthand only works when the type comes
  **after** the group of same-typed names, not before.
- Reaching for a `void`-like keyword out of habit from another language — Go simply has no return
  type written at all for a function that returns nothing.
