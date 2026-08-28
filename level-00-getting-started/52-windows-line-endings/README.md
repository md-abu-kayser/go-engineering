# 52 — Windows Line Endings

## 🎯 Learning Objectives

- Confirm Go performs **no automatic** line-ending translation when reading or writing files.
- Deliberately construct CRLF content and measure the byte-size difference versus LF.
- Understand how Git's `core.autocrlf` and `.gitattributes` interact with cross-platform repos.

## 📖 Concept

[Lesson 51](../51-unix-line-endings) looked at LF vs CRLF from the **reading** side —
`bufio.Scanner` transparently handles either. This lesson looks at the **writing** side, and at
the tooling layer above Go entirely: version control.

### Go never translates line endings for you

```go
os.WriteFile(path, []byte("hello\n"), 0644)
```

Whatever bytes you pass are written **exactly as given** — no "text mode" that silently
substitutes `\n` for the platform's native convention, the way some older languages' I/O layers
worked. If you want a file with CRLF endings, you build that string yourself
(`strings.Join(lines, "\r\n")`, as this lesson's `main.go` does) — Go won't do it implicitly, on
any operating system.

### Where the confusion usually enters: Git

Because Go itself is neutral about line endings, the place cross-platform line-ending problems
usually actually originate is **version control**, not the Go program itself:

- **`core.autocrlf`** (a per-developer Git setting) can automatically convert LF to CRLF on
  checkout (Windows) and back to LF on commit — convenient for editing text with
  Windows-native tools, but a common source of "every line shows as changed" diffs when
  teammates have it configured differently.
- **`.gitattributes`** (a per-**repository** file, checked into the repo itself) lets you pin
  down line-ending behavior explicitly and consistently for everyone, regardless of individual
  `core.autocrlf` settings:

  ```
  * text=auto
  *.go text eol=lf
  *.bat text eol=crlf
  ```

  This says: normalize text files generally, but force `.go` files to always use LF (matching Go
  community convention) and `.bat` files to always use CRLF (matching what Windows batch files
  expect).

### Why Go source specifically standardizes on LF

`gofmt` ([lesson 11](../11-go-fmt-command)) and the wider Go community convention use **LF**
line endings for `.go` files, regardless of the OS they're edited on — this repository's own
`.gitignore`/`.gitattributes` setup (if you add one, per the exercise below) should reflect that.

## 🔍 Code Walkthrough (`main.go`)

```go
crlfContent := strings.Join(lines, "\r\n") + "\r\n"
```

Building CRLF content **explicitly** like this is the point — it makes visible exactly what
"CRLF line endings" means at the byte level, rather than treating it as some invisible file
property.

## ▶️ How to Run

```bash
cd level-00-getting-started/52-windows-line-endings
go run main.go
```

## ✅ Expected Output

```
=== Windows Line Endings ===
----------------------------------
Bytes written (CRLF)   : 37
Bytes if it were LF    : 34
Contains \r\n?          : true

See the README for how Git's core.autocrlf and .gitattributes interact
with files like this one across contributors on different operating systems.
```

## 🧠 Key Takeaways

- Go writes exactly the bytes you give it — no implicit line-ending translation, ever.
- CRLF-vs-LF confusion in real projects usually originates in Git configuration, not in Go code.
- `.gitattributes` (repo-level, committed) is more reliable than `core.autocrlf` (per-developer)
  for keeping line endings consistent across a team.
- Go source files conventionally use LF, matching `gofmt`'s own output.

## 🛠️ Try It Yourself

1. Confirm the byte-count math yourself: three lines plus a trailing terminator, CRLF (2 bytes)
   vs LF (1 byte) per line-ending — check your arithmetic against the program's own output (37
   vs 34 bytes).
2. Create a `.gitattributes` file (in a scratch repo, not necessarily this one) with the content
   shown above, and read Git's documentation on what `text=auto` actually detects.
3. Check your own Git configuration with `git config --get core.autocrlf` and consider whether
   its current value matches what a shared `.gitattributes` file would enforce anyway.

## ⚠️ Common Mistakes

- Blaming Go for line-ending inconsistencies that actually originate in Git configuration or in a
  text editor's save settings — Go itself is a completely neutral party here.
- Relying on every teammate having the same `core.autocrlf` setting instead of committing a
  `.gitattributes` file that enforces it project-wide, regardless of individual configuration.
