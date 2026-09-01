# 31 — While-Style for

## 🎯 Learning Objectives

- Write a `for` loop with only a condition — Go's equivalent of `while`.
- Understand that Go deliberately has **no** separate `while` keyword; this is just another shape
  of the one `for` keyword from [lesson 30](../30-for-loop).
- Recognize when this shape is the natural choice over the full three-part form.

## 📖 Concept

[Lesson 30](../30-for-loop) covered the full `for init; condition; post { }` form. Dropping the
`init` and `post` parts (and both semicolons) leaves just a condition:

```go
for n < 20 {
    fmt.Println(n)
    n *= 2
}
```

This **is** Go's `while` loop — there's no separate keyword, just a `for` with only its middle
part present. Every language feature that a `while` loop needs (loop while a condition holds,
`break`/`continue` still work exactly the same) is already covered by this shape of `for`.

### When this shape is the natural choice

Reach for this form when you don't have a natural "count from A to B" structure — instead, you
loop until some **condition**, often based on a value that changes in a way that isn't a simple
counter:

```go
value := 100
for value > 1 {
    value /= 2
}
```

Here there's no meaningful "loop variable running from 0 to N" — the loop continues based on
`value`'s own changing state, which is exactly the situation a `while`-style loop fits naturally,
compared to forcing it into a three-part `for` with an unused or awkward post clause.

## 🔍 Code Walkthrough (`main.go`)

```go
n := 1
for n < 20 {
    fmt.Printf("n = %d\n", n)
    n *= 2
}
```

`n` is declared **before** the loop this time (with `:=`), not inside a `for` init clause —
unlike [lesson 30](../30-for-loop)'s form, there's no init slot in this shape to declare it in,
so it needs to exist in the enclosing scope already. This also means, unlike lesson 30's loop
variable, `n` **is** still visible and usable after this loop ends.

## ▶️ How to Run

```bash
cd level-01-fundamentals/31-while-style-for
go run main.go
```

## ✅ Expected Output

```
=== While-Style for ===
----------------------------------
n = 1
n = 2
n = 4
n = 8
n = 16

--- A more realistic example ---
Halved 100 down to 1 in 6 steps.

Other languages: while (n < 20) { ... }
Go:              for n < 20 { ... }        <- same idea, no separate keyword
```

## 🧠 Key Takeaways

- `for condition { }` (no init, no post) is Go's `while` loop — there's no separate keyword.
- The loop variable, if any, must be declared **before** this shape of loop, since there's no init
  slot — and it remains in scope after the loop ends, unlike the three-part form's init variable.
- Use this shape when looping depends on a condition/changing state, not a simple counted range.

## 🛠️ Try It Yourself

1. Rewrite the doubling loop (`n *= 2`) to instead loop while `n` is even, dividing by 2 each
   time, and print how many times it took to reach an odd number.
2. Confirm `value` (from the halving example) is still accessible and printable **after** its
   loop ends — unlike `i` in lesson 30's three-part form.
3. Rewrite lesson 30's simple counting loop (`for i := 0; i < 5; i++`) using this while-style
   shape instead (declaring `i` beforehand), and compare which reads more naturally for a simple
   counted loop.

## ⚠️ Common Mistakes

- Forgetting to declare the loop's controlling variable **before** the loop — there's no init
  slot in this shape to do it inline, unlike lesson 30's form.
- Forgetting to actually update the condition's variable somewhere inside the loop body — an easy
  way to accidentally write an infinite loop (see [lesson 32](../32-infinite-loop) for when that's
  intentional).
