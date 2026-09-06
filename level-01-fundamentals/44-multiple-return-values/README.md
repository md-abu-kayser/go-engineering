# 44 — Multiple Return Values

## 🎯 Learning Objectives

- Return more than one value from a function.
- Recognize Go's two dominant multi-return idioms: `(value, error)` and `(value, ok)`.
- Use the blank identifier `_` to discard a return value you don't need.

## 📖 Concept

Go functions can return **more than one value**, declared as a comma-separated list in
parentheses:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide %d by zero", a)
    }
    return a / b, nil
}
```

### The `(value, error)` pattern

This is Go's answer to exceptions — instead of throwing, a function that can fail simply returns
an additional `error` value alongside its normal result. The caller checks it explicitly
(`if err != nil`), which [level 00's error-handling lesson]
(../../level-00-getting-started/38-reading-error-messages) covers in depth from the consuming
side; this lesson focuses on the function-declaration side of the same pattern.

### The `(value, ok)` pattern

```go
func lookup(m map[string]int, key string) (int, bool) {
    value, ok := m[key]
    return value, ok
}
```

Used when "did this succeed?" is a simple yes/no, without needing a full `error` explaining why —
map lookups (`v, ok := m[k]`) and type assertions both use this exact pattern natively; wrapping
your own functions in the same shape keeps your code consistent with the rest of the language.

### Multiple return values aren't only for errors or success flags

```go
func minMax(nums []int) (min, max int) {
    ...
    return min, max
}
```

Sometimes a function genuinely has **two** independent, useful results to hand back — there's no
rule that multiple returns must follow the `(value, error)` or `(value, ok)` shape specifically.

### Discarding a value you don't need: `_`

```go
_, err = divide(20, 4) // only care whether it succeeded, not the quotient itself
```

The blank identifier explicitly discards a return value — Go requires you to account for every
return value somehow (you can't just silently drop one by only naming fewer variables than the
function returns), so `_` is the idiomatic way to say "I know this exists, I don't need it."

## 🔍 Code Walkthrough (`main.go`)

```go
func minMax(nums []int) (min, max int) {
    min, max = nums[0], nums[0]
    for _, n := range nums[1:] {
```

This preview uses **named** return values (`min, max int` in the signature) — a full treatment of
naming return values, and the "naked return" it enables, is the entire subject of
[lesson 45](../45-named-results); here they're just used as ordinary pre-declared local variables
for convenience, with an explicit `return min, max` at the end.

## ▶️ How to Run

```bash
cd level-01-fundamentals/44-multiple-return-values
go run main.go
```

## ✅ Expected Output

```
=== Multiple Return Values ===
----------------------------------
divide(10, 2) = 5, err = <nil>
divide(10, 0) = 0, err = cannot divide 10 by zero

--- Two genuinely useful values, not an error pattern ---
minMax([5 2 8 1 9 3]) = min 1, max 9

--- The (value, ok) pattern ---
lookup(ages, "Alice") = 30, ok = true
lookup(ages, "Bob")   = 0, ok = false

--- Discarding a return value with _ ---
divide(20, 4): only checking err, ignoring the quotient -> err = <nil>
```

## 🧠 Key Takeaways

- Multiple return values are declared as a parenthesized, comma-separated list of types.
- `(value, error)` is Go's idiomatic replacement for exceptions; `(value, ok)` is for simple
  success/failure without needing a full error explanation.
- Multiple returns aren't exclusively for error handling — any function can return several
  genuinely useful results.
- `_` discards a return value you don't need — Go doesn't allow silently dropping one otherwise.

## 🛠️ Try It Yourself

1. Write a function returning three values instead of two, and call it while discarding exactly
   one of them with `_`.
2. Change `lookup` to return `(int, error)` instead of `(int, bool)`, returning a real error
   ("key not found") instead of `false` — and decide which pattern you find clearer for this case.
3. Try calling `divide` and assigning its results to only **one** variable (`result :=
   divide(10, 2)`) and read the compiler's exact error about the mismatched return-value count.

## ⚠️ Common Mistakes

- Trying to assign a multi-value return to a single variable — Go requires you to account for
  every returned value, one variable (or `_`) per value, every time.
- Reaching for `(value, error)` reflexively even when a function's failure mode is a simple,
  self-explanatory yes/no — `(value, ok)` is often clearer and is what the standard library itself
  uses for exactly that situation (map lookups, type assertions).
