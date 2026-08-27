# 16 — go mod init

## 🎯 Learning Objectives

- Know exactly what `go mod init` creates.
- Choose an appropriate module path for a real, publishable project.
- Understand why the module path is more than just a label.

## 📖 Concept

`go mod init <module-path>` is the very first command you run when starting a new Go project.
It creates a `go.mod` file — this repository's root `go.mod` was created with:

```bash
go mod init go-engineering
```

which produced:

```go
module go-engineering

go 1.22
```

### Choosing a module path

For a **learning repo or local project**, a short name like `go-engineering` is perfectly fine.
But for anything you intend to `go get` from elsewhere (a real open-source library, an internal
company tool fetched by teammates), the module path should be the location it will actually be
fetched from — typically a repository URL without the scheme:

```bash
go mod init github.com/md-abu-kayser/my-cool-library
```

This matters because the module path becomes the **import path** other projects use:

```go
import "github.com/md-abu-kayser/my-cool-library/somepackage"
```

Go's module system fetches code directly from version control by convention — there's no
separate "publish to a registry" step the way npm or PyPI require. The module path *is* the
fetch location.

## 🔍 Code Walkthrough (`main.go`)

This lesson's code is deliberately simple — the real content is the command and the concept, not
the program. It exists to keep this folder consistent with every other lesson in the repo.

## ▶️ How to Run

```bash
cd level-00-getting-started/16-go-mod-init
go run main.go
```

Try it for real in a scratch folder, outside this repo:

```bash
mkdir /tmp/demo-module && cd /tmp/demo-module
go mod init example.com/demo-module
cat go.mod
```

## ✅ Expected Output

```
=== go mod init ===
----------------------------------
This repository's go.mod was created with:
  go mod init go-engineering

See the README for what that command generates and why the module
path matters beyond this repo.
```

## 🧠 Key Takeaways

- `go mod init <path>` creates `go.mod`, declaring your module's name and minimum Go version.
- For local/learning projects, any short name works.
- For anything meant to be imported by others, the module path should match its real repository
  location — the path **is** the fetch address.
- You only run `go mod init` **once** per module, at the very start of a project.

## 🛠️ Try It Yourself

1. In a scratch directory, run `go mod init` with **no** module path argument and read the
   error Go gives you — it refuses to guess.
2. Run it again with a path, then open the generated `go.mod` and compare it to this repo's root
   `go.mod`.
3. Delete your scratch directory when you're done experimenting.

## ⚠️ Common Mistakes

- Running `go mod init` **inside** an existing module (i.e., a subdirectory that already has a
  `go.mod` above it) — Go will let you create a *nested* module, which is rarely what beginners
  actually want (see [lesson 18](../18-module-paths) and [lesson 21](../21-workspace-layout)).
- Picking a module path, publishing it, and then renaming it later — every consumer's import
  paths would break; choose deliberately up front.
