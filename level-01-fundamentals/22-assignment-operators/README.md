# 22 — Assignment Operators

## 🎯 Learning Objectives

- Use every compound assignment operator Go provides for arithmetic and bitwise operations.
- Recognize `x op= y` as shorthand for `x = x op y`, for any supported operator `op`.
- Use `+=` for string concatenation, the same way `+` itself works ([lesson 18](../18-arithmetic-operators)).

## 📖 Concept

Every binary operator covered in [lessons 18](../18-arithmetic-operators) and
[21](../21-bitwise-operators) has a **compound assignment** form: `x op= y` is shorthand for
`x = x op y`.

```go
n += 5    // n = n + 5
n -= 3     // n = n - 3
n *= 2      // n = n * 2
n /= 4       // n = n / 4
n %= 3        // n = n % 3

flags &= mask   // flags = flags & mask
flags |= mask    // flags = flags | mask
flags ^= mask     // flags = flags ^ mask
flags &^= mask     // flags = flags &^ mask  (lesson 21's bit-clear)

n <<= 2   // n = n << 2
n >>= 2    // n = n >> 2
```

This is purely notational convenience — every compound form has an exactly equivalent
non-compound spelling — but it's overwhelmingly the idiomatic style in real Go code for "update
this variable based on its own current value."

### `+=` on strings

Because `+` itself means concatenation for strings ([lesson 18](../18-arithmetic-operators)),
`+=` naturally inherits that meaning too:

```go
msg := "Hello"
msg += ", Gopher!"   // msg = msg + ", Gopher!"
```

## 🔍 Code Walkthrough (`main.go`)

Every operator from lessons 18 and 21 gets its compound form demonstrated here, in the same
order, specifically so this lesson reads as a direct "and here's the shorthand for everything you
just learned" companion to both of those lessons rather than introducing new concepts of its own.

## ▶️ How to Run

```bash
cd level-01-fundamentals/22-assignment-operators
go run main.go
```

## ✅ Expected Output

```
=== Assignment Operators ===
----------------------------------
n := 10          -> n = 10
n += 5           -> n = 15
n -= 3           -> n = 12
n *= 2           -> n = 24
n /= 4           -> n = 6
n %= 3           -> n = 0

--- Bitwise compound assignments ---
flags &= 0b1010  -> 1000
flags |= 0b0001  -> 1001
flags ^= 0b1111  -> 0110
flags &^= 0b0010 -> 0100
shifted <<= 3    -> 8
shifted >>= 1    -> 4

--- += on strings ---
msg += ", Gopher!" -> "Hello, Gopher!"
```

## 🧠 Key Takeaways

- `x op= y` is exact shorthand for `x = x op y`, for every arithmetic and bitwise operator.
- This is purely stylistic convenience — both forms are equally valid — but the compound form is
  the idiomatic default in real Go code.
- `+=` on strings works the same way `+` does: concatenation.

## 🛠️ Try It Yourself

1. Rewrite every compound-assignment line in `main.go` using its fully expanded `x = x op y`
   form, and confirm the program's output is completely unchanged.
2. Chain several compound assignments on the same variable in a row and predict the final value
   before running it.
3. Try `%=`  on a `float64` variable and confirm it's a compile error — matching `%`'s own
   integers-only restriction from [lesson 18](../18-arithmetic-operators).

## ⚠️ Common Mistakes

- Forgetting compound assignment operators exist and writing out `x = x + y` everywhere — not
  wrong, just unidiomatic; real Go code overwhelmingly prefers `x += y`.
- Assuming `%=`/`&=`/`|=`/`^=`/`&^=`/`<<=`/`>>=` work on floats — they inherit exactly the same
  type restrictions as their non-compound counterparts (`%`, bitwise operators are integer-only).
