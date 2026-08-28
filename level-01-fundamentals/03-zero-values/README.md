# 03 — Zero Values

## 🎯 Learning Objectives

- State the zero value for every commonly-used Go type from memory.
- Explain why Go guarantees this, and how it differs from "uninitialized memory" in other languages.
- Know which zero values are directly usable (`nil` slices, structs) versus which need explicit
  initialization before use (`nil` maps for writing).

## 📖 Concept

In many languages, a declared-but-not-explicitly-set variable holds **whatever garbage happened
to be in that memory location** — genuinely undefined, unpredictable behavior. Go makes a
different, deliberate guarantee: **every** variable, when declared without an initializer, is
automatically set to its type's well-defined **zero value**. There is no such thing as
"uninitialized" in Go — only "explicitly given a value" or "zero value."

### The zero value table

| Type | Zero value |
|---|---|
| Numeric (`int`, `float64`, etc.) | `0` (or `0.0`) |
| `bool` | `false` |
| `string` | `""` (empty, not `nil` — a string's zero value is a valid, usable empty string) |
| `rune` | `0` (it's really just `int32` underneath) |
| `byte` | `0` (it's really just `uint8` underneath) |
| Slice | `nil` |
| Map | `nil` |
| Pointer | `nil` |
| Channel | `nil` |
| Function | `nil` |
| Interface | `nil` |
| Struct | every field gets **its own** zero value, recursively |

### An important nuance: `nil` isn't always "unusable"

A `nil` **slice** behaves like an empty slice for reading (`len(nilSlice) == 0`, ranging over it
does nothing) — it's genuinely safe to use without initializing it first, for read-only purposes,
and even `append`ing to it works fine (it allocates on first use).

A `nil` **map**, by contrast, is safe to **read** from (`m["key"]` returns the value type's zero
value, no panic) but **panics** if you try to **write** to it (`m["key"] = 1`) — maps must be
initialized with `make` or a map literal before you can write to them. This asymmetry between
slices and maps trips up a lot of newcomers.

## 🔍 Code Walkthrough (`main.go`)

```go
var slice []int
var m map[string]int
```

Both are declared with no initializer, so both are `nil` — but printing `slice` shows `[]`
(an empty-looking slice), and `len(slice)` correctly reports `0`, all without ever having called
`make`. This is exactly the "nil slice is safely usable" behavior described above.

```go
type point struct{ X, Y int }
var pt point
```

A zero-value **struct** isn't some special "empty" marker — each field independently gets its
own zero value (`X: 0, Y: 0` here), recursively, however deeply nested the struct is.

## ▶️ How to Run

```bash
cd level-01-fundamentals/03-zero-values
go run main.go
```

## ✅ Expected Output

```
=== Zero Values ===
----------------------------------
int              : 0
float64          : 0
bool             : false
string           : "" (empty, not nil)
rune             : 0
byte             : 0
[]int (slice)    : [] (nil: true, len: 0)
map[string]int   : map[] (nil: true)
*int (pointer)   : <nil> (nil: true)
chan int         : <nil> (nil: true)
func()           : nil: true
interface{}      : <nil> (nil: true)
struct{X,Y int}  : {X:0 Y:0} (fields get THEIR zero values too)
```

## 🧠 Key Takeaways

- Go guarantees every variable a well-defined zero value — never undefined memory.
- A zero-value string is a usable empty string, not `nil` — strings don't have a `nil` state.
- A `nil` slice is safely readable and even appendable; a `nil` map panics on **write** (but not
  read) until initialized with `make` or a literal.
- A zero-value struct recursively zero-values every one of its fields.

## 🛠️ Try It Yourself

1. Try writing to the `nil` map (`m["key"] = 1`) and observe the panic — then fix it with
   `m = make(map[string]int)` beforehand and confirm the write succeeds.
2. Try `append`ing to the `nil` slice (`slice = append(slice, 1)`) and confirm it works fine with
   no prior initialization.
3. Add a nested struct field (a struct containing another struct) and confirm the zero value
   propagates correctly at every level.

## ⚠️ Common Mistakes

- Assuming a `nil` map behaves like a `nil` slice for writing — reading from either is safe;
  writing to a `nil` map panics, writing (via `append`) to a `nil` slice does not.
- Checking `if s == ""` to mean "string wasn't set" when actually receiving a **non-nil** empty
  string is a perfectly normal, valid value — the zero value and "explicitly set to empty" are
  indistinguishable for strings, which is sometimes exactly what you want and sometimes a real
  ambiguity worth designing around (compare to `os.LookupEnv` in
  [lesson 43 of level 00](../../level-00-getting-started/43-environment-basics)).
