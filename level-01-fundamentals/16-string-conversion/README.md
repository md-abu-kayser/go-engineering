# 16 — String Conversion

## 🎯 Learning Objectives

- Convert numbers to strings and back using the `strconv` package's `Itoa`/`Atoi`/`FormatFloat`/
  `ParseFloat`.
- Handle the error `Atoi`/`ParseFloat` return for genuinely invalid input.
- Recognize the classic `string(someInt)` gotcha, and why it doesn't do what most people expect.

## 📖 Concept

Converting between numbers and their text representation is common enough that Go has a
dedicated package, `strconv`, for exactly this — and one, very common trap involving the `string`
conversion syntax that this lesson exists specifically to head off.

### Number → string: `strconv.Itoa`

```go
strconv.Itoa(42) // "42" — "Integer to ASCII", an old, C-derived naming convention
```

### String → number: `strconv.Atoi`

```go
n, err := strconv.Atoi("123")
```

Unlike a numeric-to-numeric conversion ([lesson 15](../15-numeric-conversion)), parsing text
**can genuinely fail** — `"hello"` isn't a valid number — so `Atoi` returns an `error`
([lesson 38 of level 00](../../level-00-getting-started/38-reading-error-messages)) you must
check, exactly like any other fallible operation.

### Floats: `FormatFloat` and `ParseFloat`

```go
strconv.FormatFloat(f, 'f', 2, 64)   // 'f' = plain decimal format, 2 = decimal places, 64 = bit size
strconv.ParseFloat("2.71828", 64)     // the bit size must match the float type you'll store it as
```

`FormatFloat`'s extra parameters give you control `Itoa` doesn't need for integers — which
notation to use (`'f'` for plain decimal, `'e'` for scientific, and others), and how many decimal
places to keep.

### The classic gotcha: `string(anInt)` is NOT `strconv.Itoa`

This is one of the most common points of confusion for people learning Go:

```go
code := 65
string(rune(code))   // "A" — NOT "65"!
```

Converting an integer directly to `string` doesn't produce the number's **text** — it treats the
integer as a **Unicode code point** and produces the single **character** at that code point.
`65` is the ASCII/Unicode code point for `'A'`, so `string(rune(65))` gives you the letter `"A"`,
not the text `"65"`. If you actually want the digits as text, you want `strconv.Itoa`, not a bare
`string(...)` conversion.

(Modern Go actually requires the explicit `rune(...)` step shown above — `go vet` flags a direct
`string(intValue)` conversion as suspicious precisely because this confusion is so common.)

## 🔍 Code Walkthrough (`main.go`)

```go
wrong := string(rune(code))
right := strconv.Itoa(code)
```

Both lines are placed directly next to each other, using the **same** `code := 65`, specifically
so the difference in output — `"A"` versus `"65"` — is impossible to miss, rather than being an
abstract warning you might forget until you actually trip over it.

## ▶️ How to Run

```bash
cd level-01-fundamentals/16-string-conversion
go run main.go
```

## ✅ Expected Output

```
=== String Conversion ===
----------------------------------
strconv.Itoa(42)        = "42"
strconv.Atoi("123")     = 123, err = <nil>
strconv.Atoi("not a number") errors: strconv.Atoi: parsing "not a number": invalid syntax
strconv.FormatFloat(3.14159, 2 places) = "3.14"
strconv.ParseFloat("2.71828")          = 2.71828

--- The classic gotcha ---
string(rune(65)) = "A" (the CHARACTER with code point 65 — 'A')
strconv.Itoa(65) = "65" (the actual TEXT "65" — almost always what you want)
```

## 🧠 Key Takeaways

- `strconv.Itoa`/`strconv.Atoi` are the idiomatic number ↔ string conversions for integers.
- `strconv.Atoi`/`strconv.ParseFloat` return an `error` — invalid input is a normal, expected case
  to handle, not a program-ending event.
- `string(someInt)` treats the integer as a Unicode code point, producing a character — not the
  number's text. Use `strconv.Itoa` when you want the digits as text.

## 🛠️ Try It Yourself

1. Try `strconv.Itoa` and `string(rune(...))` on a few more numbers (`97`, `48`) and predict which
   characters `string(rune(...))` will produce before running it (hint: ASCII table).
2. Handle `strconv.Atoi`'s error properly with an `if err != nil` check, instead of just printing
   it, the way [lesson 38 of level 00](../../level-00-getting-started/38-reading-error-messages)
   demonstrated.
3. Try `strconv.FormatFloat` with the `'e'` format verb instead of `'f'` and observe the
   scientific-notation output.

## ⚠️ Common Mistakes

- Writing `string(n)` where `n` is an `int`, expecting the number's digits as text — this is
  precisely the gotcha this lesson exists to prevent; use `strconv.Itoa(n)` instead.
- Ignoring the error from `strconv.Atoi`/`ParseFloat` — user-supplied or file-read text is
  exactly the kind of input that can and will occasionally be invalid.
