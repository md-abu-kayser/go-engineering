# 45 — Working Directory

## 🎯 Learning Objectives

- Explain what a process's "current working directory" (CWD) is.
- Read it from Go with `os.Getwd`.
- Understand that every relative path a program uses is resolved against it.

## 📖 Concept

Every running process has a **current working directory** — the directory it's "standing in,"
which the operating system tracks per-process. It's not necessarily where your program's source
code or binary lives; it's wherever the **process was started from**.

### Reading it: `os.Getwd`

```go
wd, err := os.Getwd()
```

Returns the absolute path of the process's current working directory, or an error in rare cases
(e.g. the directory was deleted out from under a running process).

### Why this matters: relative paths

Any relative path your program opens — `"README.md"`, `"./data/config.json"`,
`"../shared/file.txt"` — is resolved **against the current working directory**, not against
where your source file or compiled binary happens to live. This trips up beginners constantly:

```bash
cd level-00-getting-started/45-working-directory
go run main.go              # CWD is .../45-working-directory -> finds README.md

cd /home/claude
go run GO-ENGINEERING/level-00-getting-started/45-working-directory/main.go
# CWD is /home/claude -> "README.md" does NOT resolve to the lesson's README at all
```

Both invocations run the exact same source file — the difference in behavior comes entirely from
**where you ran the command from**, not the file itself.

### Changing it: `os.Chdir`

```go
if err := os.Chdir("/tmp"); err != nil {
    // handle error
}
```

Changes the process's working directory — affecting every relative path resolved afterward, for
the rest of that process's lifetime. Used sparingly in real programs; it's usually clearer to
build an explicit path than to change global process state.

## 🔍 Code Walkthrough (`main.go`)

```go
if _, err := os.Stat("README.md"); err == nil {
```

`os.Stat` on a bare relative path (`"README.md"`, no directory prefix) resolves it against
whatever `os.Getwd()` just reported — this is the exact mechanism the concept section above
describes, made concrete: run this from a different directory and the result changes.

## ▶️ How to Run

```bash
cd level-00-getting-started/45-working-directory
go run main.go
```

Then, to see the CWD-dependence directly:

```bash
cd /home/claude    # or any other directory
go run GO-ENGINEERING/level-00-getting-started/45-working-directory/main.go
```

## ✅ Expected Output (run from inside the lesson folder)

```
=== Working Directory ===
----------------------------------
Current working directory: /home/claude/GO-ENGINEERING/level-00-getting-started/45-working-directory
"README.md" found relative to the working directory above.
```

## 🧠 Key Takeaways

- The current working directory is a per-process, OS-tracked concept — `os.Getwd()` reads it.
- Relative paths are always resolved against the CWD, **not** against your source file's location.
- `os.Chdir` changes it, affecting every relative path resolved afterward.
- Where you **run a command from** can change a program's behavior even with identical code.

## 🛠️ Try It Yourself

1. Run `go run main.go` from inside this lesson's folder, then again from the repository root —
   compare the two outputs.
2. Add an `os.Chdir` call at the top of `main` that changes into this lesson's own folder using an
   absolute path, so the `README.md` check succeeds regardless of where the program was launched
   from.
3. Print `os.Args[0]` (see [lesson 44](../44-command-arguments)) alongside `os.Getwd()` and
   confirm they're two genuinely different pieces of information — one is where the binary lives,
   the other is where the process was started from.

## ⚠️ Common Mistakes

- Hardcoding a relative path in a program and assuming it will always resolve correctly,
  regardless of where the program gets run from in practice (a cron job, a systemd service, a
  different developer's shell).
- Confusing "the directory the source file is in" with "the current working directory" — they are
  frequently different, and Go provides no automatic way to find the former without extra work
  (see [`runtime.Caller`] for that unrelated, rarely-needed use case).
