# Level 00 — Getting Started

Welcome to the very first module of **GO-ENGINEERING**. This level assumes **zero prior Go
experience** and focuses entirely on getting your environment right, understanding the absolute
fundamentals, and becoming fluent with the Go toolchain and project conventions before any real
application logic is introduced.

By the end of this level you will be able to:

- Verify that Go is installed correctly and understand what the toolchain is doing.
- Explain what `package main`, `import`, and `func main()` actually mean.
- Write, format, and run a Go program from the command line and from VS Code.
- Use `fmt` confidently to print and format output.
- Use every core `go` subcommand: `build`, `install`, `test`, `fmt`, `vet`, `doc`, `env`, `list`.
- Initialize and maintain a module with `go mod init` / `go mod tidy`, and explain how import
  paths and semantic import versioning work.
- Structure a Go project using the community's real conventions: `internal/`, `cmd/`, `pkg/`,
  `api/`, `examples/`.
- Write proper documentation comments and navigate a codebase productively in VS Code using
  `gopls`.
- Debug confidently with Delve — breakpoints, stepping, variable inspection, and stack traces —
  both from the command line and through VS Code.
- Read compiler errors and runtime panics correctly, and handle Go's `error` values idiomatically.
- Work with standard input, output, and error streams the way real command-line tools do.
- Read and set environment variables, handle command-line arguments and flags, and reason
  correctly about working directories and relative vs. absolute paths.
- Cross-compile for any platform, understand build tags and file permissions, and produce
  reproducible, properly-named, version-stamped release binaries.
- Explain what Go's module cache and checksum database actually protect against.
- Apply a complete, practical checklist before calling a new Go repository "done."

## Lessons

**Part 1 — Language basics**

| #   | Folder                        | Topic                  |
| --- | ----------------------------- | ---------------------- |
| 01  | `01-install-and-verify-go`    | Install & Verify Go    |
| 02  | `02-go-version-and-toolchain` | Go Version & Toolchain |
| 03  | `03-hello-world`              | Hello, World           |
| 04  | `04-package-main`             | package main           |
| 05  | `05-imports`                  | Imports                |
| 06  | `06-fmt-printing`             | fmt & Printing         |
| 07  | `07-go-run`                   | go run                 |

**Part 2 — The Go command-line toolchain**

| #   | Folder               | Topic      |
| --- | -------------------- | ---------- |
| 08  | `08-go-build`        | go build   |
| 09  | `09-go-install`      | go install |
| 10  | `10-go-test-command` | go test    |
| 11  | `11-go-fmt-command`  | go fmt     |
| 12  | `12-go-vet-command`  | go vet     |
| 13  | `13-go-doc-command`  | go doc     |
| 14  | `14-go-env-command`  | go env     |
| 15  | `15-go-list-command` | go list    |

**Part 3 — Modules & dependency management**

| #   | Folder                          | Topic                      |
| --- | ------------------------------- | -------------------------- |
| 16  | `16-go-mod-init`                | go mod init                |
| 17  | `17-go-mod-tidy`                | go mod tidy                |
| 18  | `18-module-paths`               | Module Paths               |
| 19  | `19-semantic-import-versioning` | Semantic Import Versioning |

**Part 4 — Project structure conventions**

| #   | Folder                  | Topic               |
| --- | ----------------------- | ------------------- |
| 20  | `20-source-layout`      | Source Layout       |
| 21  | `21-workspace-layout`   | Workspace Layout    |
| 22  | `22-internal-packages`  | internal/ Packages  |
| 23  | `23-cmd-directory`      | cmd/ Directory      |
| 24  | `24-pkg-directory`      | pkg/ Directory      |
| 25  | `25-api-directory`      | api/ Directory      |
| 26  | `26-examples-directory` | examples/ Directory |

**Part 5 — Documentation & editor tooling**

| #   | Folder                      | Topic                     |
| --- | --------------------------- | ------------------------- |
| 27  | `27-documentation-comments` | Documentation Comments    |
| 28  | `28-godoc-style`            | godoc Style               |
| 29  | `29-editor-workflow`        | Editor Workflow (VS Code) |
| 30  | `30-gopls-overview`         | gopls Overview            |
| 31  | `31-debugging-with-delve`   | Debugging with Delve      |

**Part 6 — Debugging in depth & diagnostics**

| #   | Folder                      | Topic                  |
| --- | --------------------------- | ---------------------- |
| 32  | `32-breakpoints`            | Breakpoints            |
| 33  | `33-step-through-debugging` | Step-Through Debugging |
| 34  | `34-variable-inspection`    | Variable Inspection    |
| 35  | `35-stack-traces`           | Stack Traces           |
| 36  | `36-compile-errors`         | Compile Errors         |
| 37  | `37-runtime-panics`         | Runtime Panics         |
| 38  | `38-reading-error-messages` | Reading Error Messages |
| 39  | `39-exit-status`            | Exit Status            |

**Part 7 — Standard streams**

| #   | Folder               | Topic           |
| --- | -------------------- | --------------- |
| 40  | `40-standard-input`  | Standard Input  |
| 41  | `41-standard-output` | Standard Output |
| 42  | `42-standard-error`  | Standard Error  |

**Part 8 — Environment, paths & platform basics**

| #   | Folder                          | Topic                      |
| --- | ------------------------------- | -------------------------- |
| 43  | `43-environment-basics`         | Environment Basics         |
| 44  | `44-command-arguments`          | Command Arguments          |
| 45  | `45-working-directory`          | Working Directory          |
| 46  | `46-relative-paths`             | Relative Paths             |
| 47  | `47-absolute-paths`             | Absolute Paths             |
| 48  | `48-cross-compilation-overview` | Cross-Compilation Overview |
| 49  | `49-build-tags-overview`        | Build Tags Overview        |
| 50  | `50-file-permissions`           | File Permissions           |
| 51  | `51-unix-line-endings`          | Unix Line Endings          |

**Part 9 — Reproducibility, releases & supply chain**

| #   | Folder                          | Topic                      |
| --- | ------------------------------- | -------------------------- |
| 52  | `52-windows-line-endings`       | Windows Line Endings       |
| 53  | `53-reproducible-build-basics`  | Reproducible Build Basics  |
| 54  | `54-binary-naming`              | Binary Naming              |
| 55  | `55-build-metadata`             | Build Metadata             |
| 56  | `56-module-cache-overview`      | Module Cache Overview      |
| 57  | `57-checksum-database-overview` | Checksum Database Overview |
| 58  | `58-first-repository-checklist` | First Repository Checklist |

## How to work through this level

Go through the lessons **in order** — each one assumes you understand the ones before it.
For every lesson:

1. Read the lesson's `README.md`.
2. Open `main.go` (and any accompanying files) and read the comments before running anything.
3. Run the program (`go run main.go`) and compare the output to what the README predicts.
4. Try the "Try It Yourself" exercise at the bottom of the README before moving on.

Once you've completed all fifty-eight lessons, you're ready for Level 01 (Control Flow).
