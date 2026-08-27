# 38 — Reading Error Messages

## 🎯 Learning Objectives

- Read Go's `error` interface and know what satisfies it.
- Construct an error with context using `fmt.Errorf` and `%w`.
- Check for a specific error with `errors.Is`, and understand why that beats `==`.
- Define and use a sentinel error.

## 📖 Concept

`error` in Go is refreshingly small — it's just an interface with one method:

```go
type error interface {
    Error() string
}
```

Anything with an `Error() string` method **is** an error. Most of the time you don't implement
this yourself — you construct one with `errors.New` or `fmt.Errorf`.

### Sentinel errors

```go
var ErrNotFound = errors.New("item not found")
```

A **sentinel error** is a specific, package-level `error` value that calling code can check for by
identity. `lookup` in this lesson returns this exact value (wrapped, see below) whenever an item
genuinely isn't found — as opposed to some other, unrelated failure.

### Wrapping with `%w`

```go
return 0, fmt.Errorf("lookup %q: %w", name, ErrNotFound)
```

`%w` is special — unlike `%s` or `%v`, it doesn't just format `ErrNotFound`'s message into the
string, it **wraps** it, preserving a link back to the original error value. The resulting error's
`.Error()` string reads naturally (`lookup "cherry": item not found`), but the identity of
`ErrNotFound` underneath is still recoverable.

### Checking with `errors.Is`, not `==`

```go
if errors.Is(err, ErrNotFound) {
    // ...
}
```

Because `err` here is actually `fmt.Errorf(...)`'s wrapper value, **not** `ErrNotFound` itself,
a plain `err == ErrNotFound` comparison would be `false` — the wrapping changed the concrete
value. `errors.Is` specifically unwraps any `%w`-wrapped chain and checks each layer, so it
correctly reports `true` here. **Always prefer `errors.Is` over `==`** once wrapping is involved
anywhere in your codebase — which, in practice, means always.

### `errors.As` — for wrapped custom error *types*

`errors.Is` checks for a specific error *value*. When you need to check for a specific error
*type* instead (to extract extra fields from it), the equivalent tool is `errors.As`:

```go
var pathErr *fs.PathError
if errors.As(err, &pathErr) {
    fmt.Println("failed path:", pathErr.Path)
}
```

## 🔍 Code Walkthrough (`main.go`)

```go
switch {
case err == nil:
    // success
case errors.Is(err, ErrNotFound):
    // the specific, expected failure
default:
    // something else entirely
}
```

This three-way switch is a common, idiomatic shape: handle success, handle the **specific**
expected error you know how to react to, and have a fallback for anything else — rather than
trying to pattern-match on the error's string text, which is fragile and can change.

## ▶️ How to Run

```bash
cd level-00-getting-started/38-reading-error-messages
go run main.go
```

## ✅ Expected Output

```
apple   : 12 in stock
banana  : 0 in stock
cherry  : lookup "cherry": item not found

See the README for exactly what %w does, and why errors.Is beats ==.
```

## 🧠 Key Takeaways

- `error` is just an interface with one method — `Error() string`.
- A sentinel error (`var ErrX = errors.New(...)`) is a specific value calling code can check for.
- `%w` in `fmt.Errorf` wraps an error, adding context while preserving its identity.
- `errors.Is` correctly sees through `%w` wrapping; `==` does not.

## 🛠️ Try It Yourself

1. Change `%w` to `%v` in `lookup` and add an `errors.Is` check — watch it now report `false`,
   proving wrapping specifically requires `%w`.
2. Add a second sentinel error, `ErrOutOfStock`, returned when `count == 0`, and handle it as its
   own case in `main`'s switch.
3. Look up `fmt.Errorf`'s documentation and find where it explains `%w` — note it explicitly
   allows at most limited multiple wraps depending on your Go version.

## ⚠️ Common Mistakes

- Comparing wrapped errors with `==` instead of `errors.Is` — this is probably the single most
  common real-world Go bug involving errors, since it silently "works" until wrapping is
  introduced somewhere in the call chain.
- Wrapping with `%v` instead of `%w` and wondering why `errors.Is` never matches — `%v` formats
  the message but discards the link back to the original error entirely.
