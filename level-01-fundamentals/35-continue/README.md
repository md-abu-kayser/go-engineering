# 35 — continue

## 🎯 Learning Objectives

- Use `continue` to skip a loop iteration's remaining body.
- Understand that in nested loops, `continue` always targets the **innermost** enclosing loop.
- Know this is a deliberate default — matching how `break` also targets its nearest construct
  ([lesson 34](../34-break)).

## 📖 Concept

[Lesson 30](../30-for-loop) already introduced the basics:

```go
for i := 0; i < 5; i++ {
    if i%2 == 0 {
        continue // skip the rest of THIS iteration's body; go to the post clause
    }
    fmt.Println(i)
}
```

### Nested loops: `continue` only affects the innermost loop

```go
for row := 0; row < 3; row++ {
    for col := 0; col < 4; col++ {
        if col == 1 {
            continue // affects ONLY the inner (col) loop
        }
        // ...
    }
    // the outer (row) loop's own iteration finishes normally regardless
}
```

Just as [lesson 34](../34-break) showed `break` targeting its nearest enclosing construct,
`continue` follows the same principle: written inside a nested loop, it **only** skips ahead in
that **innermost** loop — the outer loop's own iteration count and behavior are completely
unaffected.

### When you need the outer loop instead: labeled continue

If you genuinely need to skip ahead in an **outer** loop from inside a nested inner loop, that
requires a **labeled continue** — covered fully in [lesson 37](../37-labeled-continue).

## 🔍 Code Walkthrough (`main.go`)

```go
for row := 0; row < 3; row++ {
    fmt.Printf("row %d: ", row)
    for col := 0; col < 4; col++ {
        if col == 1 {
            continue
        }
        fmt.Printf("%d ", col)
    }
    fmt.Println()
}
```

Every row still prints its own line (`fmt.Println()` after the inner loop), and every row still
runs the full range of `col` values it should (just skipping `1` specifically) — direct evidence
that the inner `continue` never touched the outer `row` loop's own progress at all.

## ▶️ How to Run

```bash
cd level-01-fundamentals/35-continue
go run main.go
```

## ✅ Expected Output

```
=== continue ===
----------------------------------
--- continue in a plain loop ---
  i = 1 (odd)
  i = 3 (odd)

--- continue in nested loops (targets the INNER loop only) ---
row 0: 0 2 3 
row 1: 0 2 3 
row 2: 0 2 3 

See lesson 37 (labeled continue) for how to continue an OUTER loop
from inside a nested inner loop.
```

## 🧠 Key Takeaways

- `continue` skips the rest of the current iteration's body and moves to the loop's next iteration.
- In nested loops, `continue` always targets the **innermost** loop it's written inside.
- This mirrors `break`'s "nearest enclosing construct" rule from [lesson 34](../34-break).
- [Lesson 37](../37-labeled-continue) covers the labeled form, for targeting an outer loop
  specifically.

## 🛠️ Try It Yourself

1. Add a second `continue` condition to the inner loop (skip `col == 3` too) and confirm each row
   still prints its full expected set of remaining values.
2. Move the `continue` to the **outer** loop instead (skip `row == 1` entirely) and confirm the
   inner loop for that row never runs at all.
3. Predict, before running it, what happens with a `continue` inside a `switch` that's inside a
   loop — does `continue` have the same "nearest construct" ambiguity that `break` does? (Hint:
   `switch` isn't a valid target for `continue` at all — think about why that resolves the
   ambiguity `break` has.)

## ⚠️ Common Mistakes

- Expecting `continue` inside a nested inner loop to advance the **outer** loop's iteration —
  it only ever affects the loop it's directly written inside.
- Assuming `continue`'s behavior with `switch` mirrors `break`'s gotcha ([lesson 34](../34-break))
  — it doesn't, because `continue` can only ever target a loop in the first place; there's no
  "which nearby construct did you mean" ambiguity the way there is with `break`.
