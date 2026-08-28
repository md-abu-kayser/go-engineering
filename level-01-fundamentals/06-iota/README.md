# 06 — iota

## 🎯 Learning Objectives

- Use `iota` to generate a sequence of related constants without numbering each one by hand.
- Skip a value in an `iota` sequence deliberately.
- Recognize and use the classic bit-shifted-flags pattern built on `iota`.

## 📖 Concept

`iota` is a special identifier usable only inside a `const (...)` block. It starts at `0` and
increments by `1` for every new line within that block — Go's built-in tool for numbered
enumerations.

### The basic pattern

```go
const (
    Sunday weekday = iota // 0
    Monday                 // 1
    Tuesday                 // 2
    ...
)
```

Only the first line needs `= iota` written explicitly — every subsequent constant in the block
**repeats the same expression** from the first line, with `iota` incrementing each time. This is
exactly why `Monday` doesn't need to repeat `weekday = iota` itself.

### Skipping a value with `_`

```go
const (
    Bronze = iota + 1  // 1
    Silver               // 2
    _                     // 3 — consumed, but not named
    Platinum             // 4
)
```

The blank identifier `_` still consumes an `iota` value (advancing the counter), it just doesn't
create a usable named constant for it — useful when a numeric protocol or format has a reserved
or deliberately-unused value you still need to "skip past."

### The classic pattern: bit flags

```go
const (
    PermRead    permission = 1 << iota // 1 << 0 = 1   (binary 001)
    PermWrite                            // 1 << 1 = 2   (binary 010)
    PermExecute                          // 1 << 2 = 4   (binary 100)
)
```

Each constant gets a **distinct bit**, rather than a distinct sequential number. This matters
because distinct bits can be **combined** with `|` (bitwise OR) into one value representing
multiple flags simultaneously, and **tested** with `&` (bitwise AND):

```go
readWrite := PermRead | PermWrite     // 1 | 2 = 3 (binary 011 — both bits set)
hasWrite := readWrite&PermWrite != 0   // true — the PermWrite bit is set in readWrite
```

This is exactly the pattern behind flag-style APIs throughout the Go standard library (e.g.
`os.O_RDONLY | os.O_CREATE` when opening a file).

## 🔍 Code Walkthrough (`main.go`)

```go
const (
    Bronze = iota + 1
    Silver
    _
    Platinum
)
```

`iota` **resets to 0** at the start of every new `const` block — this is a completely separate
sequence from the `weekday` block above it, which is why `Bronze` starts at `1` (via `iota + 1`
with `iota == 0`), not continuing from wherever the previous block left off.

## ▶️ How to Run

```bash
cd level-01-fundamentals/06-iota
go run main.go
```

## ✅ Expected Output

```
=== iota ===
----------------------------------
Sunday=0 Monday=1 ... Saturday=6
Bronze=1 Silver=2 Platinum=4 (value 3 skipped)
PermRead=1 PermWrite=2 PermExecute=4
readWrite (Read|Write) = 3
readWrite has PermWrite?   true
readWrite has PermExecute? false
```

## 🧠 Key Takeaways

- `iota` starts at `0` and increments by one per line, and only within a `const` block.
- Every constant after the first repeats that first line's expression, with `iota` substituted in.
- `iota` resets to `0` at the start of each new `const (...)` block — blocks don't share a counter.
- `_` skips a value without naming it; `1 << iota` produces the classic bit-flag pattern.

## 🛠️ Try It Yourself

1. Add an eighth constant after `Saturday` with no explicit value and confirm it continues the
   sequence correctly (it should be `7`).
2. Add a fourth permission flag, `PermDelete`, continuing the `1 << iota` pattern, and confirm it
   gets the value `8`.
3. Build a combined value with three flags OR'd together and test for all three with `&`.

## ⚠️ Common Mistakes

- Assuming `iota` continues counting across **separate** `const` blocks — it resets to `0` every
  time a new `const (...)` block begins.
- Forgetting that `_` still **advances** `iota`, even though it doesn't create a named constant —
  the constant after a skipped `_` is not off-by-one from what you might expect.
