# 56 — Doc Comments

## 🎯 Learning Objectives

- Know **which** identifiers genuinely need a doc comment, and why — tied directly to
  [lesson 54](../54-exported-identifiers)/[55](../55-unexported-identifiers)'s exported/unexported
  distinction.
- Confirm `go doc` only ever shows **exported** identifiers to outside consumers of a package.
- Recognize that unexported identifiers can still benefit from comments, for a different audience.

## 📖 Concept

> **Scope note:** [level 00's lessons 13, 27, and 28]
> (../../level-00-getting-started/13-go-doc-command) already covered doc comment **formatting**
> in depth — headings, code blocks, links, package-level style. This lesson doesn't repeat that;
> it answers a narrower, different question: **which** identifiers actually need one, tied
> directly to exported vs. unexported status.

### Exported identifiers: documentation is for a real audience you can't see

```go
// Temperature represents a temperature reading in Celsius.
type Temperature struct {
    // Celsius is the reading itself.
    Celsius float64
    ...
}
```

An exported identifier can be used by **any** other package that imports yours — code you'll
never see, written by people who can't casually go read your internal implementation to figure
out what something does. `go doc` (and pkg.go.dev, for published modules) shows these doc
comments precisely because they're the **only** information an external consumer has to go on.
This is why exported identifiers should always have a doc comment: there's a real, distant
audience depending on it.

### Unexported identifiers: still worth commenting, for a different reason

```go
// clampToRealistic is unexported — no outside caller will ever see it.
func clampToRealistic(c float64) float64 {
```

An unexported identifier is only ever visible to people editing **this exact package** — often,
that's still a real audience (a teammate, or future-you) who benefits from a short explanation,
just not the same strict "this is the only info an external user has" stakes an exported doc
comment carries. A brief, informal note is often enough; the elaborate structure
([lesson 27's headings, code blocks, and cross-references]
(../../level-00-getting-started/27-documentation-comments)) matters far more for exported,
publicly-consumed symbols.

### `go doc` only ever shows exported identifiers to outside consumers

```bash
go doc .        # from OUTSIDE the package's own source — shows exported symbols only
go doc -all .    # same — even -all only lists what's actually exported
```

This is a direct, practical consequence of [lesson 55](../55-unexported-identifiers)'s
visibility rule: since `source` and `clampToRealistic` are literally unreachable from outside the
package, `go doc` (which is answering "what can I, an external consumer, use here?") simply never
lists them — there'd be nothing useful to say about a symbol you couldn't call anyway.

## 🔍 Code Walkthrough (`main.go`)

```go
// source is UNEXPORTED — an internal implementation detail...
source string
```

Notice `source`'s comment is genuinely useful to someone editing this package later, but it
doesn't follow the strict "start with the exact field name, write a complete sentence for an
external reader" convention as rigidly as `Celsius`'s comment does — that's a deliberate,
reasonable difference in register given the very different audiences.

## ▶️ How to Run

```bash
cd level-01-fundamentals/56-doc-comments
go run main.go
go doc .
go doc -all .
```

## ✅ Expected Output

```
=== Doc Comments ===
----------------------------------
21.5°C = 70.7°F

Run `go doc .` and `go doc -all .` here — notice unexported
identifiers like `source` and `clampToRealistic` never appear.
```

`go doc -all .` will list `Temperature`, `Temperature.Celsius`, and `Temperature.Fahrenheit` —
and nothing about `source` or `clampToRealistic` at all.

## 🧠 Key Takeaways

- Exported identifiers need real doc comments — an external consumer has no other way to
  understand them.
- Unexported identifiers can still benefit from a comment, for the package's own future
  maintainers — a different, closer audience with lower documentation stakes.
- `go doc` (from outside the package) only ever surfaces exported identifiers — this is a direct
  consequence of Go's visibility rule, not a separate documentation-specific behavior.
- Deciding "does this need a doc comment" is really "who is the audience, and what do they already
  have access to" — exported vs. unexported is exactly that question, made concrete.

## 🛠️ Try It Yourself

1. Run `go doc -all .` in this lesson's folder and confirm `source` and `clampToRealistic` never
   appear anywhere in the output.
2. Add a new exported method to `Temperature` with **no** doc comment at all, then run
   `go doc -all .` and notice it still appears — undocumented, but still exported and thus still
   listed; presence in `go doc` and having a *good* comment are two separate things.
3. Consider a real project you've worked on (or this repository itself) and identify one
   exported identifier that's under-documented relative to how it's actually used elsewhere.

## ⚠️ Common Mistakes

- Writing elaborate, external-consumer-oriented doc comments on unexported helpers that will
  never be seen by anyone outside the package — proportionate effort matters; save the full
  formatting toolkit for genuinely exported, externally-consumed symbols.
- Assuming an exported identifier with **no** comment at all is somehow invisible or flagged by
  the compiler — Go compiles it fine either way; only human review, linters
  (`golint`/`staticcheck`), or team convention actually enforce documenting exported symbols.
