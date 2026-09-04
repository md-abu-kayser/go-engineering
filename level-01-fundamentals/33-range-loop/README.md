# 33 — range Loop

## 🎯 Learning Objectives

- Use `for range` over a slice, an array, a string, and a map.
- Know exactly what `range` yields for each: `(index, value)`, `(byte index, rune)`, or
  `(key, value)`.
- Know that map iteration order is deliberately **not guaranteed** — and why Go makes that choice.

## 📖 Concept

`range` is Go's iteration keyword, used after `for`, that walks a collection and yields its
elements — the specific shape of what it yields depends on what you're ranging over.

### Slices and arrays: `(index, value)`

```go
for i, fruit := range fruits {
    // i is the index, fruit is the value at that index
}
```

Both work identically for this purpose — arrays and slices range the same way.

### Discarding what you don't need

```go
for i := range fruits { }         // index only
for _, fruit := range fruits { }   // value only, index explicitly discarded
```

Both forms are common and idiomatic — use whichever half of the pair your logic actually needs.

### Strings: `(byte index, rune)`

```go
for i, r := range "Hi, 世" {
    // i is the BYTE index where this rune starts; r is the decoded rune
}
```

This is exactly [lesson 09](../09-runes)'s UTF-8-aware iteration — worth remembering here as one
more case of "what does `range` yield," alongside slices and maps.

### Maps: `(key, value)` — in NO guaranteed order

```go
for name, age := range ages {
    // order across different runs of the SAME program is NOT guaranteed to be the same
}
```

This is deliberate: Go's runtime intentionally **randomizes** map iteration order, specifically
so no one accidentally writes code that depends on an ordering the language never actually
promised. If you need a specific order, sort the map's keys explicitly first.

## 🔍 Code Walkthrough (`main.go` and the test file)

```go
for range s {
    count++
}
```

In the test file's `countRunes`, `range` is used with **no** variables captured at all (not even
`_`) — this is valid when you only care that an iteration happened, not its index or value; here,
it's simply used to count how many runes a string decodes into.

```go
{"nil slice", nil, 0},
```

`sumSlice`'s test includes a `nil` slice case specifically — `range` over a `nil` slice is
completely safe and simply doesn't iterate at all (zero times), matching
[lesson 03](../03-zero-values)'s coverage of `nil` slices being safely usable for reading.

## ▶️ How to Run

```bash
cd level-01-fundamentals/33-range-loop
go run main.go
go test -v ./...
```

## ✅ Expected Output

```
=== range Loop ===
----------------------------------
--- range over a slice ---
  [0] = apple
  [1] = banana
  [2] = cherry

--- range with index only ---
  index 0
  index 1
  index 2

--- range with value only (index discarded) ---
  apple
  banana
  cherry

--- range over an array ---
  [0] = 10
  [1] = 20
  [2] = 30

--- range over a string (decodes runes) ---
  byte 0: H
  byte 1: i
  byte 2: ,
  byte 3:  
  byte 4: 世

--- range over a map (order NOT guaranteed) ---
  Alice is 30
  Bob is 25
  (2 entries total, in unspecified order)
```

(The map section's line order may differ between runs — that's expected, not a bug.)

## 🧠 Key Takeaways

- `range` over a slice/array yields `(index, value)`; discard either half with `_` or by omitting it.
- `range` over a string yields `(byte index, rune)`, decoding UTF-8 automatically.
- `range` over a map yields `(key, value)`, in **deliberately unspecified** order.
- `range` over a `nil` slice or map is completely safe — it simply doesn't iterate.

## 🛠️ Try It Yourself

1. Run `go run main.go` several times in a row and confirm the map section's line order can (but
   might not always) differ between runs.
2. Sort `ages`' keys explicitly (using `sort.Strings` on a `[]string` of the keys) before ranging,
   to get a deterministic, repeatable order.
3. Run `go test -v ./...` and confirm every subtest passes, including the `nil` slice case.

## ⚠️ Common Mistakes

- Relying on map iteration order being consistent — it never was guaranteed, and Go actively
  randomizes it to prevent exactly this assumption from creeping into code.
- Forgetting `range` over a string yields **byte indices**, not sequential character positions —
  see [lesson 09](../09-runes) if this is a surprise.
