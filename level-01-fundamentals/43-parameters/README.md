# 43 — Parameters

## 🎯 Learning Objectives

- Confirm Go is **always** pass-by-value — there is no pass-by-reference for ordinary types.
- Understand why slices and maps *appear* mutable through a function call, despite this rule.
- See the specific case where a slice modification inside a function does **not** propagate back.

## 📖 Concept

Go has exactly one parameter-passing convention: **pass-by-value**. Every argument is **copied**
into the function's parameter — there is no pass-by-reference mechanism for ordinary values, full
stop.

### Plain values: the simple case

```go
func tryToDouble(n int) {
    n *= 2 // modifies the LOCAL COPY only
}
```

`n` inside the function is a completely separate copy of whatever the caller passed — changes to
it never propagate back. This matches every plain value type: `int`, `float64`, `bool`, `string`,
structs, arrays.

### Slices and maps: what actually gets copied is a small "header"

A slice's value isn't the underlying array itself — it's a small **header**: a pointer to the
array, a length, and a capacity. When you pass a slice to a function, that **header** is copied
(matching the pass-by-value rule exactly) — but the header's **pointer** still points at the same
underlying array the caller has:

```go
func modifyFirstElement(s []int) {
    s[0] = 999 // writes through the SHARED pointer — the caller sees this
}
```

Because both the caller's slice and the function's copy of it point at the **same** underlying
array, writing to an existing element through either one is visible to both. This is why slices
*look* like they're passed by reference, even though — strictly — only their small header was
actually copied. Maps work the same way: a map value is a pointer to shared underlying storage.

### The case where this breaks down: `append`

```go
func appendToSlice(s []int) {
    s = append(s, 100) // might allocate a brand-new array
}
```

If a slice is already at full **capacity**, `append` must allocate an entirely **new**,
larger underlying array to fit the additional element — and copies the existing elements into it.
The function's **local** slice header now points at this new array, but the **caller's** slice
header (their own separate copy) still points at the **old** one — completely unaware anything
happened. The caller's slice is left exactly as it was before the call.

## 🔍 Code Walkthrough (`main.go`)

```go
small := make([]int, 2, 2) // length 2, CAPACITY 2 — deliberately no room to grow
...
appendToSlice(small)
fmt.Printf("... small = %v (UNCHANGED ...)\n", small)
```

`small` is deliberately created with **capacity exactly equal to its length**, guaranteeing
`append` inside the function must reallocate — making the "caller doesn't see the append" behavior
reliable and reproducible, rather than depending on happenstance capacity.

```go
func modifyMapEntry(m map[string]int) {
    m["new"] = 1
}
```

Unlike the slice-append case, there's no equivalent "capacity" gotcha for maps in this basic
scenario — writing a new key through the shared underlying map storage is reliably visible to the
caller.

## ▶️ How to Run

```bash
cd level-01-fundamentals/43-parameters
go run main.go
```

## ✅ Expected Output

```
=== Parameters ===
----------------------------------
--- Plain values: always copied, never affected ---
before tryToDouble: x = 5
  inside tryToDouble: n is now 10
after tryToDouble:  x = 5 (COMPLETELY unchanged)

--- Slices: the HEADER is copied, but it points at SHARED data ---
before modifyFirstElement: nums = [1 2 3]
after modifyFirstElement:  nums = [999 2 3] (element WAS changed — shared array)

--- append() inside a function may NOT affect the caller ---
before appendToSlice: small = [1 2] (len=2, cap=2)
after appendToSlice:  small = [1 2] (UNCHANGED — append had to allocate a new array)

--- Maps: behave like slices' shared-data property, but for the WHOLE map ---
before modifyMapEntry: ages = map[Alice:30]
after modifyMapEntry:  ages = map[Alice:30 new:1] (new entry IS visible — shared storage)
```

## 🧠 Key Takeaways

- Go is always pass-by-value — plain types are fully copied, with zero effect on the caller.
- Slices and maps carry a pointer to shared underlying data inside their (copied) header, which
  is why modifying an existing element/key through them is visible to the caller.
- `append` can allocate a new underlying array, silently breaking the "shared data" illusion —
  the caller's original slice header has no way to observe that reallocation happened.
- If a function needs to change a slice's **length** in a way the caller sees, it must **return**
  the new slice and have the caller reassign it — this is exactly why `append`'s own idiomatic
  usage is always `s = append(s, x)`, never a bare `append(s, x)`.

## 🛠️ Try It Yourself

1. Create `small` with extra spare capacity instead (`make([]int, 2, 10)`) and confirm
   `appendToSlice` **does** now affect the caller — since there's room to grow in place, no
   reallocation is needed.
2. Write a function that takes a slice and reassigns its **entire contents** with a fresh
   `s = []int{...}` inside — confirm the caller is unaffected, since this only changes the local
   header's pointer, not any shared array.
3. Rewrite `tryToDouble` to take `n *int` (a pointer) instead of a plain `int`, dereference it to
   modify the caller's value directly — a preview of pointers, covered in a later lesson.

## ⚠️ Common Mistakes

- Assuming slices/maps are passed "by reference" in the way some other languages use that term —
  strictly, only their small header is copied; the *shared underlying data* is what creates the
  illusion, and `append`'s reallocation case shows exactly where that illusion breaks down.
- Calling `append` inside a function and expecting the caller's slice to reflect the new,
  possibly-longer result — without returning the new slice back to the caller, there's no
  guarantee it will.
