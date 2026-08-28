# 09 — Runes

## 🎯 Learning Objectives

- Explain what a `rune` is: Go's name for a Unicode code point, an alias for `int32`.
- Correctly count and access characters in text that might contain multi-byte characters.
- Use `range` over a string to iterate character-by-character, correctly, without manual decoding.

## 📖 Concept

[Lesson 08](../08-strings) showed that `len()` and byte-indexing both operate on **bytes**, which
is wrong for non-ASCII text. `rune` is the fix: it represents one actual Unicode **code point**
("character," loosely) — Go's built-in type `rune` is simply an alias for `int32`, large enough
to hold any Unicode code point.

### Counting characters correctly

```go
len(greeting)                    // bytes — 13 for "Hello, 世界"
utf8.RuneCountInString(greeting) // actual characters — 9
```

`unicode/utf8.RuneCountInString` decodes the UTF-8 byte sequence and gives you the true character
count, regardless of how many bytes each individual character took.

### Converting to `[]rune` for indexed access

```go
runes := []rune(greeting)
runes[7] // the ACTUAL 8th character — 世
```

Converting a string to `[]rune` decodes the whole thing upfront into a slice where **each
element is one character**, regardless of its original byte width — now ordinary slice indexing
gives you correct results. The trade-off: this decodes the entire string in one pass, which costs
something for very long strings you only need to partially inspect.

### `range` over a string: decoding without converting first

```go
for i, r := range greeting {
    // i = the BYTE index this rune STARTS at (not a sequential character index!)
    // r = the actual rune (character) at that position
}
```

`range` on a `string` is special-cased to decode UTF-8 automatically, giving you each rune in
turn — without needing to convert to `[]rune` first. The index `i` you get is the **byte offset**
where that rune begins, which is why it can jump by more than 1 between iterations (e.g. `+3`
after a 3-byte character like `世`) — it's not a plain 0,1,2,3 counter.

## 🔍 Code Walkthrough (`main.go`)

```go
for i, r := range greeting {
    fmt.Printf("  byte index %2d: %c (%U)\n", i, r, r)
}
```

Watch the printed byte indices carefully: they go `0, 1, 2, 3, 4, 5, 6, 7, 10` — jumping by 3
between the last two, because `世` (starting at byte 7) occupies 3 bytes, so the next rune (`界`)
doesn't begin until byte 10. This is `range`'s UTF-8-aware iteration made directly visible.

## ▶️ How to Run

```bash
cd level-01-fundamentals/09-runes
go run main.go
```

## ✅ Expected Output

```
=== Runes ===
----------------------------------
len(greeting)               : 13 (bytes)
utf8.RuneCountInString(...) : 9 (actual characters)
[]rune(greeting)[7]          : 世 (the ACTUAL 8th character, 世)

ranging over the string directly:
  byte index  0: H (U+0048)
  byte index  1: e (U+0065)
  byte index  2: l (U+006C)
  byte index  3: l (U+006C)
  byte index  4: o (U+006F)
  byte index  5: , (U+002C)
  byte index  6:   (U+0020)
  byte index  7: 世 (U+4E16)
  byte index 10: 界 (U+754C)

'A' as a rune: 65, +1 = B
```

## 🧠 Key Takeaways

- `rune` is Go's name for a Unicode code point — an alias for `int32`.
- `utf8.RuneCountInString` gives the true character count; `len()` gives the byte count.
- `[]rune(s)` decodes a string into indexable, correctly-sized characters.
- `range` over a `string` decodes UTF-8 automatically, yielding each rune with its starting byte
  index — which is not a simple sequential counter once multi-byte characters appear.

## 🛠️ Try It Yourself

1. Print `runes[8]` (not `runes[7]`) and confirm it's `界`, the character right after `世`.
2. Write a loop using `range` that builds a `[]rune` manually by appending each `r` — and confirm
   the result equals `[]rune(greeting)` built directly.
3. Try `string(rune(0x1F600))` (an emoji code point) and confirm it correctly produces a
   multi-byte UTF-8 string when printed.

## ⚠️ Common Mistakes

- Using a plain `for i := 0; i < len(s); i++` loop with `s[i]` to "iterate characters" — this
  iterates **bytes**, silently breaking on any non-ASCII input; use `range` or `[]rune` instead.
- Treating a `range` string loop's index as a sequential character counter (`0, 1, 2, ...`) — it's
  a **byte offset**, and will skip values whenever a multi-byte character appears earlier in the
  string.
