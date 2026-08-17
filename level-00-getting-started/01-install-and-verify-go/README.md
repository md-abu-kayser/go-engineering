# 01 — Install & Verify Go

## 🎯 Learning Objectives

By the end of this lesson you will be able to:

- Confirm Go is installed correctly using both the command line and a Go program.
- Explain what the `runtime` package is and why it's useful.
- Read basic Go syntax: package declaration, imports, structs, functions, `main()`.
- Run a Go file and a Go test file.

## 📖 Concept

Most tutorials tell you to check your installation by running:

```bash
go version
```

That's correct, but it only tells you what the **command-line tool** reports. This lesson goes
one step further and asks the **Go runtime itself**, from inside a running program, using the
standard library's [`runtime`](https://pkg.go.dev/runtime) package. This is a gentle way to see:

- Your first real Go **struct** (`envInfo`) — a way to group related data together.
- Your first **function that returns a value** (`collectEnvInfo`).
- How `main()` calls other functions instead of doing everything itself.

## 🔍 Code Walkthrough (`main.go`)

```go
type envInfo struct {
    GoVersion string
    OS        string
    Arch      string
    NumCPU    int
    Compiler  string
}
```

A `struct` is Go's way of bundling related fields into one named type — similar to an object's
shape in TypeScript or a plain class in other languages, but without methods baked in by default.

```go
func collectEnvInfo() envInfo {
    return envInfo{...}
}
```

`collectEnvInfo` is a function with **no parameters** that **returns an `envInfo` value**. It's
kept separate from `main()` on purpose — separating "get the data" from "display the data" is a
habit you'll want throughout your career. It also means this function can be **tested in
isolation**, which is exactly what `01_install_and_verify_go_test.go` does.

```go
func main() {
    info := collectEnvInfo()
    fmt.Println(...)
}
```

`main()` is the entry point (more on that in [lesson 04](../04-package-main)). Here it just calls
`collectEnvInfo()` and prints the result.

## ▶️ How to Run

```bash
cd level-00-getting-started/01-install-and-verify-go
go run main.go
```

## ✅ Expected Output

Your exact values will differ based on your machine, but the shape will look like this:

```
✅ Go is installed and working!
----------------------------------
Go version : go1.22.3
Operating system : linux
Architecture : amd64
Logical CPUs : 8
Compiler : gc
```

## 🧪 Run the Test

```bash
go test -v ./...
```

You should see `PASS` for `TestCollectEnvInfo`. Tests like this are how you'll verify your code
does what you think it does — without manually re-running and eyeballing output every time.

## 🧠 Key Takeaways

- `go version` checks the CLI tool; `runtime.Version()` checks it from **inside** a program.
- Structs group related data under one type.
- Separating "compute" from "print" makes code easier to test and reuse.
- Test files (`_test.go`) live next to the code they test, in the same package.

## 🛠️ Try It Yourself

1. Add a new field to `envInfo` called `NumGoroutine` using `runtime.NumGoroutine()`.
2. Print it in `main()`.
3. Add an assertion for it in the test file.

## ⚠️ Common Mistakes

- **"go: command not found"** — Go isn't installed or isn't on your `PATH`. Reinstall from
  [go.dev/dl](https://go.dev/dl/) and restart your terminal.
- Forgetting `go run main.go` needs to be run **from inside the lesson folder** (or you must
  point it at the right file path).
