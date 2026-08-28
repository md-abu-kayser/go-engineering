# 04 — Constants

## 🎯 Learning Objectives

- Declare constants with `const`, singly and in a grouped block.
- Explain how `const` differs fundamentally from `var`: immutability and compile-time evaluation.
- Compute one constant's value from other constants.

## 📖 Concept

`const` declares a value that **never changes** for the lifetime of the program, and — unlike
`var` — must be **fully computable at compile time**, not just "assigned once."

```go
const appVersion = "v1.0.0"

const (
    maxRetries   = 3
    timeoutSecs  = 30
    defaultDebug = false
)
```

### Why `const`, not just a `var` you promise not to change

Three real differences from `var`:

1. **The compiler enforces immutability.** Attempting to assign to a `const` anywhere in your
   code is a compile error, not a runtime bug waiting to happen.
2. **Constants can be computed from other constants**, at compile time, with zero runtime cost:

   ```go
   const totalTimeout = timeoutSecs * maxRetries // computed once, at compile time
   ```

3. **Constants can be more flexible about type** than `var` — this is explored fully in
   [lesson 05](../05-typed-constants), where "untyped constants" let the same literal work as
   multiple different numeric types depending on context.

### What can and can't be a constant

A `const`'s value must be knowable **at compile time** — literals, and expressions built purely
from other constants and constant-safe operations. Things that **cannot** be constants: the
result of a function call (even one that always returns the same thing), a value read from a
file or the environment, anything computed from a variable.

```go
const x = 5 + 3        // fine — both are literals
const y = len("hello") // fine — len() of a STRING LITERAL is actually a compile-time constant
const z = someFunc()   // NOT fine, in general — function calls aren't compile-time constants
```

## 🔍 Code Walkthrough (`main.go`)

```go
const totalTimeout = timeoutSecs * maxRetries
```

This line only compiles because **both** `timeoutSecs` and `maxRetries` are themselves
constants — the entire expression is evaluated once, at compile time, and `totalTimeout` becomes
a plain constant `90`, with zero multiplication happening when the program actually runs.

## ▶️ How to Run

```bash
cd level-01-fundamentals/04-constants
go run main.go
```

## ✅ Expected Output

```
=== Constants ===
----------------------------------
appVersion   : v1.0.0
maxRetries   : 3
timeoutSecs  : 30
defaultDebug : false
greeting     : Hello
totalTimeout : 90 (computed from two other constants)
```

## 🧠 Key Takeaways

- `const` values are immutable and enforced by the compiler, not just convention.
- A constant's value must be computable entirely at compile time.
- Constants can reference other constants in their own definition.
- `const` works at both package level and inside a function, just like `var`.

## 🛠️ Try It Yourself

1. Try assigning to `maxRetries` somewhere in `main()` (`maxRetries = 5`) and read the exact
   compiler error ("cannot assign to maxRetries").
2. Try `const z = someRuntimeVariable` where `someRuntimeVariable` is an ordinary `var` — confirm
   this fails to compile, proving the "must be compile-time computable" rule.
3. Add a new constant computed from `appVersion` (e.g. its length via `len(appVersion)`) and
   confirm it works, since `len()` of a string **literal or constant** is itself a compile-time
   constant.

## ⚠️ Common Mistakes

- Trying to use `const` for something that genuinely needs to be computed at runtime (a parsed
  config value, a timestamp) — that's what `var` is for.
- Forgetting that a grouped `const (...)` block, unlike a `var (...)` block, has a special
  repetition behavior when combined with `iota` — covered fully in [lesson 06](../06-iota).
