# 30 — for Loop

## 🎯 Learning Objectives

- Write the classic three-part `for` loop: `for init; condition; post { }`.
- Know the scope rule for the init variable — identical to `if`/`switch`'s init statements.
- Use `break` and `continue` correctly, and know the difference between them.

## 📖 Concept

**`for` is Go's only looping keyword.** There is no separate `while`, `do-while`, or `until` —
every kind of loop in Go is written with `for`, in one of a few different shapes. This lesson
covers the classic, fully-specified shape; [lesson 31](../31-while-style-for) and
[lesson 32](../32-infinite-loop) cover the others.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Three parts, separated by semicolons:

1. **init** (`i := 0`) — runs exactly **once**, before the loop begins.
2. **condition** (`i < 5`) — checked **before every** iteration; the loop ends the moment this
   is `false`.
3. **post** (`i++`) — runs **after every** iteration's body completes, before the condition is
   checked again.

### Scope, exactly like `if`/`switch`'s init statement

```go
for i := 0; i < 5; i++ { ... }
fmt.Println(i) // COMPILE ERROR — i is undefined out here
```

`i` is scoped to the loop only — the same rule from [lesson 25](../25-if-with-init) (`if`'s init)
and [lesson 28](../28-expression-switches) (`switch`'s init) applies here too.

### The post clause can be anything, not just `++`/`--`

```go
for i := 0; i < 10; i += 2 { ... }   // step by 2
```

Any valid statement works in the post position — `i += 2`, a function call, anything — `++`/`--`
are just the most common choice.

### `break` vs. `continue`

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue  // skip the REST of this iteration's body, go straight to the post clause
    }
    if i == 6 {
        break      // exit the loop ENTIRELY, right now
    }
    fmt.Println(i)
}
```

`continue` skips the remainder of the **current** iteration's body but keeps the loop running;
`break` ends the loop immediately, with no further iterations at all.

## 🔍 Code Walkthrough (`main.go`)

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue
    }
    if i == 6 {
        break
    }
    fmt.Printf("i = %d\n", i)
}
```

This prints `0, 1, 2, 4, 5` — `3` is skipped (via `continue`) but the loop keeps going, while
`6` stops the loop entirely (via `break`) before it or anything after it ever prints.

## ▶️ How to Run

```bash
cd level-01-fundamentals/30-for-loop
go run main.go
```

## ✅ Expected Output

```
=== for Loop ===
----------------------------------
i = 0
i = 1
i = 2
i = 3
i = 4

See the README: `i` above is NOT visible out here.

--- Counting down ---
i = 3
i = 2
i = 1

--- Stepping by 2 ---
i = 0
i = 2
i = 4
i = 6
i = 8

--- break and continue ---
i = 0
i = 1
i = 2
i = 4
i = 5
```

## 🧠 Key Takeaways

- `for` is Go's **only** looping keyword — no separate `while`/`do-while` syntax exists.
- The init variable is scoped to the loop only, matching `if`/`switch`'s init-statement rule.
- The post clause can be any statement — stepping by something other than 1 is trivial.
- `continue` skips to the next iteration; `break` exits the loop entirely, immediately.

## 🛠️ Try It Yourself

1. Rewrite the "stepping by 2" loop to step by 3 instead, and predict the printed values before
   running it.
2. Swap the order of the `continue` and `break` checks and predict how the output changes.
3. Write a loop that counts from 10 down to 1, printing only the even numbers, using `continue`.

## ⚠️ Common Mistakes

- Confusing `continue` and `break` — `continue` keeps looping; `break` stops it completely. Mixing
  them up either skips too little or stops far too early.
- Forgetting the init variable's scope ends with the loop, and trying to inspect its final value
  afterward — declare a variable **before** the loop if you need its value to persist past it.
