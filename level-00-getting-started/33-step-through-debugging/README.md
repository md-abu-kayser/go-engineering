# 33 — Step-Through Debugging

## 🎯 Learning Objectives

- Tell `next`, `step`, and `stepout` apart, precisely, not just by vague description.
- Navigate a real chain of nested function calls using all three deliberately.
- Know which one to reach for depending on what you're trying to investigate.

## 📖 Concept

[Lesson 31](../31-debugging-with-delve) introduced these commands briefly. This lesson is
dedicated entirely to using them correctly, since mixing them up is the single most common
source of "wait, why did the debugger just do that?" confusion for people new to Delve.

| Command | What it does |
|---|---|
| `next` (`n`) | Executes the current line and stops at the **next line in the same function** — if the current line calls another function, that call runs to completion **without** pausing inside it. |
| `step` (`s`) | Executes the current line, but if it's a function call, **pauses at the first line inside that call** instead of skipping over it. |
| `stepout` (`so`) | Runs until the **current function returns**, then pauses in the caller, right after the call that got you here. |

### A concrete call chain to practice on

```
main() -> checkout() -> applyTax() -> (returns to checkout) -> formatPrice() -> (returns to checkout) -> (returns to main)
```

## 🔍 A worked session

```bash
dlv debug .
```

```
(dlv) break main.checkout
(dlv) continue
> main.checkout() ./main.go:18
(dlv) next
> main.go:19 (withTax now holds applyTax's result)
```

Now rewind mentally and instead **step into** `applyTax`:

```
(dlv) restart
(dlv) continue
> main.checkout() ./main.go:18
(dlv) step
> main.applyTax() ./main.go:13   (we're now INSIDE applyTax)
(dlv) stepout
> main.checkout() ./main.go:19   (back in checkout, right after the call)
```

Notice `next` on line 18 (which calls `applyTax`) would have run the entire function invisibly;
`step` instead took you inside it; and `stepout`, once inside, took you straight back out to the
caller without you having to `next` through every remaining line of `applyTax` one at a time.

## 🔍 Code Walkthrough (`main.go`)

`checkout` deliberately calls two different functions (`applyTax`, then `formatPrice`) so you can
practice `step`-ing into one, `stepout`-ing back, and then choosing `next` vs `step` again for the
second call — the same decision, twice, in one debugging session.

## ▶️ How to Run

```bash
cd level-00-getting-started/33-step-through-debugging
go run main.go
dlv debug .
```

## ✅ Expected Output (normal run)

```
Final price: $21.68

See the README to practice next/step/stepout on this exact call chain.
```

## 🧠 Key Takeaways

- `next` stays in the current function, treating any call on that line as a black box.
- `step` follows a call into the callee, one level deeper.
- `stepout` runs to the end of the current function and pauses back in its caller.
- Choosing `next` vs `step` is really "do I trust this function call, or do I need to look inside it?"

## 🛠️ Try It Yourself

1. Set a breakpoint at the top of `main`, then use `step` (not `next`) all the way into
   `checkout`, `applyTax`, and `formatPrice`, one call at a time.
2. From inside `formatPrice`, use `stepout` and observe you land back in `checkout`, at the
   `return formatPrice(withTax)` line.
3. Redo the whole walk using only `next` from `main`, and notice you never leave `main` at all —
   confirming `next` really does treat function calls as opaque.

## ⚠️ Common Mistakes

- Using `step` when you meant `next` on a call into a well-trusted standard-library function
  (e.g. `fmt.Sprintf`) — you'll suddenly find yourself deep in `fmt`'s internals, which is rarely
  where you wanted to be.
- Forgetting `stepout` exists and instead mashing `next` repeatedly to "get back" to the caller —
  `stepout` does the same thing in one command.
