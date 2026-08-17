# 03 — Hello, World

## 🎯 Learning Objectives

- Write, save, and run your first Go program.
- Understand exactly what each line of the smallest possible Go program does.
- Understand that Go is a **compiled** language, and what that means in practice.

## 📖 Concept

Every programming language tutorial starts here for a reason: "Hello, World" is small enough
that every single line matters, which makes it the perfect place to understand the anatomy of a
program before adding any real logic.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Four lines, four concepts:

| Line                           | What it does                                                                                                                                   |
| ------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `package main`                 | Declares this file belongs to the `main` package — the one that produces a runnable program. Full details in [lesson 04](../04-package-main).  |
| `import "fmt"`                 | Brings in the standard library's formatting/printing package. Full details in [lesson 05](../05-imports).                                      |
| `func main() { ... }`          | Declares the entry point. Execution starts at the first line inside these braces.                                                              |
| `fmt.Println("Hello, World!")` | Calls the `Println` function from the `fmt` package, printing the text followed by a newline. Full details in [lesson 06](../06-fmt-printing). |

## 🔬 Compiled, not interpreted

Go is a **compiled** language. When you run `go run main.go`, Go doesn't read and execute your
source code line-by-line the way something like Python does. Instead it:

1. **Compiles** your `.go` files into a native machine-code binary.
2. **Executes** that binary.
3. Cleans up the temporary binary (when using `go run` — with `go build` the binary is kept).

This is why Go programs start fast and run fast — by the time your program executes, it's
already plain machine code, not source text being parsed on the fly.

## ▶️ How to Run

```bash
cd level-00-getting-started/03-hello-world
go run main.go
```

## ✅ Expected Output

```
Hello, World!
```

## 🧠 Key Takeaways

- Every runnable Go program needs `package main` and a `func main()`.
- `fmt.Println` is the simplest way to print text with a trailing newline.
- Go compiles to machine code before running — it does not interpret source line-by-line.

## 🛠️ Try It Yourself

1. Change the greeting to include your name.
2. Add a second `fmt.Println(...)` call and see both lines print in order.
3. Try `go build` instead of `go run`, then look at the folder — you'll see a new binary file.
   Run it directly (`./03-hello-world` on macOS/Linux, or `03-hello-world.exe` on Windows).

## ⚠️ Common Mistakes

- Forgetting the double quotes around string literals — `Println(Hello, World!)` will not compile.
- Naming the file something other than what you intend to run and then running the wrong file.
- Missing the closing brace `}` — Go's compiler errors point at line numbers, so read them
  carefully; the reported line is often _after_ the actual mistake.
