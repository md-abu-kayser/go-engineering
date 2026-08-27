# 27 — Documentation Comments

## 🎯 Learning Objectives

- Use headings, lists, and links inside Go doc comments.
- Write an indented code block inside a doc comment.
- Mark a symbol as deprecated in a way tooling recognizes.
- Link to another symbol from within a doc comment.

## 📖 Concept

[Lesson 13](../13-go-doc-command) covered the basic rule — a plain comment starting with the
symbol's name. This lesson covers the **specific formatting syntax** Go's doc renderer (used by
`go doc`, `godoc`, and pkg.go.dev) understands inside that comment.

### Headings

```go
// # Units
```

A line starting with `#` becomes a heading — useful for splitting a long package or type doc
comment into sections (`# Overview`, `# Example`, `# Errors`, etc.).

### Code blocks

```go
// A typical reading is created like this:
//
//	t := Temperature{Celsius: 21.5}
//	fmt.Println(t.Fahrenheit())
```

A line indented with a **tab** (or that starts after a blank comment line and is indented)
renders as a code block — exactly how the usage snippet in `main.go`'s `Temperature` doc comment
is written.

### Links

```go
// See also the Go blog post on doc comments: https://go.dev/doc/comment
```

Bare URLs are automatically turned into clickable links in rendered output (e.g. on pkg.go.dev).

### Linking to another symbol

```go
// Use [Temperature.Fahrenheit] to convert for display when needed.
```

Square brackets around a symbol name create a **cross-reference link** to that symbol's own
documentation — this only works for symbols in the same package or an imported one.

### Deprecation notices

```go
// Deprecated: prefer working in Celsius directly; this method is kept
// only for display purposes in US-locale output.
```

A paragraph starting with exactly `Deprecated:` is recognized by tooling (including IDEs like VS
Code) and typically renders with a strikethrough or warning — a real, machine-readable signal,
not just a comment a human might skim past.

## 🔍 Code Walkthrough (`main.go`)

The `Temperature` type's doc comment in this lesson intentionally uses **every** formatting
feature above in one place — a heading (`# Units`), a cross-reference link
(`[Temperature.Fahrenheit]`), a code block, and a bare URL — so you can see them all together
and then look at how `go doc` renders each one.

## ▶️ How to Run

```bash
cd level-00-getting-started/27-documentation-comments
go run main.go
go doc -all .
go doc Temperature
```

## ✅ Expected Output

```
21.5°C = 70.7°F

Run `go doc -all .` in this folder to see how the comments above render.
```

## 🧠 Key Takeaways

- `# Heading` creates a section heading inside a doc comment.
- Tab-indented lines render as a code block.
- `[Symbol]` or `[Type.Method]` creates a cross-reference link.
- A paragraph starting with `Deprecated:` is a recognized, tooling-visible deprecation marker.

## 🛠️ Try It Yourself

1. Run `go doc -all .` and compare the rendered output line-by-line against the raw comment in
   `main.go`.
2. Add a second heading, `# Limitations`, with a short bullet list under it (lines starting with
   `-` or `+`) describing a constraint of `Temperature`.
3. Hover over `Temperature` in VS Code (with the Go extension installed) and confirm the same
   formatted documentation appears in the hover tooltip.

## ⚠️ Common Mistakes

- Using spaces instead of a real tab to indent a code block — the renderer specifically looks for
  the indentation `gofmt` produces.
- Writing `Deprecated - ...` or `DEPRECATED:` instead of exactly `Deprecated:` — tooling matches
  this marker literally.
