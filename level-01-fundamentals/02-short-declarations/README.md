# 02 — Short Declarations

## 🎯 Learning Objectives

- Use `:=` to declare one or more local variables with inferred types.
- Know the "at least one new variable" rule that lets `:=` sometimes look like redeclaration.
- Understand variable shadowing — a genuinely new variable with the same name in a nested scope.

## 📖 Concept

`:=` is Go's **short variable declaration** — it declares a variable and infers its type from the
value, all in one step, and is by far the most common way to declare variables inside a function
body.

```go
name := "Gopher"
// exactly equivalent to:
var name = "Gopher"
```

### It only works inside a function

```go
name := "Gopher"   // fine INSIDE main() or any function

// but NOT at package level — that requires var (lesson 01)
```

This is the single most important distinction from `var`: `:=` is a **statement**, valid only
inside a function body; `var` is a **declaration**, valid both inside functions and at package
level.

### Multiple variables at once

```go
city, country := "Dhaka", "Bangladesh"
```

Extremely common with functions returning `(value, error)` or `(value, ok)` pairs — you'll write
`result, err := someFunc()` constantly in idiomatic Go.

### The "at least one new variable" rule

```go
age, job := 16, "Student"
```

If `age` already exists in the **same scope**, this doesn't redeclare it — it **assigns** to the
existing `age` and **declares** the new `job`. Go allows this specifically because it's so common
with multi-value returns. The rule: `:=` is legal as long as **at least one** name on the left is
genuinely new **in that scope** — if all names already exist, you'd need a compile error, since
there'd be nothing left to actually declare.

### Shadowing

```go
name := "Gopher"
{
    name := "Shadowed Gopher"  // a BRAND NEW variable, not a reassignment
}
```

Every `{ }` block introduces a new scope. `:=` inside a nested scope, using a name that already
exists in an outer scope, creates a genuinely **separate** variable — the outer `name` is
completely unaffected. This is called shadowing, and while intentional shadowing is sometimes
useful, accidental shadowing is a real, common source of bugs.

## 🔍 Code Walkthrough (`main.go`)

```go
age, job := 16, "Student"
```

Notice `age` was already declared a few lines earlier with `age := 15`. This line **reuses** that
same variable (updating it to `16`) while genuinely declaring `job` for the first time — both in
one `:=` statement, legal because of the "at least one new" rule above.

## ▶️ How to Run

```bash
cd level-01-fundamentals/02-short-declarations
go run main.go
```

## ✅ Expected Output

```
=== Short Declarations ===
----------------------------------
name: Gopher, age: 15
city: Dhaka, country: Bangladesh
age (updated): 16, job (new): Student
inside block, name: Shadowed Gopher
outside block, name is still: Gopher
```

## 🧠 Key Takeaways

- `:=` infers the type and works only inside function bodies, never at package level.
- `a, b := x, y` declares multiple variables at once — common with `(value, error)` returns.
- `:=` can mix reassignment and new declaration in one statement, as long as at least one name
  is genuinely new in that scope.
- A nested `{ }` block's `:=` shadows an outer variable of the same name — a separate variable,
  not the same one.

## 🛠️ Try It Yourself

1. Try `age, job := 17, "Student"` a second time in the **same** scope with **no** new names on
   the left, and read the exact compiler error ("no new variables on left side of :=").
2. Remove the inner `{ }` block's `name :=` line, replacing it with plain `name = "..."` instead
   (no colon), and confirm this **does** modify the outer `name` — the opposite of shadowing.
3. Deliberately shadow a variable inside an `if` block's scope, and confirm the outer variable is
   unaffected once the `if` block ends.

## ⚠️ Common Mistakes

- Trying to use `:=` at package level — this is a compile error; use `var` there instead
  ([lesson 01](../01-var-declarations)).
- Accidental shadowing inside an `if`/`for`/`{}` block — declaring `err := ...` inside a nested
  block when you meant to reuse an outer `err` is one of the most common real-world Go bugs,
  since the outer variable silently stays unchanged.
