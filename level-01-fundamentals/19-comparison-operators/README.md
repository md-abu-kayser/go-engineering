# 19 — Comparison Operators

## 🎯 Learning Objectives

- Use all six comparison operators: `==` `!=` `<` `<=` `>` `>=`.
- Know that string comparison is lexicographic (byte-by-byte), not based on length.
- Understand exactly which types are comparable with `==` in Go — and which genuinely aren't.

## 📖 Concept

```go
a == b   // equal
a != b   // not equal
a < b     // less than
a <= b    // less than or equal
a > b     // greater than
a >= b    // greater than or equal
```

Straightforward for numbers. The genuinely important nuance in this lesson is **comparability**:
which Go types `==`/`!=` are even legal to use on in the first place.

### Strings compare lexicographically

```go
"apple" < "banana"  // true — compared byte-by-byte, like dictionary order
"zoo" < "apple"      // false — length has nothing to do with it; 'z' > 'a'
```

String comparison walks both strings byte-by-byte and compares based on byte values — exactly
like alphabetizing words in a dictionary, and completely unrelated to which string is longer.

### Structs: comparable, conditionally

```go
type point struct{ X, Y int }
p1 := point{1, 2}
p2 := point{1, 2}
p1 == p2  // true — compares EVERY field
```

A struct is comparable with `==` **if and only if every one of its fields is itself comparable**.
Since `int` is comparable, `point` (built entirely from `int` fields) is comparable too, and `==`
compares all fields together, field by field.

### Slices, maps, and functions: NOT comparable (with one exception)

```go
var s1, s2 []int
s1 == s2   // COMPILE ERROR — invalid operation, slices are not comparable
```

Slices, maps, and functions cannot be compared with `==` **at all** — this is a compile-time
error, not something that evaluates to `false` at runtime. The **one** exception: any of these can
be compared against the literal `nil`:

```go
s1 == nil   // fine — checking "is this slice nil" is allowed
```

This means if a struct **contains** a slice or map field, that struct itself becomes
**not comparable** either — `==` on it becomes a compile error, since Go can't compare one of
its fields.

## 🔍 Code Walkthrough (`main.go`)

```go
p1 := point{X: 1, Y: 2}
p2 := point{X: 1, Y: 2}
fmt.Printf("p1 == p2 : %t\n", p1 == p2)
```

This works specifically because `point`'s only fields are `int`s — both comparable on their own.
If `point` had a `Tags []string` field instead, this exact same `==` comparison would become a
compile error, not a runtime value.

## ▶️ How to Run

```bash
cd level-01-fundamentals/19-comparison-operators
go run main.go
```

## ✅ Expected Output

```
=== Comparison Operators ===
----------------------------------
10 == 20 : false
10 != 20 : true
10 <  20 : true
10 <= 20 : true
10 >  20 : false
10 >= 20 : false

--- String comparison is lexicographic ---
"apple" < "banana" : true (compares byte-by-byte, like dictionary order)
"zoo" < "apple"  : false (shorter isn't automatically "less")

--- Struct comparability ---
p1 == p2 : true (same field values)
p1 == p3 : false (different field values)

--- Slices/maps are NOT comparable ---
See the README: `mySlice == otherSlice` is a compile-time error, not `false`.
s == nil : true (comparing to nil IS allowed, even though slice == slice is not)
```

## 🧠 Key Takeaways

- Go's six comparison operators work as expected for numbers and strings.
- String comparison is lexicographic (byte-by-byte), independent of length.
- A struct is comparable with `==` only if every one of its fields is comparable.
- Slices, maps, and functions can never be compared with `==` to each other — only to `nil` — and
  this restriction propagates to any struct containing one of them as a field.

## 🛠️ Try It Yourself

1. Add a `Tags []string` field to `point`, and try `p1 == p2` again — read the exact compiler
   error that results.
2. Compare two maps directly with `==` and read the compiler's error message.
3. Write a manual "deep equal" check for two slices (comparing length, then each element) since
   `==` can't do it — or look up `reflect.DeepEqual`/`slices.Equal` for the standard-library way.

## ⚠️ Common Mistakes

- Trying to compare two slices or maps directly with `==`, expecting either `true`/`false` — it's
  a compile error, not a comparison that runs and returns `false`.
- Assuming string comparison is based on length ("longer string is greater") — it's strictly
  lexicographic, byte by byte, exactly like dictionary ordering.
