# 13 — go doc

## 🎯 Learning Objectives

- Write doc comments that follow Go's conventions.
- Use the `go doc` command to read documentation directly in the terminal.
- Understand how these same comments end up on [pkg.go.dev](https://pkg.go.dev).

## 📖 Concept

In Go, **documentation comments are just regular comments** placed immediately above a
declaration, with no special syntax, tags, or annotations required:

```go
// Greeter produces greetings for a configured language.
type Greeter struct {
    ...
}
```

The convention is simple and strictly enforced by tooling: the comment should be a complete
sentence that **starts with the name being documented**. `go doc`, `godoc`, and pkg.go.dev all
rely on this convention to render clean documentation automatically — there's no separate
documentation file to maintain and let go stale.

### Reading docs with `go doc`

```bash
go doc .                 # doc comment for the package itself, plus a summary of its exports
go doc Greeter            # doc comment for the Greeter type
go doc Greeter.Greet      # doc comment for one specific method
```

This is the exact same command you'd use to read documentation for **any** Go package, including
the standard library:

```bash
go doc fmt.Println
go doc strings.Builder
```

No internet connection required — `go doc` reads directly from source, either from your module
cache or the standard library shipped with your Go installation.

## 🔍 Code Walkthrough (`main.go`)

```go
// Greeter produces greetings for a configured language.
//
// The zero value of Greeter is ready to use and defaults to English.
type Greeter struct {
```

Two doc-comment conventions worth noting:

1. It starts with `Greeter` — the exact name of the thing being documented.
2. It documents the **zero value** behavior explicitly (`Greeter{}` with no fields set). This is
   a very common and appreciated pattern in idiomatic Go docs, since Go doesn't have constructors
   and the zero value matters.

```go
// Greet returns a greeting for name, in the Greeter's configured language.
```

Method docs follow the same rule — start with the method name (`Greet`), describe what it
returns and any notable behavior (the language fallback, in this case).

## ▶️ How to Run

```bash
cd level-00-getting-started/13-go-doc-command
go run main.go
go doc .
go doc Greeter.Greet
```

## ✅ Expected Output

```
Hello, Gopher!
হ্যালো, Gopher!

Run `go doc .` and `go doc Greeter.Greet` in this folder to see these
comments rendered as documentation.
```

## 🧠 Key Takeaways

- Doc comments are ordinary comments directly above a declaration — no special syntax.
- Convention: start the comment with the exact name being documented.
- `go doc <symbol>` reads documentation from source, offline, for your code or the standard library.
- The same comments power pkg.go.dev for any publicly published module.

## 🛠️ Try It Yourself

1. Run `go doc fmt` and skim the package-level documentation for `fmt` itself.
2. Add a new exported method to `Greeter` with a proper doc comment, then read it back with
   `go doc Greeter.<YourMethod>`.
3. Try `go doc -all .` to see every exported declaration's documentation at once.

## ⚠️ Common Mistakes

- Leaving a blank line between the comment and the declaration — this **disconnects** the
  comment from the symbol, and tooling will no longer associate them.
- Writing comments that don't start with the symbol's name — `go doc`/`golint`-style tools flag
  this as a convention violation.
