# 53 — Package Scope

## 🎯 Learning Objectives

- Declare a variable at **package scope** — outside any function.
- Confirm package-scoped declarations are visible to every function in the package, with no
  parameter-passing needed.
- Understand that package scope spans **every file** in the same package, not just one.

## 📖 Concept

[Lesson 51](../51-scope) covered block and function scope. There's one scope broader than
either: **package scope** — anything declared outside every function, directly in a `.go` file
at the top level.

```go
var appName = "GO-ENGINEERING" // package-scoped: visible to EVERY function below, in ANY file
```

### Visible everywhere in the package, no passing required

```go
func logRequest(path string) {
    requestCount++ // reads AND writes requestCount directly — no parameter needed
}

func resetCount() {
    requestCount = 0
}
```

Both functions use `requestCount` directly, despite neither receiving it as a parameter or
returning it — this is only possible because `requestCount` lives at package scope, reachable
from anywhere in the package without any explicit wiring.

### Package scope spans every file in the package — not just one

This lesson's example lives in a single `main.go`, but the rule is genuinely about the
**package**, not the file: if this package had a second file, say `helpers.go`, declared with the
same `package main` line at its top, that file could use `appName` and `requestCount` directly
too — no `import` needed, since they're already in the same package. Import statements are for
reaching into **other** packages; within one package, every file already shares one common scope.

### Trade-offs of package-scoped state

Package-level variables are convenient — but the same convenience that makes `requestCount`
usable from anywhere also means **any** function in the package can silently modify it, which can
make a larger program's behavior harder to trace than passing state explicitly through function
parameters and return values. Use package scope deliberately, for genuinely shared,
package-wide state — not as a default way to avoid passing values around.

## 🔍 Code Walkthrough (`main.go`)

```go
var appName = "GO-ENGINEERING"
var requestCount int
```

Both are declared **outside** `func main` and every other function — that's what makes them
package-scoped rather than function-scoped, matching exactly [lesson 01](../01-var-declarations)'s
earlier mention of package-level `var` declarations, now explained in terms of scope specifically.

## ▶️ How to Run

```bash
cd level-01-fundamentals/53-package-scope
go run main.go
```

## ✅ Expected Output

```
=== Package Scope ===
----------------------------------
appName is visible right here too, in main(): GO-ENGINEERING

--- Multiple functions sharing package-scoped state ---
  [GO-ENGINEERING] request #1 to /home
  [GO-ENGINEERING] request #2 to /about
  [GO-ENGINEERING] request #3 to /contact
requestCount after 3 requests: 3
requestCount after resetCount(): 0

See the README: this file's package-level declarations would ALSO be
visible from a second .go file, if this package had one — no import needed.
```

## 🧠 Key Takeaways

- Package-scoped declarations (outside any function) are visible to every function in the package.
- This spans every file sharing the same `package` declaration — not just the file where the
  variable is written.
- No import is needed to use another file's package-scoped declarations within the same package;
  imports are only for reaching into **other** packages.
- Package-scoped state is convenient but can make a program's data flow harder to trace — use it
  deliberately, not as a default.

## 🛠️ Try It Yourself

1. Add a second package-scoped variable, `debugMode bool`, and use it inside `logRequest` to
   optionally print extra detail.
2. In a scratch copy of this lesson, split it into two files (`main.go` and `helpers.go`, both
   `package main`), moving `logRequest` into the new file — confirm it still compiles and runs
   identically, with no import added.
3. Rewrite `logRequest`/`resetCount` to take and return `requestCount` explicitly as a parameter
   instead of using package scope, and compare the two versions' call sites.

## ⚠️ Common Mistakes

- Reaching for a package-level variable as a shortcut to avoid passing a value through several
  function calls, when explicit parameters would make the data flow clearer to a future reader.
- Assuming package scope requires being in the same **file** — it only requires being in the same
  **package** (the same `package` declaration), which commonly spans many files.
