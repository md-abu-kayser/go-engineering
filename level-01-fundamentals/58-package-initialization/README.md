# 58 — Package Initialization

## 🎯 Learning Objectives

- Know the exact order Go initializes a package: package-level variables (in dependency order),
  then `init()` functions, then `main()`.
- Use multiple `init()` functions in one package, and know they all run automatically.
- Confirm package-level variables are initialized in **dependency order**, not source-code order.

## 📖 Concept

Before `main()` ever runs ([lesson 57](../57-main-function)), a Go program's `main` package goes
through a strict, automatic initialization sequence:

1. **Package-level variables** are initialized first — in an order Go determines based on their
   **dependencies** on each other, not necessarily the order they're written in the source file.
2. **Every `init()` function** in the package then runs — in the order they appear (across
   possibly multiple files, in filename order; within one file, top to bottom).
3. **Only then** does `main()` finally start.

### Package-level variables: dependency order, not source order

```go
var total = base * multiplier // WRITTEN first, but DEPENDS on the two below
var base = 10
var multiplier = 3
```

Even though `total` is declared **before** `base` and `multiplier` in the source, Go recognizes
that computing `total` requires their values, and automatically initializes `base` and
`multiplier` **first** — the dependency graph, not the text order, determines what happens when.

### `init()`: automatic, no explicit call, and there can be several

```go
func init() {
    fmt.Println("first init")
}

func init() {
    fmt.Println("second init")
}
```

A package can declare **any number** of `init()` functions — even multiple in the same file, as
shown here. Every single one runs automatically before `main()` starts; you never call `init()`
yourself, and in fact **can't** — calling `init()` explicitly is not valid Go.

### Full ordering, put together

```
package-level vars (dependency order) -> init() #1 -> init() #2 -> ... -> main()
```

By the time either `init()` function runs, every package-level variable is already fully
initialized — which is exactly what this lesson's first `init()` demonstrates by printing
`total`'s already-computed value.

## 🔍 Code Walkthrough (`main.go`)

```go
var total = base * multiplier
var base = 10
var multiplier = 3
```

This ordering is deliberately "backwards" from what dependency order would suggest, specifically
to make the point concrete: if Go simply initialized variables top-to-bottom as written, `total`
would be computed from `base`/`multiplier`'s **zero values** (`0 * 0 = 0`), not `10 * 3 = 30`. The
correct output (`30`) is direct proof Go resolved the dependency correctly rather than just
running top to bottom.

## ▶️ How to Run

```bash
cd level-01-fundamentals/58-package-initialization
go run main.go
```

## ✅ Expected Output

```
init() #1: package-level vars are ALREADY initialized here:
  total = 30 (computed from base=10, multiplier=3)
init() #2: multiple init() functions run in the order they APPEAR in the file

=== Package Initialization ===
----------------------------------
main() runs AFTER every package-level var AND every init() function above.
total is still 30 here in main() — nothing re-runs it.
```

Notice both `init()` messages print **before** anything from `main()` — including `main()`'s own
first `fmt.Println`, which is textually the very first line of the function but still runs after
every `init()`.

## 🧠 Key Takeaways

- Initialization order is: package-level vars (dependency order) → `init()` functions (appearance
  order) → `main()`.
- Package-level variables are ordered by their actual dependencies, not by where they're written.
- A package can have multiple `init()` functions, even in one file — all run automatically, with
  no explicit call (and no way to call one yourself).
- `init()` functions are guaranteed to see every package-level variable already fully initialized.

## 🛠️ Try It Yourself

1. Reorder `base`, `multiplier`, and `total`'s declarations so `total` is written **last**, and
   confirm the output is completely unchanged — proving the order is genuinely dependency-based,
   not just "Go got lucky" with the original ordering.
2. Add a third `init()` function between the two existing ones and confirm it runs in the position
   you'd expect, based on where it appears in the file.
3. Add a `fmt.Println` as the very first line inside `main()` (before anything else) and confirm
   it still prints **after** both `init()` functions' output.

## ⚠️ Common Mistakes

- Assuming package-level variables initialize strictly top-to-bottom, as written — Go resolves
  dependencies between them first, which can reorder things in ways that only matter when one
  variable's initializer depends on another's value.
- Reaching for `init()` as a place to put "setup code" by default — for anything a caller might
  reasonably want control over (whether it runs, with what configuration, in what order relative
  to other setup), an explicit function called from `main` is usually clearer than an automatic,
  invisible `init()`.
