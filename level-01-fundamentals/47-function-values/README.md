# 47 — Function Values

## 🎯 Learning Objectives

- Assign a function to a variable, and declare a variable of a function type.
- Write a "higher-order" function — one that accepts another function as a parameter.
- Store functions in a slice, the same way you'd store any other value.

## 📖 Concept

Functions in Go are **first-class values** — they can be assigned to variables, passed as
arguments, returned from other functions ([lesson 48](../48-anonymous-functions) covers that
side), and stored in data structures, exactly like an `int` or a `string` can be.

### Assigning a function to a variable

```go
var operation func(int, int) int
operation = add
operation(3, 4) // calls add(3, 4) through the variable
```

`func(int, int) int` here is a **function type** — it describes the shape (parameter types,
return type) a function must have to be assignable to `operation`. Writing a function's name
**without** parentheses (`add`, not `add()`) refers to the function itself as a value; adding
parentheses would call it.

### Higher-order functions: accepting a function as a parameter

```go
func applyOp(a, b int, op func(int, int) int) int {
    return op(a, b)
}
```

`applyOp`'s third parameter, `op`, is itself a function — any function matching the exact
signature `func(int, int) int` can be passed in. This is what "higher-order function" means: a
function that operates on other functions, not just on plain data.

### Storing functions in a slice

```go
operations := []func(int, int) int{add, multiply}
```

Since functions are ordinary values, a `[]func(int, int) int` is a completely normal slice type —
nothing special is needed to hold functions in a collection.

## 🔍 Code Walkthrough (`main.go`)

```go
var operation func(int, int) int
operation = add
...
operation = multiply
```

The **same** variable, `operation`, is reassigned from `add` to `multiply` partway through — proof
that a function variable genuinely holds a value that can change, exactly like any other variable,
not some special fixed binding.

## ▶️ How to Run

```bash
cd level-01-fundamentals/47-function-values
go run main.go
```

## ✅ Expected Output

```
=== Function Values ===
----------------------------------
operation = add;      operation(3, 4) = 7
operation = multiply; operation(3, 4) = 12

--- Passing functions as arguments ---
applyOp(5, 6, add)      = 11
applyOp(5, 6, multiply) = 30

--- A slice of functions ---
operations[0](2, 3) = 5
operations[1](2, 3) = 6
```

## 🧠 Key Takeaways

- Functions are first-class values in Go — assignable, passable, and storable like any other value.
- `func(paramTypes...) returnType` is a function **type**, usable anywhere a type is expected.
- A higher-order function is simply one that accepts (or returns) a function as a value.
- A function's name without parentheses refers to it as a value; with parentheses, it's a call.

## 🛠️ Try It Yourself

1. Write a third function, `subtract(a, b int) int`, and add it to the `operations` slice.
2. Write a higher-order function `applyTwice(n int, f func(int) int) int` that applies `f` to `n`
   twice in a row, and use it with a simple doubling function.
3. Try assigning a function with a **mismatched** signature to `operation` (e.g. one taking a
   `string` instead of two `int`s) and read the compiler's exact type-mismatch error.

## ⚠️ Common Mistakes

- Writing `add()` when you meant to refer to the function itself — the parentheses **call** it
  immediately (and, since `add` needs arguments, this would actually be a compile error here);
  omit them to treat the function as a plain value.
- Assuming any two functions with the "same idea" are interchangeable — Go matches function types
  **structurally and exactly**: parameter types, order, and return types must all match precisely.
