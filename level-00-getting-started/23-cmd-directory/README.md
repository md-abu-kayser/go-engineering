# 23 — cmd/ Directory

## 🎯 Learning Objectives

- Use the `cmd/` directory convention to structure a module with **multiple binaries**.
- Understand why each subdirectory of `cmd/` is its own independent `package main`.
- Recognize this pattern in real, large Go projects.

## 📖 Concept

A single Go **module** can produce more than one executable. The de facto community convention
for organizing this is a top-level `cmd/` directory, with one subdirectory per binary:

```
my-project/
├── go.mod
├── cmd/
│   ├── server/
│   │   └── main.go      # package main -> builds the "server" binary
│   └── cli/
│       └── main.go      # package main -> builds the "cli" binary
└── internal/
    └── ...                # shared logic used by both binaries
```

Each `cmd/<name>/main.go` is a **separate, independent `package main`** — they don't import each
other. Instead, shared logic typically lives in `internal/` (see
[lesson 22](../22-internal-packages)) or another package both `cmd/` entries import.

### Where this shows up in the real world

This exact pattern is used by many large, well-known Go projects — Kubernetes and Docker/Moby
both ship several binaries from one repository, each rooted under a `cmd/` subdirectory, sharing
the bulk of their logic through internal packages.

### Building and running a specific `cmd/` binary

```bash
go run ./cmd/greeter
go build -o bin/greeter ./cmd/greeter
go build ./cmd/...          # build every binary under cmd/ at once
```

## 🔍 Code Walkthrough

```
23-cmd-directory/
├── main.go              # this lesson's own top-level program
└── cmd/
    └── greeter/
        └── main.go        # a second, independent package main
```

Both files declare `package main` and both have their own `func main()` — but they are compiled
completely separately. Running `go run main.go` at the top level never touches
`cmd/greeter/main.go`, and vice versa; you choose which one to build or run by which path you
point the `go` command at.

## ▶️ How to Run

```bash
cd level-00-getting-started/23-cmd-directory
go run main.go
go run ./cmd/greeter
```

## ✅ Expected Output

```
=== cmd/ directory ===
----------------------------------
This lesson's real example lives in ./cmd/greeter — a second,
independent `package main`. Run it with:
  go run ./cmd/greeter
```

and separately, from `go run ./cmd/greeter`:

```
I'm a separate binary, built from ./cmd/greeter.
```

## 🧠 Key Takeaways

- `cmd/<name>/` is the community convention for a module that builds multiple binaries.
- Each `cmd/<name>/main.go` is its own, independent `package main`.
- Shared logic between binaries typically lives in `internal/` or another importable package.
- `go build ./cmd/...` builds every binary under `cmd/` in one command.

## 🛠️ Try It Yourself

1. Add a third binary, `cmd/farewell/main.go`, that prints a goodbye message.
2. Build all of them at once: `go build -o bin/ ./cmd/...` and check the `bin/` folder.
3. Have both `cmd/greeter` and your new `cmd/farewell` import a **shared** helper function from
   a common package (or from `internal/greeting`, reusing [lesson 22](../22-internal-packages)'s
   structure) instead of duplicating the string.

## ⚠️ Common Mistakes

- Trying to `import` one `cmd/` binary's code from another `cmd/` binary directly — `package
  main` packages generally aren't meant to be imported; put shared code in a proper library
  package instead.
- Forgetting the `./` prefix when pointing `go run`/`go build` at a subdirectory —
  `go run cmd/greeter` (no `./`) can behave differently than you expect depending on your shell
  and Go version; `./cmd/greeter` is the safe, explicit form.
