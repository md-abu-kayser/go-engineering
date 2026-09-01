# 32 — Infinite Loop

## 🎯 Learning Objectives

- Write `for {}` — Go's explicit infinite loop shape.
- Use `break` correctly to exit one, and understand this is the loop's **only** way to stop.
- Recognize this as the standard pattern behind real-world event/server/polling loops.

## 📖 Concept

Dropping **everything** from `for` — no init, no condition, no post — gives you an infinite loop:

```go
for {
    // runs forever, unless something inside explicitly stops it
}
```

This is a completely deliberate, idiomatic Go construct — not a mistake or a bug pattern. It's
the natural shape for "keep doing this until some internal event says to stop," which doesn't
always map cleanly onto a simple counted range or condition-at-the-top loop.

### `break` is the way out

```go
count := 0
for {
    count++
    if count >= 5 {
        break
    }
}
```

Without a `break` (or a `return`, which also exits the enclosing function and therefore the
loop) somewhere inside, `for {}` genuinely never ends — that's the entire point of the shape, and
it's your responsibility to ensure some path through the loop body actually reaches a `break`.

### A realistic pattern: work until there's no more work

```go
for {
    if len(queue) == 0 {
        break
    }
    item := queue[0]
    queue = queue[1:]
    // process item
}
```

Notice there's no natural "count" here — the loop continues based on whether there's still work
left, checked **inside** the loop body itself, rather than in a condition Go checks automatically
before each iteration. This is precisely the situation `for {}` with an internal `break` fits
better than either of the other two `for` shapes.

### Where this shows up in real programs

This exact pattern — `for {}` with a `break` (or sometimes no `break` at all, running genuinely
forever) — is how you'll see, almost universally:

- **Server accept loops**: `for { conn := listener.Accept(); go handle(conn) }` — accepting
  connections forever, for the life of the server process.
- **Event loops**: repeatedly reading from a channel or queue until told to stop.
- **Polling loops**: checking some external condition on an interval, until it's satisfied.

## 🔍 Code Walkthrough (`main.go`)

```go
for {
    if len(queue) == 0 {
        fmt.Println("queue is empty, stopping")
        break
    }
    item := queue[0]
    queue = queue[1:]
    fmt.Printf("processing %s ...\n", item)
}
```

The condition that eventually stops this loop (`len(queue) == 0`) is checked **inside** the loop
body, not automatically by `for` itself — this is the defining characteristic of this shape versus
[lesson 31](../31-while-style-for)'s while-style `for condition {}`, where the condition check is
built into the loop's own syntax.

## ▶️ How to Run

```bash
cd level-01-fundamentals/32-infinite-loop
go run main.go
```

## ✅ Expected Output

```
=== Infinite Loop ===
----------------------------------
iteration 1
iteration 2
iteration 3
iteration 4
iteration 5

Loop exited cleanly via break.

--- Realistic pattern: process until empty ---
processing task-1 (queue has 2 item(s) left)
processing task-2 (queue has 1 item(s) left)
processing task-3 (queue has 0 item(s) left)
queue is empty, stopping

This exact shape — for {} with an internal break — is how event loops,
server accept loops, and polling loops are almost always written in Go.
```

## 🧠 Key Takeaways

- `for {}` (nothing after `for`) loops forever, by explicit design.
- `break` (or `return`) is the only way out — without one reachable, the loop genuinely never ends.
- This shape fits naturally when the stopping condition is checked **inside** the loop body,
  rather than automatically before each iteration.
- Server accept loops, event loops, and polling loops are almost universally written this way.

## 🛠️ Try It Yourself

1. Remove the `break` from the first loop (in a scratch copy) and — carefully — consider what
   would happen if you ran it; don't actually run an unbounded loop that prints forever.
2. Rewrite the queue-processing loop using [lesson 31](../31-while-style-for)'s while-style form
   instead (`for len(queue) > 0 { ... }`), and compare the two versions.
3. Add a `continue` inside the queue loop to skip processing any item equal to `"task-2"`, while
   still eventually reaching the empty-queue `break`.

## ⚠️ Common Mistakes

- Writing `for {}` without any reachable `break`/`return` by accident — genuinely infinite, and
  (depending on the loop body) can consume 100% of a CPU core or hang the program entirely.
- Reaching for `for {}` when a condition-based `for condition {}` ([lesson 31](../31-while-style-for))
  would express the same logic more directly — prefer the simplest shape that clearly expresses
  the loop's actual stopping rule.
