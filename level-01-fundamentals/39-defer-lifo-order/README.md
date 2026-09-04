# 39 — defer LIFO Order

## 🎯 Learning Objectives

- Confirm directly that multiple `defer` statements run in LIFO (Last In, First Out) order.
- Explain why LIFO is exactly the correct order for nested resource cleanup.
- Predict the output of a function with several `defer` statements, before running it.

## 📖 Concept

[Lesson 38](../38-defer-basics) introduced `defer` for a single deferred call. This lesson
focuses specifically on what happens with **multiple** defers in the same function: they run in
**LIFO order** — Last In, First Out, exactly like a stack.

```go
defer fmt.Println("call #1")
defer fmt.Println("call #2")
defer fmt.Println("call #3")
// when the function returns, output is: call #3, call #2, call #1
```

Think of each `defer` as **pushing** onto a stack; when the function is about to return, that
stack is **popped**, one call at a time, from the top — meaning the **most recently** deferred
call runs **first**.

### Why LIFO order is exactly right for resource cleanup

This ordering isn't arbitrary — it matches how resource acquisition and cleanup naturally nest in
real code:

```go
// open resource A
defer closeA()

// (using A) open resource B, which DEPENDS on A being open
defer closeB()

// ... use both A and B ...

// when this function returns: closeB() runs FIRST, then closeA()
```

If `B` was opened using `A` (e.g. B is a connection built on top of A), it makes sense for `B` to
be closed **before** `A` — cleaning up in the reverse order of acquisition is exactly what you
want, and that's precisely what LIFO ordering gives you automatically, with zero extra bookkeeping.

## 🔍 Code Walkthrough (`main.go`)

```go
defer fmt.Println("deferred call #1 (registered FIRST, runs LAST)")
defer fmt.Println("deferred call #2 (registered SECOND, runs SECOND-to-last)")
defer fmt.Println("deferred call #3 (registered LAST, runs FIRST)")
```

Each message describes its own position in **both** orderings (registration order and actual run
order) directly in its text — so the printed output itself demonstrates the LIFO rule without
needing to cross-reference anything else.

## ▶️ How to Run

```bash
cd level-01-fundamentals/39-defer-lifo-order
go run main.go
```

## ✅ Expected Output

```
=== defer LIFO Order ===
----------------------------------
Registering three deferred calls, in this order: 1, 2, 3...
...main() body continues normally here...
...and now main() is about to return; watch the order below:
deferred call #3 (registered LAST, runs FIRST)
deferred call #2 (registered SECOND, runs SECOND-to-last)
deferred call #1 (registered FIRST, runs LAST)
```

## 🧠 Key Takeaways

- Multiple `defer` statements run in **LIFO** order — last registered, first executed.
- This matches a stack's push/pop behavior exactly, one call popped per return.
- LIFO ordering is precisely correct for nested resource cleanup: close things in the reverse
  order you opened them, automatically, with no manual tracking needed.
- Combined with [lesson 38](../38-defer-basics)'s "arguments evaluated immediately" rule, you can
  now fully predict the output of any function using multiple defers.

## 🛠️ Try It Yourself

1. Add a fourth `defer` statement and predict exactly where in the output it will appear before
   running the program.
2. Write a small function simulating opening and closing two nested resources (just printing
   "opening X" / "closing X"), using `defer` for each close, and confirm the closes happen in
   reverse order.
3. Combine this lesson with [lesson 38](../38-defer-basics)'s argument-evaluation-timing rule:
   write a loop that `defer`s a `Printf` referencing the loop variable on each iteration, and
   predict what values actually get printed (this is also closely related to the historical
   Go loop-variable-capture behavior worth reading about separately).

## ⚠️ Common Mistakes

- Assuming deferred calls run in the order they were **written** (FIFO) — it's the reverse (LIFO);
  this is easy to get backwards when predicting output for the first few times.
- Registering cleanup defers in the wrong order relative to acquisition (deferring the "outer"
  resource's cleanup before the "inner" one that depends on it) — since LIFO always reverses
  registration order, defer cleanup calls in the **same** order you acquire the resources, and
  let LIFO handle reversing them correctly for you.
