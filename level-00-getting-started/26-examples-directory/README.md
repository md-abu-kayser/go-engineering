# 26 — examples/ Directory

## 🎯 Learning Objectives

- Understand why libraries maintain a runnable `examples/` directory.
- Distinguish "usage examples" from both tests and "godoc example functions".
- Write a small, self-contained usage example for a package you build.

## 📖 Concept

When you publish a reusable library, documentation alone (see
[lesson 13](../13-go-doc-command)) often isn't enough — people learn fastest from **complete,
runnable code** they can copy, run, and modify. That's what an `examples/` directory is for:

```
26-examples-directory/
├── greetlib/
│   └── greetlib.go        # the library itself
└── examples/
    └── basic/
        └── main.go           # a runnable demonstration of how to use it
```

Each subdirectory of `examples/` is typically its own small `package main` — a complete,
standalone program that imports the library and shows one realistic usage pattern, exactly the
way a real consumer of the library would.

### `examples/` vs. tests vs. "godoc Example functions"

It's worth being precise about three related-but-different things:

| Mechanism | Purpose | Runs as part of `go test`? |
|---|---|---|
| `examples/<name>/main.go` (this lesson) | A full, runnable demo program for humans to read and copy | ❌ No — run manually with `go run` |
| `_test.go` files ([lesson 10](../10-go-test-command)) | Automated correctness checks | ✅ Yes |
| `func ExampleXxx()` inside a `_test.go` file | A doc-comment-adjacent example whose **output is checked** by `go test`, and which also renders on pkg.go.dev | ✅ Yes |

All three are valuable and often coexist in a mature library — this lesson focuses on the first,
since it's the simplest and most beginner-friendly to start with.

## 🔍 Code Walkthrough

```go
// go-engineering/level-00-getting-started/26-examples-directory/greetlib/greetlib.go
type Greeting struct {
    Prefix string
}

func (g Greeting) For(name string) string { ... }
```

`greetlib` is the library — small, exported, with a doc comment following the convention from
[lesson 13](../13-go-doc-command).

```go
// go-engineering/level-00-getting-started/26-examples-directory/examples/basic/main.go
defaultGreeting := greetlib.Greeting{}
fmt.Println(defaultGreeting.For("Gopher"))

custom := greetlib.Greeting{Prefix: "Welcome"}
fmt.Println(custom.For("Gopher"))
```

This is exactly the code a new user of `greetlib` would want to read first: the zero value in
use, then a customized instance — both in a few lines, with no test framework or extra ceremony
in the way.

## ▶️ How to Run

```bash
cd level-00-getting-started/26-examples-directory
go run main.go
go run ./examples/basic
```

## ✅ Expected Output

```
=== examples/ directory ===
----------------------------------
This lesson's real example lives in ./examples/basic — a small,
standalone program showing how to use ./greetlib. Run it with:
  go run ./examples/basic
```

and separately, from `go run ./examples/basic`:

```
Hello, Gopher!
Welcome, Gopher!
```

## 🧠 Key Takeaways

- `examples/<name>/main.go` gives new users of a library complete, runnable, copy-pasteable code.
- It's distinct from `_test.go` files and from `func ExampleXxx()` godoc examples — all three
  serve different, complementary purposes.
- Good examples are short, realistic, and show more than one usage pattern (default vs. customized).

## 🛠️ Try It Yourself

1. Add a second example, `examples/uppercase/main.go`, that uses `greetlib.Greeting` together
   with `strings.ToUpper` to shout the greeting.
2. Read about `func ExampleXxx()` (godoc example functions) and try converting
   `examples/basic/main.go`'s logic into one, inside a `_test.go` file, with an `// Output:`
   comment Go verifies automatically.
3. Find a real popular Go library on GitHub with an `examples/` directory and skim one of its
   examples.

## ⚠️ Common Mistakes

- Letting examples rot out of sync with the library's actual API — since `examples/` programs
  aren't run by `go test` automatically (unlike `ExampleXxx` functions), they need a habit of
  being run manually after changes, or converted to godoc examples for automatic verification.
- Writing examples that are too clever or too abstract — the best examples are boring and literal,
  optimized for someone copy-pasting them as a starting point.
