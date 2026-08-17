# 05 — Imports

## 🎯 Learning Objectives

- Write single and grouped `import` statements correctly.
- Use an **aliased** import to rename a package locally.
- Understand what a **blank import** (`_`) does and when it's appropriate.
- Understand why Go treats unused imports as a **compile error**, not a warning.

## 📖 Concept

### Single import

```go
import "fmt"
```

### Grouped import (idiomatic for 2+ packages)

```go
import (
    "fmt"
    "math"
)
```

`gofmt` automatically sorts grouped imports alphabetically and keeps standard-library imports
separated from third-party ones (when there are any), so you rarely need to think about ordering
by hand.

### Aliased import

```go
import m "math"
```

This lets you refer to the package as `m` instead of `math` everywhere in the file. Useful when:

- Two imported packages would otherwise have the same name.
- The package name is long and you want a shorter local name.
- You're intentionally shadowing a name for clarity (used sparingly).

### Blank import

```go
import _ "time/tzdata"
```

The underscore `_` means: _"import this package purely for its side effects — its `init()`
function — and don't give me a name to call it by."_ Without the `_`, Go would refuse to compile
the file, because you'd have an imported package you never reference (see below). Common real
uses:

- `_ "time/tzdata"` — embeds time zone data into the binary.
- Database drivers, e.g. `_ "github.com/lib/pq"` — registers a driver with `database/sql` without
  needing to call anything from the package directly.

## 🔍 Why unused imports don't compile

Try this thought experiment: comment out every line that uses `m.Pi` and `m.Sqrt(2)` in
`main.go`, but leave `m "math"` in the import block. Run `go run main.go` and you'll get:

```
./main.go:7:2: "math" imported and not used
```

Go enforces this **at compile time**, not as a linter warning, because unused imports are almost
always leftover clutter from editing — Go's philosophy is to keep code exactly as clean as it
needs to be, no more, no less.

## ▶️ How to Run

```bash
cd level-00-getting-started/05-imports
go run main.go
```

## ✅ Expected Output

```
=== Import Styles ===
----------------------------------
Pi (via aliased import `m "math"`) : 3.14159
Square root of 2 : 1.41421

The blank import `_ "time/tzdata"` above contributes nothing we call
directly. It runs only for its side effect: it embeds the IANA time zone
database into this binary, so time.LoadLocation() works even on a machine
that has no system time zone data installed.
```

## 🧠 Key Takeaways

- Grouped imports are idiomatic for anything beyond a single import.
- Aliased imports rename a package locally with `alias "path"`.
- Blank imports (`_ "path"`) run a package's `init()` for its side effects only.
- Unused imports are a **compile error** in Go, not a warning.

## 🛠️ Try It Yourself

1. Remove the alias `m` and change every `m.` to `math.` — confirm the program still works
   identically.
2. Deliberately import a package you don't use (e.g. `"strings"`) and observe the compiler error.
3. Look up one real-world blank import in a popular Go project on GitHub (search for
   `import _ "` in a repository) and read what side effect it relies on.

## ⚠️ Common Mistakes

- Forgetting the alias name goes **before** the path: `m "math"`, not `"math" m`.
- Using a blank import when you actually meant to call something from the package — if you find
  yourself blank-importing and then wondering how to use it, you probably wanted a normal import.
