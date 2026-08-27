# 34 — Variable Inspection

## 🎯 Learning Objectives

- Use `print` to inspect a struct, a slice, and a map from inside a Delve session.
- Use `whatis` to check a value's exact type.
- Use `set` to change a variable's value while paused, and observe the effect.

## 📖 Concept

[Lesson 31](../31-debugging-with-delve) used `print` on plain integers. Real programs have
richer data — structs, slices, maps — and Delve can inspect all of them in full, without you
writing a single custom `fmt.Printf`.

### Inspecting a struct

```
(dlv) print o
main.order {ID: 42, Items: []string len: 3, cap: 3, [...], Metadata: map[string]string [...]}
```

Delve prints the **entire nested value** — every field, recursively — not just a memory address.

### Inspecting a slice and a map specifically

```
(dlv) print o.Items
[]string len: 3, cap: 3, ["Keyboard","Mouse","Monitor"]
(dlv) print o.Metadata
map[string]string ["priority": "high", "region": "eu-west"]
(dlv) print o.Items[1]
"Mouse"
(dlv) print len(o.Items)
3
```

Notice Delve understands Go expressions directly — indexing, `len()`, field access — not just bare
variable names.

### Checking a type with `whatis`

```
(dlv) whatis o
main.order
(dlv) whatis o.Metadata
map[string]string
```

Useful when you're not sure exactly what type an interface value or a `%v`-printed value actually
holds.

### Changing a value with `set`

```
(dlv) set o.ID = 999
(dlv) print o.ID
999
```

`set` lets you change a variable's value mid-debug and then `continue` to see how the rest of the
program behaves differently — a fast way to test "what if this had been different?" without
editing code and recompiling.

## 🔍 Code Walkthrough (`main.go`)

`order` deliberately has one of each common shape — a scalar field (`ID`), a slice (`Items`), and
a map (`Metadata`) — specifically so this lesson gives you all three to practice `print`/`whatis`/
`set` against in one session.

## ▶️ How to Run

```bash
cd level-00-getting-started/34-variable-inspection
go run main.go
dlv debug .
```

Set a breakpoint inside `summarize` (`break main.summarize`), `continue`, then work through the
commands above.

## ✅ Expected Output (normal run)

```
Order #42 has 3 item(s)

See the README to inspect (and even modify) o's fields with dlv print/set.
```

## 🧠 Key Takeaways

- `print` understands full Go expressions: field access, indexing, `len()`, and more.
- `print` on a struct shows every field recursively, not just a summary.
- `whatis` tells you a value's exact type when that's genuinely in question.
- `set` changes a variable mid-session — a fast way to test hypotheses without recompiling.

## 🛠️ Try It Yourself

1. Break inside `summarize`, and `print o` to see the whole struct at once.
2. Use `set o.ID = 999` before `continue`, and confirm the printed summary still shows the
   original `#42` (since `summarize` already read `o.ID` into `total` differently — think about
   *why*, and check whether setting it before vs after the breakpoint's line changes the outcome).
3. Add a `Price float64` field to `order`, and practice `print o.Price` and `whatis o.Price`.

## ⚠️ Common Mistakes

- Assuming `print` on a large slice or map always shows everything — Delve truncates very large
  collections by default; there are configuration options to raise that limit if you need it.
- Using `set` on a variable's *type* mismatched value (e.g. trying to set a `string` field to a
  bare number) — Delve will reject it, the same way the Go compiler would.
