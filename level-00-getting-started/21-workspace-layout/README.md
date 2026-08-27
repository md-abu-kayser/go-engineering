# 21 — Workspace Layout

## 🎯 Learning Objectives

- Distinguish a **module** ([lesson 16](../16-go-mod-init)) from a **workspace**.
- Understand what problem `go.work` solves.
- Know the basic `go work` commands.

## 📖 Concept

Everything so far in this repository lives inside **one module** (one `go.mod`). But real
organizations often maintain **several separate modules** — say, a shared library module and two
different service modules that both depend on it. While developing, you frequently want to make a
change in the shared library and immediately see its effect in a service, **without** publishing
a new version and bumping a `go.mod` requirement every time.

That's exactly what a **Go workspace** is for.

### Creating a workspace

```bash
go work init ./shared-library ./service-a ./service-b
```

This creates a `go.work` file like:

```
go 1.22

use (
    ./shared-library
    ./service-a
    ./service-b
)
```

With a `go.work` file present, Go tooling (build, test, `gopls` in your editor) resolves imports
**across all listed modules using their local, on-disk source** — instead of whatever version is
pinned in each module's `go.mod`. Change a function in `shared-library`, and `service-a` picks it
up immediately, with no `go mod tidy`, no version bump, no publish step.

### Key commands

```bash
go work init [modules...]   # create go.work, listing the given module directories
go work use ./another-mod    # add another module to an existing go.work
go work sync                 # sync go.work's module requirements back into each module's go.mod
```

### Important: `go.work` is a local development tool

`go.work` is almost always **excluded from version control** (this repo's `.gitignore` already
excludes it — see [lesson 07](../07-go-run)). It represents *your* local multi-module setup, not
something that should affect how the project builds for anyone else, or in CI.

## 🔍 Code Walkthrough (`main.go`)

This repository is intentionally a **single module**, so there's genuinely no `go.work` file to
show here — the lesson is entirely conceptual, which is why `main.go` just states that directly
rather than faking a multi-module example.

## ▶️ How to Run

```bash
cd level-00-getting-started/21-workspace-layout
go run main.go
```

To see `go.work` for real, try this in a scratch directory:

```bash
mkdir -p /tmp/ws-demo/lib /tmp/ws-demo/app && cd /tmp/ws-demo
(cd lib && go mod init example.com/lib)
(cd app && go mod init example.com/app)
go work init ./lib ./app
cat go.work
```

## ✅ Expected Output

```
=== Workspace Layout (go.work) ===
----------------------------------
This repository is a single module, so it has no go.work file.
See the README for when multi-module workspaces are worth using.
```

## 🧠 Key Takeaways

- A **module** (`go.mod`) is one unit of versioned code; a **workspace** (`go.work`) is a local
  grouping of multiple modules for simultaneous development.
- `go.work` makes cross-module changes visible immediately, without publishing or version bumps.
- Workspaces are a **local development convenience** — `go.work` is normally gitignored.
- Single-module projects (like this repository) simply don't need one.

## 🛠️ Try It Yourself

1. Work through the scratch-directory example above and open the generated `go.work` file.
2. Add a function to `lib`, and confirm `app` can call it immediately after `go work init`, with
   no `go get` or version change.
3. Delete `/tmp/ws-demo` once you're done experimenting.

## ⚠️ Common Mistakes

- Committing `go.work` to a shared repository, unintentionally forcing your personal local module
  layout onto every other contributor and CI.
- Reaching for a workspace when you actually just need a subpackage inside one module — a
  workspace is specifically for **multiple, separately-versioned modules**, not general code
  organization within one.
