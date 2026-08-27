# 36 — Compile Errors

## 🎯 Learning Objectives

- Read a Go compiler error's `file:line:column` location and message confidently.
- Recognize the half-dozen compile errors you'll hit constantly as a beginner.
- Understand why Go treats some things (unused variables, unused imports) as errors, not warnings.

## 📖 Concept

Every Go compiler error follows the same shape:

```
./main.go:12:6: <message>
```

`file:line:column`, then a colon, then a description. Learning to read that prefix quickly — go
straight to line 12, column 6 — saves far more time than reading the message first.

### The errors you'll see constantly

**Unused variable**

```go
func main() {
    x := 5
}
```
```
./main.go:2:2: declared and not used: x
```
Go refuses to compile a variable you declared but never read. Either use it, or replace `:=` with
`_ =` if you genuinely only need the side effect (rare).

**Unused import**

```go
import "strings"

func main() {}
```
```
./main.go:1:8: "strings" imported and not used
```
Covered in depth in [lesson 05](../05-imports). Remove the import, or actually use the package.

**Mismatched types**

```go
var age int = "fifteen"
```
```
./main.go:1:14: cannot use "fifteen" (untyped string constant) as int value in variable declaration
```
Go's static typing catches this before your program ever runs — contrast with dynamically-typed
languages where this might not surface until the line actually executes.

**Missing return**

```go
func getName() string {
    if true {
        return "Gopher"
    }
}
```
```
./main.go:5:1: missing return
```
Go's compiler isn't smart enough to prove that `if true` always returns — every code path out of a
function with a return type must have an explicit `return`.

**Undefined identifier (typo)**

```go
fmt.Prinln("hello")
```
```
./main.go:2:6: undefined: fmt.Prinln
```
Almost always a typo (`Prinln` vs `Println`) — read it literally; Go isn't guessing what you meant.

**Mismatched argument count**

```go
func add(a, b int) int { return a + b }
// ...
add(1, 2, 3)
```
```
./main.go:3:1: too many arguments in call to add
	have (number, number, number)
	want (int, int)
```
Notice Go shows you both what you **passed** and what was **expected** — read both lines.

## 🔍 Code Walkthrough (`main.go`)

This lesson's actual file is deliberately valid and compiles cleanly — it exists as a "known
good" baseline. The errors above are things to **try yourself** in a scratch copy, not things
present in the committed file.

## ▶️ How to Run

```bash
cd level-00-getting-started/36-compile-errors
go run main.go
```

## ✅ Expected Output

```
10 / 2 = 5

This file compiles cleanly. See the README for the exact compiler
errors you'd get from several common, deliberate mistakes.
```

## 🧠 Key Takeaways

- Every error starts with `file:line:column` — go there first.
- Unused variables and unused imports are compile errors in Go, not warnings.
- Type mismatches are caught before your program ever runs.
- "Missing return" means the compiler can't prove every path returns, even if you can.

## 🛠️ Try It Yourself

1. Copy `main.go` to a scratch file, introduce each of the six mistakes above one at a time, and
   confirm the exact error message you see matches what's shown here.
2. Fix each one and confirm the file compiles again before moving to the next.
3. Deliberately misspell `safeDivide` at its call site and read the "undefined" error Go gives you.

## ⚠️ Common Mistakes

- Reading only the message and skipping the `file:line:column` prefix — for anything beyond a
  one-file program, the location is often the most useful part.
- Panicking at a wall of errors from one typo — fixing the **first** reported error and
  recompiling often makes several "downstream" errors disappear on their own.
