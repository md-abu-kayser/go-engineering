# 49 — Closures

## 🎯 Learning Objectives

- Write a closure: a function literal that captures a variable from its surrounding scope.
- Use the classic "counter generator" pattern, and confirm each generated closure has independent
  state.
- Confirm Go's modern (1.22+) per-iteration loop variable behavior directly — a real, historically
  significant change specifically relevant to closures created inside loops.

## 📖 Concept

[Lesson 48](../48-anonymous-functions) covered the *syntax* of a function literal. A **closure**
is a function literal that references a variable from **outside** its own body — it "closes
over" that variable, keeping a genuine reference to it (not a copy) for as long as the closure
itself exists, even after the function that originally declared the variable has returned.

### The classic example: a counter generator

```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

Normally, `count` would be a local variable that disappears once `makeCounter` returns. But
because the returned function literal **references** `count`, Go keeps it alive for as long as
that specific function value exists — this is exactly what makes it a closure rather than a plain
anonymous function. Each **call** to `makeCounter()` creates a brand-new `count`, so:

```go
counterA := makeCounter()
counterB := makeCounter()
counterA() // 1
counterA() // 2
counterB() // 1 — completely independent of counterA
```

### Closures modify the ORIGINAL variable, not a copy

```go
total := 0
addToTotal := func(n int) {
    total += n // this is the SAME total declared above, not a copy
}
```

Unlike a function *parameter* ([lesson 43](../43-parameters), always pass-by-value), a closure's
captured variable is a genuine reference to the original — changes made through the closure are
visible in the outer scope too, and vice versa.

### Go 1.22+: loop variables are per-iteration

This matters specifically for closures created **inside a loop**:

```go
var funcs []func()
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() {
        fmt.Println(i)
    })
}
for _, f := range funcs {
    f()
}
// Go 1.22+: prints 0, 1, 2 — each closure captured ITS OWN iteration's i
```

**Historical note:** before Go 1.22, `for` loops used a **single, shared** `i` variable across
every iteration — so all three closures above would have captured the *same* variable, and by the
time they ran (after the loop finished), it would hold its final value, `3`, for every single one
of them (printing `3, 3, 3` instead). This was a famous, extremely common source of bugs. Go 1.22
changed the language specification so each iteration gets its **own** copy of the loop variable,
eliminating this entire category of mistake — the behavior demonstrated in this lesson (`0, 1, 2`)
is the modern, current behavior, verified directly against this repository's Go 1.22 toolchain.

## 🔍 Code Walkthrough (`main.go`)

```go
counterA := makeCounter()
counterB := makeCounter()
```

Calling `makeCounter` **twice** is the key detail — each call executes `count := 0` fresh,
creating a genuinely separate variable each time, which is exactly why `counterA` and `counterB`
don't interfere with each other despite being built from the identical function.

## ▶️ How to Run

```bash
cd level-01-fundamentals/49-closures
go run main.go
```

## ✅ Expected Output

```
=== Closures ===
----------------------------------
counterA(): 1
counterA(): 2
counterA(): 3
counterB(): 1 (independent of counterA — its OWN count)

--- Closing over an outer variable directly ---
total = 15 (modified via the closure, twice)

--- Closures inside a loop (Go 1.22+ per-iteration variables) ---
  captured i = 0
  captured i = 1
  captured i = 2
```

## 🧠 Key Takeaways

- A closure is a function literal that captures a variable from its enclosing scope by reference,
  not by value.
- Each call to a function that returns a closure creates fresh, independent captured state.
- A closure can both read and modify its captured variable — the outer scope sees those changes too.
- Since Go 1.22, each loop iteration gets its own copy of the loop variable — closures created
  inside a loop correctly capture that iteration's specific value, fixing a long-standing,
  well-known historical gotcha.

## 🛠️ Try It Yourself

1. Call `makeCounter()` a third time and confirm the new counter starts from `1`, completely
   unaffected by `counterA`'s or `counterB`'s current counts.
2. Write a closure-returning function that takes a starting value as a parameter
   (`makeCounterFrom(start int) func() int`), instead of always starting at `0`.
3. If you have access to an older Go toolchain (or can research it), compare this lesson's loop
   example's output against what it would have printed before Go 1.22 — `3, 3, 3` instead of
   `0, 1, 2`.

## ⚠️ Common Mistakes

- Assuming a closure captures a variable's **value at the time of capture** — it doesn't; it
  captures a live **reference**, which is exactly why the loop-variable behavior above was such a
  well-known trap in versions of Go before 1.22.
- Reading old Go code, blog posts, or tutorials that manually work around the pre-1.22
  loop-variable-sharing behavior (commonly by adding `i := i` inside the loop body) and assuming
  that workaround is still necessary — on Go 1.22+, it no longer is, though it remains harmless if
  present.
