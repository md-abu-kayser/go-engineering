# 10 — Bytes

## 🎯 Learning Objectives

- Know that `byte` is literally an alias for `uint8` — same type, different conventional name.
- Convert between `string` and `[]byte` cleanly, and understand the cost of doing so.
- Understand that `[]byte`, unlike `string`, is mutable.

## 📖 Concept

`byte` is Go's conventional name for **raw byte data** — but it isn't a distinct type from
`uint8` ([lesson 12](../12-unsigned-integers)); it's a genuine, built-in **alias**. Anywhere you
see `byte`, you can substitute `uint8` and the code means exactly the same thing — the alias
exists purely for readability, signaling "this represents a raw byte" rather than "this is a
small number I'm doing arithmetic on."

### Converting between `string` and `[]byte`

```go
data := []byte(greeting)   // string -> []byte: the raw UTF-8 bytes, no decoding
back := string(data)        // []byte -> string: the reverse
```

Unlike converting to `[]rune` ([lesson 09](../09-runes)), converting to `[]byte` does **no**
UTF-8 decoding — it's a direct, cheap copy of the string's existing bytes. This is why `[]byte`
is the type most of Go's I/O machinery (file reads, network reads, `bufio`, `io.Reader`/`Writer`)
actually operates on: raw bytes are exactly what comes off the wire or out of a file, before any
interpretation as text happens.

### `[]byte` is mutable; `string` is not

```go
data[0] = 'h'  // perfectly legal — []byte supports in-place modification
```

This is the single biggest practical reason to convert a string to `[]byte` in the first place:
you need to **modify** byte content in place, which a `string` (immutable,
[lesson 08](../08-strings)) simply doesn't allow. Once converted, `data` is a completely
independent copy — modifying it never affects the original string.

## 🔍 Code Walkthrough (`main.go`)

```go
data[0] = 'h'
fmt.Printf("after data[0]='h' : %s\n", string(data))
fmt.Printf("original greeting : %s\n", greeting)
```

This is the concrete proof that `[]byte(s)` produces an independent **copy**: mutating `data`
leaves `greeting` completely untouched — if `[]byte(s)` somehow shared the original string's
underlying storage, this mutation would be impossible in the first place, since strings are
immutable.

## ▶️ How to Run

```bash
cd level-01-fundamentals/10-bytes
go run main.go
```

## ✅ Expected Output

```
=== Bytes ===
----------------------------------
var b byte = 65   : 65, as a character: A
[]byte(greeting)  : 13 bytes, first few: [72 101 108 108 111]
after data[0]='h' : hello, 世界 (converted back to string)
original greeting : Hello, 世界 (completely unaffected — data is a SEPARATE copy)

string(data)      : hello, 世界
```

## 🧠 Key Takeaways

- `byte` is an alias for `uint8` — identical type, used by convention for raw byte data.
- `string` ↔ `[]byte` conversion is cheap and does no UTF-8 decoding, unlike `[]rune` conversion.
- `[]byte` is mutable; converting a string to `[]byte` gives you an independent, editable copy.
- Most of Go's I/O APIs operate on `[]byte` directly, since that's what raw data actually is.

## 🛠️ Try It Yourself

1. Print `data[:5]` before and after the `data[0] = 'h'` mutation, confirming only index `0`
   changed.
2. Try mutating the original `greeting` string directly (e.g. `greeting[0] = 'h'`) and read the
   compiler error — proof strings genuinely disallow this, unlike `[]byte`.
3. Round-trip a string through `[]byte` and back with no mutation in between, and confirm the
   result is identical to the original (`back == greeting`).

## ⚠️ Common Mistakes

- Assuming `[]byte(s)` and `[]rune(s)` do the same thing — they don't: `[]byte` gives raw UTF-8
  bytes (fast, no decoding); `[]rune` gives decoded Unicode code points (slower, but each element
  is a real character — [lesson 09](../09-runes)).
- Converting a string to `[]byte` in a hot loop unnecessarily — each conversion copies the entire
  underlying data; avoid repeated round-trips when a single conversion (or none at all) would do.
