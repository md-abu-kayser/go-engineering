# 08 — Strings

## 🎯 Learning Objectives

- Explain that Go strings are UTF-8 encoded byte sequences, and that `len()` counts **bytes**.
- Explain string immutability, and what "modifying" a string actually means under the hood.
- Use the most common `strings` package helpers.

## 📖 Concept

A Go `string` is, under the hood, an **immutable sequence of bytes**, conventionally interpreted
as UTF-8-encoded text. Two consequences of this follow directly, and both trip up beginners.

### `len()` counts bytes, not "characters"

```go
greeting := "Hello, 世界"
len(greeting) // 13, not 9
```

In UTF-8, ASCII characters (`H`, `e`, `l`, …) take **1 byte** each, but `世` and `界` each take
**3 bytes**. `len()` on a string always reports the total **byte count** — for text that's
entirely ASCII this happens to match the character count, which is exactly what makes the
distinction easy to overlook until it silently breaks on non-ASCII input.

### Indexing gives you a byte, not a "character"

```go
greeting[0] // a single byte (uint8), NOT necessarily one visible character
```

For ASCII text, indexing happens to look correct, since each character *is* one byte. For
multi-byte characters, indexing gives you one raw byte of a larger sequence — almost never what
you actually want. [Lesson 09](../09-runes) covers the **correct** way to work with individual
characters: converting to `[]rune` or ranging over the string.

### Strings are immutable

```go
upper := strings.ToUpper(greeting)
```

There is no way to change a byte of an existing string in place — every string "transformation"
function in Go (`strings.ToUpper`, concatenation with `+`, etc.) returns a **brand new** string,
leaving the original completely untouched. This is why `greeting` is unaffected after
`strings.ToUpper(greeting)` — the result is assigned to a new variable, `upper`.

### Common `strings` package functions

```go
strings.Contains(s, substr)  // does s contain substr?
strings.Split(s, sep)         // split into a []string
strings.TrimSpace(s)           // remove leading/trailing whitespace
strings.ToUpper(s) / ToLower(s)
strings.Join(slice, sep)        // the inverse of Split
```

## 🔍 Code Walkthrough (`main.go`)

```go
fmt.Printf("greeting[0]       : %d (%q) — the BYTE at index 0\n", greeting[0], string(greeting[0]))
```

`greeting[0]` is `72` (the ASCII value of `'H'`) — this happens to be correct here because `H` is
a single-byte ASCII character right at the start. Try indexing further into the string, past
`世` or `界`, and this same technique would give you a meaningless partial byte instead.

## ▶️ How to Run

```bash
cd level-01-fundamentals/08-strings
go run main.go
```

## ✅ Expected Output

```
=== Strings ===
----------------------------------
greeting          : Hello, 世界
len(greeting)     : 13 bytes
greeting[0]       : 72 ("H") — the BYTE at index 0
strings.ToUpper() : HELLO, 世界 (a NEW string — greeting itself is unchanged)
greeting (after)  : Hello, 世界 (still the original)
greeting + "!"    : Hello, 世界!
strings.Contains  : true
strings.Split     : ["a" "b" "c"]
strings.TrimSpace : "padded"
```

## 🧠 Key Takeaways

- `len(s)` returns the byte count, which only equals the character count for pure ASCII text.
- `s[i]` gives you a single byte, not necessarily one full character.
- Strings are immutable — every transformation returns a new string; the original is untouched.
- The `strings` package covers the vast majority of everyday text manipulation needs.

## 🛠️ Try It Yourself

1. Print `len("世界")` alone and confirm it's `6` (two 3-byte characters), not `2`.
2. Try indexing into `"世界"` at index `1` (not `0`) and print the raw byte value — notice it's not
   a valid, meaningful character on its own.
3. Use `strings.Join` to reverse what `strings.Split("a,b,c", ",")` produced, back into `"a,b,c"`.

## ⚠️ Common Mistakes

- Using `len(s)` to mean "number of characters" for text that might contain non-ASCII content —
  use `utf8.RuneCountInString(s)` (or convert to `[]rune`, [lesson 09](../09-runes)) instead.
- Indexing a string expecting a "character" and getting a meaningless byte value for non-ASCII
  input — range over the string ([lesson 09](../09-runes)) when you need actual characters.
