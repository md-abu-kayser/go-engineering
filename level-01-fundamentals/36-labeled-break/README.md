# 36 — Labeled break

## 🎯 Learning Objectives

- Attach a label to a loop, and use `break label` to target it specifically.
- Fix [lesson 34](../34-break)'s switch-inside-loop gotcha using a label.
- Use a labeled break to exit multiple nested loops at once.

## 📖 Concept

[Lesson 34](../34-break) showed the problem: a bare `break` inside a `switch` nested in a loop
only exits the switch, since `break` targets its **nearest** enclosing construct. A **label**
solves this by giving `break` an explicit, unambiguous target.

### Labeling a loop

```go
outer:
for i := 0; i < 5; i++ {
    // "outer" now refers to THIS loop specifically
}
```

A label is just an identifier followed by a colon, placed directly before the statement it names
— conventionally a loop. It doesn't change how the loop runs on its own; it just gives it a name
other statements can refer to.

### `break label`: targeting a specific loop, regardless of nesting

```go
outer:
for i := 0; i < 5; i++ {
    switch {
    case i == 2:
        break outer // exits the LOOP NAMED "outer" — not just the switch
    }
}
```

`break outer` explicitly says "exit the loop labeled `outer`," resolving the ambiguity a bare
`break` had inside the switch — this is precisely [lesson 34](../34-break)'s gotcha, fixed.

### Breaking out of multiple nested loops at once

```go
search:
for row := 0; row < 3; row++ {
    for col := 0; col < 3; col++ {
        if found {
            break search // exits BOTH loops immediately
        }
    }
}
```

This is the other major use case for labeled break: a bare `break` in the inner loop would only
exit that inner loop, leaving the outer loop to continue with its next iteration. Labeling the
**outer** loop and breaking to that label exits both at once — exactly the right tool for a
"search a grid, stop everywhere once found" pattern.

## 🔍 Code Walkthrough (`main.go`)

```go
outer:
for i := 0; i < 5; i++ {
    switch {
    case i == 2:
        break outer
    ...
    }
    fmt.Printf("  (this line will NOT print for i=2 ...)\n")
}
fmt.Println("  loop fully exited")
```

Compare this directly against [lesson 34](../34-break)'s output: there, the "loop continues"
line printed for **every** iteration, including `i == 2`. Here, with `break outer` instead of a
bare `break`, that line is correctly skipped once the label is targeted — proof the loop itself,
not just the switch, actually exited this time.

## ▶️ How to Run

```bash
cd level-01-fundamentals/36-labeled-break
go run main.go
```

## ✅ Expected Output

```
=== Labeled break ===
----------------------------------
--- Fixing lesson 34's gotcha with a label ---
  i=0: default case
  (this line will NOT print for i=2 — the loop already exited)
  i=1: default case
  (this line will NOT print for i=2 — the loop already exited)
  i=2: breaking the OUTER LOOP this time, via the label
  loop fully exited

--- Breaking out of nested loops entirely ---
  checking row=0, col=0
  checking row=0, col=1
  checking row=0, col=2
  checking row=1, col=0
  found target at row=1, col=1 — stopping BOTH loops
  search finished
```

## 🧠 Key Takeaways

- A label (`name:`) placed before a loop gives it an explicit name other statements can target.
- `break label` exits the labeled loop specifically, regardless of how deeply nested the `break`
  itself is (inside a switch, an inner loop, or both).
- This is the direct fix for [lesson 34](../34-break)'s switch-inside-loop gotcha.
- Labeled break is also the standard way to exit multiple nested loops at once, for a
  "search until found, then stop everywhere" pattern.

## 🛠️ Try It Yourself

1. Remove the label from the first example (back to a bare `break`) and confirm you're back to
   [lesson 34](../34-break)'s original gotcha behavior.
2. In the grid-search example, change the target to `row == 2, col == 2` and confirm the search
   correctly visits every earlier cell before stopping.
3. Add a third level of nesting (a loop inside the inner loop) and use a label to break all three
   at once from the innermost level.

## ⚠️ Common Mistakes

- Forgetting the label must be placed directly before the loop statement itself — a label on the
  wrong line, or with anything else in between, won't compile as intended.
- Overusing labeled break/continue for logic that could be more clearly expressed as a separate
  function with an early `return` instead — labels are the right tool for genuine multi-level
  loop control, not a substitute for restructuring genuinely tangled logic.
