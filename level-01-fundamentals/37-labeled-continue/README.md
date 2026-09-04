# 37 — Labeled continue

## 🎯 Learning Objectives

- Attach a label to an outer loop and use `continue label` to target it from a nested inner loop.
- Contrast labeled `continue` directly against [lesson 35](../35-continue)'s plain, inner-loop-only
  behavior.
- Know when abandoning the rest of an inner loop's current pass is genuinely the right tool.

## 📖 Concept

[Lesson 35](../35-continue) showed that a plain `continue` inside a nested inner loop only skips
ahead within that inner loop — the outer loop's current iteration keeps running normally
afterward. A **labeled continue** changes this: it skips straight to the **outer** loop's next
iteration, abandoning whatever remained of the inner loop's current pass entirely.

```go
rows:
for row := 0; row < 4; row++ {
    for col := 0; col < 4; col++ {
        if someCondition {
            continue rows // jump straight to the OUTER loop's next row — skip the rest of THIS row
        }
    }
}
```

This is the exact labeling mechanism from [lesson 36](../36-labeled-break), used with `continue`
instead of `break` — same label syntax, different target action (skip ahead vs. exit entirely).

### Side-by-side: plain vs. labeled continue

| | Plain `continue` | Labeled `continue outerLabel` |
|---|---|---|
| Skips the rest of | the **innermost** loop's current iteration only | the **entire outer** iteration, abandoning any remaining inner-loop work |
| Outer loop's remaining inner passes | still run normally | **skipped entirely** for that outer iteration |

## 🔍 Code Walkthrough (`main.go`)

```go
rows:
for row := 0; row < 4; row++ {
    for col := 0; col < 4; col++ {
        if col == 2 && row%2 == 0 {
            continue rows
        }
        fmt.Printf("  row=%d, col=%d\n", row, col)
    }
    fmt.Printf("  (finished row %d normally)\n", row)
}
```

For even rows, once `col` reaches `2`, `continue rows` jumps straight to the next `row` — notice
`col == 3` for those rows **never prints**, and the `"finished row %d normally"` line is also
skipped for them, since `continue rows` abandons the rest of that outer iteration entirely,
including code after the inner loop.

The second example in `main.go` repeats the same grid with a **plain** `continue` instead, so you
can see `col == 3` (and the "finished row" line) print normally there — a direct, side-by-side
contrast with the labeled version above it.

## ▶️ How to Run

```bash
cd level-01-fundamentals/37-labeled-continue
go run main.go
```

## ✅ Expected Output

```
=== Labeled continue ===
----------------------------------
--- Skipping an entire outer iteration from an inner loop ---
  row=0, col=0
  row=0, col=1
  row=0, col=2: skipping the REST of this row entirely
  row=1, col=0
  row=1, col=1
  row=1, col=2
  row=1, col=3
  (finished row 1 normally)
  row=2, col=0
  row=2, col=1
  row=2, col=2: skipping the REST of this row entirely
  row=3, col=0
  row=3, col=1
  row=3, col=2
  row=3, col=3
  (finished row 3 normally)

--- Contrast: a PLAIN continue only skips within the inner loop ---
  row=0, col=0
  row=0, col=1
  row=0, col=3
  (finished row 0 — notice col=3 STILL printed, unlike the labeled version above)
  row=1, col=0
  row=1, col=1
  row=1, col=3
  (finished row 1 — notice col=3 STILL printed, unlike the labeled version above)
```

## 🧠 Key Takeaways

- `continue label` skips straight to the **labeled outer loop's** next iteration, abandoning any
  remaining inner-loop work for that pass entirely.
- This is a genuinely different effect from a plain `continue`, which only affects the innermost
  loop it's written in ([lesson 35](../35-continue)).
- The label syntax itself is identical to [lesson 36](../36-labeled-break)'s — only the keyword
  (`break` vs `continue`) changes what happens once the label is reached.

## 🛠️ Try It Yourself

1. Change the condition from `col == 2 && row%2 == 0` to trigger on **every** row, and confirm
   `col == 3` never prints for any row in the labeled version.
2. Add a third nesting level and use `continue` labeled at the **outermost** loop from the
   innermost one — confirm it skips both remaining inner levels for that outer pass.
3. Rewrite the labeled example using a boolean flag and a plain `continue` instead of a label, and
   compare which version you find clearer.

## ⚠️ Common Mistakes

- Confusing labeled `continue`'s effect with labeled `break`'s — `continue label` still runs the
  labeled loop's **next** iteration; it does not exit the loop the way `break label` would.
- Reaching for labeled continue/break as a first resort for complex nested logic, when extracting
  the inner loop into its own named function (using a plain early `return` instead of a label)
  might read more clearly.
