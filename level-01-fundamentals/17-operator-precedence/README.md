# 17 — Operator Precedence

## 🎯 Learning Objectives

- Recall Go's operator precedence ordering, from highest to lowest.
- Predict how a mixed-operator expression actually evaluates, without running it.
- Know when to add parentheses for human readability, even where the compiler doesn't need them.

## 📖 Concept

Go evaluates expressions according to a fixed **precedence** table — higher-precedence operators
bind more tightly, meaning they're effectively grouped first.

### The precedence levels (highest to lowest)

| Precedence | Operators |
|---|---|
| 5 (highest) | `*` `/` `%` `<<` `>>` `&` `&^` |
| 4 | `+` `-` `\|` `^` |
| 3 | `==` `!=` `<` `<=` `>` `>=` |
| 2 | `&&` |
| 1 (lowest) | `\|\|` |

A few things worth internalizing directly from this table:

- **Arithmetic binds tighter than comparison.** `1 + 2 == 3` is `(1 + 2) == 3`, not
  `1 + (2 == 3)` (which wouldn't even compile — you can't add an `int` to a `bool`).
- **`&&` binds tighter than `||`.** `a || b && c` is `a || (b && c)`, not `(a || b) && c` — these
  can genuinely produce different results depending on the values involved.
- **Bitwise operators are split across two precedence levels**, interleaved with arithmetic:
  `&` and `&^` are as tight as `*`/`/`; `|` and `^` are as loose as `+`/`-`. This means
  `1 | 2 & 3` is `1 | (2 & 3)`, not `(1 | 2) & 3` — easy to get wrong if you assume all bitwise
  operators share one precedence level.

### Parentheses: use them for humans, not just the compiler

Every expression in this lesson is **fully unambiguous** to the Go compiler without any
parentheses at all. But "technically unambiguous to a parser" and "instantly readable to a human
skimming the code" are different bars. Adding parentheses that don't change behavior, purely to
make grouping obvious at a glance, is a low-cost, high-value habit — especially once `&&`/`||` or
bitwise operators are mixed with comparisons in the same expression.

## 🔍 Code Walkthrough (`main.go`)

```go
result3 := true || false && false
```

Read casually, left to right, this might look like `(true || false) && false`, which would be
`false`. But because `&&` binds tighter, it's actually `true || (false && false)`, which is
`true` — the two readings genuinely disagree, which is exactly why this example is included
rather than a case where both readings happen to agree.

## ▶️ How to Run

```bash
cd level-01-fundamentals/17-operator-precedence
go run main.go
```

## ✅ Expected Output

```
=== Operator Precedence ===
----------------------------------
2 + 3 * 4        = 14 (NOT 20 — * happens first)
1 + 2 == 3       = true (arithmetic happens BEFORE the comparison)
true || false && false = true (&& binds tighter: true || (false && false))
1 | 2 & 3        = 3 (& binds tighter: 1 | (2 & 3) = 1 | 2 = 3)

--- When parentheses are technically unnecessary but STILL a good idea ---
without parens: true, with parens: true (same result, but which was easier to read?)
```

## 🧠 Key Takeaways

- Go's precedence order, high to low: `* / % << >> & &^`, then `+ - | ^`, then comparisons, then
  `&&`, then `||`.
- Arithmetic always happens before comparison, and comparison always happens before `&&`/`||`.
- `&&` binds tighter than `||` — genuinely change results if you mentally swap them.
- Bitwise `&`/`&^` are as tight as multiplication; `|`/`^` are as loose as addition — not all
  bitwise operators share one precedence level.

## 🛠️ Try It Yourself

1. Predict the result of `2 + 3 * 4 > 10 && false` on paper before running it, then check your
   work.
2. Rewrite `1 | 2 & 3` with explicit parentheses matching its ACTUAL evaluation order, and
   confirm the result is unchanged.
3. Find one expression elsewhere in this repository (any lesson) that relies on operator
   precedence, and add clarifying parentheses to it as an exercise — without changing its
   behavior.

## ⚠️ Common Mistakes

- Assuming `&&` and `||` share one precedence level (as in some other languages' simplified mental
  models) — Go's `&&` genuinely binds tighter, which can flip a result if you get it backwards.
- Assuming all bitwise operators (`&`, `|`, `^`, `&^`) sit at one shared precedence level — `&`
  and `&^` are tighter than `|` and `^`, matching how tightly they bind relative to `*` vs `+`.
