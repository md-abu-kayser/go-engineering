# 51 — Scope

## 🎯 Learning Objectives

- Identify the **block** as Go's fundamental unit of scope.
- Predict exactly where a variable is visible, given where it was declared.
- Understand that nested blocks can see outward (into enclosing blocks), but never inward.

## 📖 Concept

A variable in Go is visible from the point it's declared until the end of the **innermost
enclosing block** — where a "block" is anything delimited by `{ }`: a function body, an `if`
body, a `for` body, or even a bare `{ }` with nothing else around it.

```go
func classifyScore(score int) string {
    label := "unknown"       // scoped to the WHOLE function body

    if score >= 90 {
        bonus := " (with honors)" // scoped to JUST this if-block
        label = "A" + bonus
    }
    // `bonus` does not exist here — this line is outside its block

    return label
}
```

This is the same underlying rule behind several things earlier lessons already touched on
individually: `if`/`switch`/`for`'s init-statement variables
([lesson 25](../25-if-with-init), [lesson 28](../28-expression-switches),
[lesson 30](../30-for-loop)) are scoped to their statement precisely because that statement
introduces its own block; this lesson names the **general** rule those were all specific cases of.

### Nested blocks see outward, never inward

```go
outer := "outer value"
{
    fmt.Println(outer)   // fine — inner block sees the ENCLOSING block's variables
    inner := "inner value"
}
fmt.Println(inner) // COMPILE ERROR — outer block cannot see INTO the inner block
```

Think of scope as nested boundaries: code inside a block can always reach outward to anything
declared in an enclosing block (as long as it was declared **before** the inner block began), but
nothing declared inside an inner block leaks back out once that block ends.

## 🔍 Code Walkthrough (`main.go` and the test file)

```go
func classifyScore(score int) string {
    label := "unknown"
    if score >= 90 {
        bonus := " (with honors)"
        label = "A" + bonus
    } else if score >= 70 {
        label = "B"
    }
    return label
}
```

`classifyScore` is deliberately written as a small, pure function specifically so this lesson can
verify its scoping-dependent behavior with a **real test**, rather than only printed output — the
test in `51_scope_test.go` checks every branch, including the boundary case (`score == 90`) where
the honors bonus should and shouldn't apply.

## ▶️ How to Run

```bash
cd level-01-fundamentals/51-scope
go run main.go
go test -v ./...
```

## ✅ Expected Output

```
=== Scope ===
----------------------------------
classifyScore(95) = A (with honors)
classifyScore(75) = B
classifyScore(40) = unknown

--- Block scope, made visible ---
  inside loop: i=0, x=0
  inside loop: i=1, x=1
  inside loop: i=2, x=4
  outside the loop: neither i nor x exist here anymore

--- Nested blocks see outward, never inward ---
  inner block CAN see: "I'm from the outer block"
  inner block can also see its own: "I'm from the inner block"
  outer block canNOT see `inner` — it never existed out here
```

## 🧠 Key Takeaways

- A block (`{ }`) is Go's fundamental unit of scope — function bodies, `if`/`for`/`switch`
  bodies, and bare `{ }` blocks all introduce one.
- A variable is visible from its declaration to the end of its innermost enclosing block.
- Inner blocks can see outward into enclosing blocks; outer blocks can never see into inner ones.
- This single rule explains why every init-statement variable from earlier lessons
  (`if`/`switch`/`for`) is scoped the way it is.

## 🛠️ Try It Yourself

1. Add a `bonus` variable **outside** the `if` block in `classifyScore`, at the top with `label`,
   and confirm it's now visible in the `else if` branch too — contrasting with the original,
   block-scoped version.
2. Add a bare `{ }` block (not attached to any `if`/`for`) directly inside `main`, declare a
   variable inside it, and confirm it's inaccessible immediately after the block closes.
3. Run `go test -v ./...` and confirm every subtest passes, including the `score == 90` boundary
   case.

## ⚠️ Common Mistakes

- Declaring a variable inside a block "just in case it's needed later," then being surprised it's
  unavailable a few lines below the block's closing brace — move the declaration to the
  appropriate enclosing scope instead.
- Assuming a deeply nested block can't see variables from several levels up — nesting depth
  doesn't matter; any enclosing block's variables remain visible all the way down, as long as
  they were declared before the inner block began.
