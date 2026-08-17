# 04 — package main

## 🎯 Learning Objectives

- Explain what a **package** is in Go.
- Explain why `package main` specifically is special.
- Understand that a package can contain many functions, types, and variables — not just `main()`.
- Know the difference between a **main package** (produces a program) and a **library package**
  (produces reusable code that other packages import).

## 📖 Concept

Every `.go` file starts with a package declaration:

```go
package main
```

A **package** is simply "a collection of `.go` files, in the same folder, that belong together."
Everything declared in those files — functions, types, variables, constants — is visible to every
other file in the same package, with no import needed.

### Why `main` is special

Go treats the package name `main` as a signal: _"this package produces an executable program,
not a library."_ Two rules follow from that:

1. A package named `main` **must** contain a function named exactly `func main()` with no
   parameters and no return values. That function is the program's entry point.
2. When you run `go build` or `go run` on a `main` package, Go produces a runnable binary. Every
   other package name produces something that can only be **imported**, not executed directly.

### Library packages, for comparison

Imagine a second folder in this repo, `greeting/greeting.go`, that starts with:

```go
package greeting

func Hello(name string) string {
    return "Hello, " + name + "!"
}
```

That package has **no `main()`** and **cannot be run directly**. Instead, another `package main`
file would `import` it and call `greeting.Hello("Gopher")`. This is exactly how the Go standard
library works — `fmt`, `runtime`, and every package you've used so far are library packages that
`package main` files import and call.

## 🔍 Code Walkthrough (`main.go`)

This lesson's `main.go` deliberately defines **three** functions — `greet`, `farewell`, and
`main` — to make the point visible: a package is a _collection_ of declarations, and `main()` is
just the one Go calls first when the program starts.

```go
func greet(name string) string {
    return "Hello, " + name + "!"
}
```

This is a regular function: it takes a `string` parameter and returns a `string`. Nothing about
it is special — it just happens to live in `package main` alongside the entry point.

## ▶️ How to Run

```bash
cd level-00-getting-started/04-package-main
go run main.go
```

## ✅ Expected Output

```
Hello, Gopher!
Goodbye, Gopher.

This file declares 'package main', which is what tells the Go
toolchain: 'this produces a runnable program', not a library.
```

## 🧠 Key Takeaways

- A package is "the `.go` files in one folder, working together."
- `package main` + `func main()` is the specific combination that makes a program runnable.
- Any other package name produces a library that must be imported to be used.
- A package is free to contain many functions — `main()` is just the designated starting point.

## 🛠️ Try It Yourself

1. Add a third function, `shout`, that returns the greeting in uppercase (hint: look up
   `strings.ToUpper` — you'll need to `import "strings"` too).
2. Call it from `main()` and print the result.
3. Rename `package main` to `package demo` and try `go run main.go` — read the error Go gives
   you, then rename it back.

## ⚠️ Common Mistakes

- Naming a file's package something other than `main` and then being confused why `go run`
  refuses to execute it.
- Assuming a folder's **name** determines the package name — it doesn't. The package name comes
  from the `package` declaration in the files themselves (though by convention they usually match).
