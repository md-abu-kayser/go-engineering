# 50 — Recursion

## 🎯 Learning Objectives

- Write a recursive function with a clear base case and recursive case.
- Recognize three different shapes recursion commonly takes: counting down, branching, and
  digit/structural decomposition.
- Know that Go does **not** perform tail-call optimization, and what that means in practice.

## 📖 Concept

A recursive function calls **itself**, working on a smaller version of the same problem each
time, until reaching a **base case** simple enough to answer directly without recursing further.

### The classic shape: factorial

```go
func factorial(n int) int {
    if n <= 1 {
        return 1          // base case
    }
    return n * factorial(n-1) // recursive case: smaller problem (n-1), combined with n
}
```

Every recursive function needs **both** parts: a base case that stops the recursion (without one,
or if it's unreachable, the function recurses forever, eventually crashing with a stack overflow),
and a recursive case that makes genuine progress toward that base case with each call.

### A branching shape: naive Fibonacci

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

This version calls itself **twice** per recursive case, not once — which means the same smaller
values get recomputed repeatedly (`fibonacci(5)` ends up calling `fibonacci(3)` twice,
`fibonacci(2)` three times, and so on). This is intentionally left naive here for teaching
clarity; a real implementation would typically cache previously computed results ("memoization")
or use an iterative approach instead for anything beyond small `n`.

### A structural-decomposition shape: summing digits

```go
func sumDigits(n int) int {
    if n < 10 {
        return n
    }
    return n%10 + sumDigits(n/10)
}
```

Here the "smaller problem" isn't `n-1` — it's `n` with its last digit removed (`n/10`), combined
with that removed digit (`n%10`). Recursion doesn't always mean "count down by one"; it means
"solve a smaller instance of the same shape of problem," however that reduction naturally happens
for the specific problem at hand.

### Go does NOT optimize tail calls

Some languages (Scheme, and to varying degrees others) perform **tail-call optimization**: if a
recursive call is the very last thing a function does (a "tail call"), the compiler can reuse the
current stack frame instead of pushing a new one, allowing effectively unlimited recursion depth
for tail-recursive functions. **Go's compiler does not do this** — every recursive call, tail
position or not, uses a genuine new stack frame. This means:

- A sufficiently deep recursive call chain **can** exhaust the goroutine's stack and crash the
  program (Go's goroutine stacks do grow dynamically, but they're not unlimited).
- Rewriting a naturally-recursive algorithm into "tail-recursive form," a common technique in
  languages that *do* optimize it, buys you **nothing** in Go — an iterative (`for` loop) rewrite
  is the actual fix if recursion depth is a genuine concern for your input sizes.

## 🔍 Code Walkthrough (`main.go`)

```go
return n * factorial(n-1)
```

This is **not** a tail call — after `factorial(n-1)` returns, there's still a multiplication left
to do before `factorial` itself can return. Even if Go did optimize tail calls, this particular
line wouldn't qualify; it's used here simply as the clearest possible first example of the
base-case/recursive-case shape.

## ▶️ How to Run

```bash
cd level-01-fundamentals/50-recursion
go run main.go
```

## ✅ Expected Output

```
=== Recursion ===
----------------------------------
factorial(0)  = 1
factorial(1)  = 1
factorial(5)  = 120
factorial(10)  = 3628800

fibonacci(0)  = 0
fibonacci(1)  = 1
fibonacci(5)  = 5
fibonacci(10)  = 55

sumDigits(7) = 7
sumDigits(123) = 6
sumDigits(9999) = 36

See the README for why Go does NOT optimize tail-recursive calls,
and what that means for how deep a recursive function can safely go.
```

## 🧠 Key Takeaways

- Every recursive function needs a reachable base case and a recursive case that makes genuine
  progress toward it.
- "Smaller problem" doesn't always mean "n minus one" — it means whatever natural decomposition
  fits the specific problem (digits of a number, halves of a list, and so on).
- Go performs **no** tail-call optimization — every recursive call uses a real stack frame,
  regardless of its position in the function.
- For algorithms where recursion depth could genuinely be large, prefer an iterative rewrite over
  relying on tail-call elimination that Go simply doesn't provide.

## 🛠️ Try It Yourself

1. Trace `factorial(5)` by hand, writing out each recursive call and its eventual return value,
   before checking it against the program's output.
2. Rewrite `factorial` iteratively (a `for` loop, no recursion) and confirm it produces identical
   results for every input this lesson tests.
3. Try `fibonacci(35)` or higher and notice how dramatically slower it runs compared to smaller
   inputs — direct, felt evidence of the naive version's repeated recomputation.

## ⚠️ Common Mistakes

- Writing a recursive function with an unreachable or missing base case — this causes infinite
  recursion, which in Go eventually crashes the program with a stack overflow, not an infinite
  hang.
- Assuming a "tail-recursive" rewrite will let a Go function recurse arbitrarily deep safely —
  Go's compiler doesn't grant that guarantee the way some other languages' compilers do; genuine
  unbounded depth concerns need an iterative solution instead.
