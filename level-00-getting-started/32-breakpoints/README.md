# 32 — Breakpoints

## 🎯 Learning Objectives

- Set a **conditional breakpoint** in Delve — one that only stops when an expression is true.
- Set a **tracepoint** — one that logs instead of stopping.
- List, disable, and clear breakpoints during a session.

## 📖 Concept

[Lesson 31](../31-debugging-with-delve) set a plain breakpoint that stops **every** time
execution reaches that line. That's fine for a function called once, but `findFirstOver`'s loop
in this lesson runs up to seven times — stopping on every iteration to manually check "is this
the one I care about?" is slow and error-prone. Delve breakpoints can be smarter than that.

### Conditional breakpoints

```bash
dlv debug .
```

```
(dlv) break main.go:12
Breakpoint 1 set at ... for main.findFirstOver() ./main.go:12
(dlv) condition 1 n > 8
(dlv) continue
```

Now execution only actually **stops** when `n > 8` is true for that specific loop iteration —
every other pass through the loop runs invisibly, at full speed, until the condition matches.

### Tracepoints — log without stopping

Sometimes you don't want to pause at all, just observe:

```
(dlv) trace main.go:12
(dlv) continue
```

Each time line 12 executes, Delve prints the values in scope and **keeps running** — like a
temporary, zero-code-change `fmt.Println`, without editing your source at all.

### Managing breakpoints

| Command | Effect |
|---|---|
| `breakpoints` (`bp`) | List every active breakpoint/tracepoint |
| `clear <id>` | Remove one specific breakpoint by its number |
| `clearall` | Remove every breakpoint |
| `condition <id> <expr>` | Attach (or update) a condition on an existing breakpoint |

## 🔍 Code Walkthrough (`main.go`)

```go
for i, n := range nums {
    if n > threshold {
        return n, i
    }
}
```

This loop is intentionally the kind of code where a plain breakpoint is annoying — you'd hit it
on `3`, `7`, and `2` before you ever reach the `9` you actually care about. A conditional
breakpoint (`n > 8`) skips straight to the iteration that matters.

## ▶️ How to Run

```bash
cd level-00-getting-started/32-breakpoints
go run main.go
dlv debug .
```

Then try the conditional breakpoint session above.

## ✅ Expected Output (normal run)

```
First value over 8: 9 at index 3

See the README for how to break ONLY on the loop iteration that matters,
instead of stepping through every single one.
```

## 🧠 Key Takeaways

- `condition <id> <expr>` turns any breakpoint into one that only fires when the expression is true.
- `trace <location>` logs a line's execution without pausing — useful for lightweight observation.
- `breakpoints`, `clear`, and `clearall` let you manage multiple breakpoints in one session.
- Conditional breakpoints are almost always faster than repeatedly pressing `continue`.

## 🛠️ Try It Yourself

1. Set a conditional breakpoint on `n > 8` as shown above, and confirm it stops exactly once,
   at `n = 9`.
2. Change the condition to `n > 100` (nothing qualifies) and confirm `continue` runs the program
   to completion without ever stopping.
3. Set a tracepoint instead of a breakpoint on the same line, and watch it log every iteration
   without pausing execution.

## ⚠️ Common Mistakes

- Writing a condition that references a variable not yet in scope at that line — Delve will
  report an error rather than silently ignoring it.
- Forgetting a breakpoint is still active after you're done with it — `clearall` before starting
  a fresh investigation avoids confusing leftover state.
