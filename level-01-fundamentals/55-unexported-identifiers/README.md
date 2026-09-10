# 55 — Unexported Identifiers

## 🎯 Learning Objectives

- Define an unexported identifier: any name starting with a **lowercase** letter.
- Confirm unexported identifiers are genuinely unreachable from another package — for both
  reading and writing, with no exceptions.
- Use the "exported constructor + exported methods, unexported fields" pattern for encapsulation.

## 📖 Concept

[Lesson 54](../54-exported-identifiers) covered exported (uppercase) identifiers. Anything
starting with a **lowercase** letter is the opposite: **unexported**, visible only within the
package that declares it — genuinely inaccessible from anywhere else, for any reason.

```go
type Counter struct {
    value int // unexported — invisible outside package counter, PERIOD
    step  int  // also unexported
}
```

### The standard encapsulation pattern: exported API, unexported internals

```go
func New(step int) *Counter {          // EXPORTED constructor
    return &Counter{value: 0, step: step}
}

func (c *Counter) Increment() {         // EXPORTED method
    c.value += c.step
}

func (c *Counter) Value() int {          // EXPORTED, read-only accessor
    return c.value
}
```

Because `Counter`'s fields are unexported, code outside the `counter` package **cannot**
construct one with a struct literal (`counter.Counter{value: 0, step: 5}` simply won't compile
from `main.go` — `value` and `step` aren't visible there) and cannot read or write those fields
directly, from any external package, under any circumstance. The **only** way in or out is
through the package's deliberately exported functions and methods — exactly the encapsulation
guarantee unexported fields provide.

### This is a hard compiler rule, not a convention

Unlike some languages where "private" is more of a suggestion (accessible via reflection,
underscore-prefix conventions, etc.), Go's unexported identifiers are enforced by the
**compiler**: there is no mechanism, short of being literally inside the same package, to
reference one from outside it.

## 🔍 Code Walkthrough (`counter/counter.go` and `main.go`)

```go
c := counter.New(5)
c.Increment()
fmt.Printf("... %d\n", c.Value())
```

Every interaction with `c` in `main.go` goes through `counter`'s exported API —
`New`/`Increment`/`Value` — never touching `value`/`step` directly, because it's structurally
impossible to from this package.

## ▶️ How to Run

```bash
cd level-01-fundamentals/55-unexported-identifiers
go run main.go
```

## ✅ Expected Output

```
=== Unexported Identifiers ===
----------------------------------
After 3 increments of step 5: 15

See the README for the exact compiler error you'd get trying to
access counter.Counter's fields directly from this package.
```

## 🧠 Key Takeaways

- A lowercase first letter makes an identifier unexported — reachable only within its own package.
- This is enforced by the compiler, not just a naming convention — there is no workaround from
  outside the package.
- The standard Go encapsulation pattern: unexported fields, plus an exported constructor and
  exported methods as the only sanctioned way to create and interact with the type.
- This is precisely how much of the Go standard library itself is built (e.g. `bytes.Buffer`,
  `strings.Builder` hide their internal fields behind exported methods).

## 🛠️ Try It Yourself

1. In a scratch copy of `main.go`, try `c := counter.Counter{value: 5}` (naming the unexported
   field directly) and read the exact compiler error. Note that `counter.Counter{}` with **no**
   named fields compiles fine on its own — the error specifically triggers when you try to name
   an unexported field.
2. Add a `Reset()` method to `Counter` that sets `value` back to `0`, and call it from `main.go`.
3. Add a **new**, exported field directly to `Counter` (e.g. `Label string`) and confirm you
   *can* set it directly from `main.go` — contrasting an exported field's accessibility against
   `value`/`step`'s.

## ⚠️ Common Mistakes

- Assuming an unexported field is merely "hidden by convention" and looking for a workaround to
  access it anyway from outside the package — there genuinely isn't one, short of modifying the
  package itself.
- Exporting every field on a struct by default "just in case," rather than deliberately choosing
  which parts of a type are genuinely meant to be part of its public API.
