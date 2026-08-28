# 51 — Unix Line Endings

## 🎯 Learning Objectives

- Explain the difference between LF (`\n`) and CRLF (`\r\n`) line endings.
- See both represented literally as bytes inside a Go string.
- Confirm `bufio.Scanner` handles both transparently, without extra code.

## 📖 Concept

Different operating systems historically settled on different conventions for "where one line of
text ends and the next begins":

| Convention | Bytes | Used by |
|---|---|---|
| **LF** (Line Feed) | `\n` (one byte) | Linux, macOS (modern), and virtually everything else |
| **CRLF** (Carriage Return + Line Feed) | `\r\n` (two bytes) | Windows |

This isn't just trivia — a text file with the "wrong" line ending for a given tool can cause
subtle bugs: a shebang line that mysteriously doesn't work, a string comparison that fails for no
visible reason (because there's an invisible `\r` at the end), or a diff that shows every single
line as changed when only whitespace differs.

### Seeing the difference literally

```go
unixLine := "hello\n"       // 6 bytes: h e l l o \n
windowsLine := "hello\r\n"   // 7 bytes: h e l l o \r \n
```

`%q` formatting makes the invisible characters visible: `"hello\n"` prints exactly as written,
with the escape sequence shown rather than an actual newline — useful for spotting a stray `\r`
that would otherwise be invisible in a normal terminal.

### `bufio.Scanner` handles both, transparently

[Lesson 40](../40-standard-input) introduced `bufio.Scanner` for reading lines. Its default
split function (`bufio.ScanLines`) specifically strips a trailing `\r\n` **or** a bare `\n` —
you get clean line content either way, with no manual `strings.TrimRight(line, "\r\n")` needed.

```go
scanner := bufio.NewScanner(strings.NewReader(mixed))
for scanner.Scan() {
    line := scanner.Text() // never has a trailing \r or \n, regardless of source
}
```

### Where this bites in practice

- **Git**: cloning a Windows-authored repo on Linux (or vice versa) can introduce mixed line
  endings unless `.gitattributes` or `core.autocrlf` is configured — worth knowing exists, even
  if not a Go-specific concern.
- **Hand-rolled line splitting**: code that does `strings.Split(text, "\n")` on Windows-authored
  text ends up with a stray `\r` at the end of every line — `bufio.Scanner` avoids this
  automatically; manual splitting doesn't.

## 🔍 Code Walkthrough (`main.go`)

```go
mixed := "unix line\nwindows line\r\nanother unix line\n"
```

This string deliberately mixes both conventions in one input — exactly the situation you might
hit reading a file that was edited on more than one OS over its history — to prove
`bufio.Scanner` doesn't care which one it encounters.

## ▶️ How to Run

```bash
cd level-00-getting-started/51-unix-line-endings
go run main.go
```

## ✅ Expected Output

```
=== Unix Line Endings ===
----------------------------------
Unix line    : "hello\n" (len 6)
Windows line : "hello\r\n" (len 7)

Scanning mixed line endings with bufio.Scanner:
  "unix line" (no leftover \r or \n)
  "windows line" (no leftover \r or \n)
  "another unix line" (no leftover \r or \n)
```

## 🧠 Key Takeaways

- LF (`\n`) is one byte; CRLF (`\r\n`) is two — the extra `\r` is easy to miss visually.
- `bufio.Scanner`'s default split strips either ending automatically — prefer it over manual
  `strings.Split(text, "\n")` when reading text that might have either.
- A stray, invisible `\r` at the end of a manually-split line is a classic source of "why doesn't
  this string comparison work" bugs.

## 🛠️ Try It Yourself

1. Change `mixed` to end with `\r\n` instead of `\n` on the last line and confirm the scanner
   still strips it correctly.
2. Replace `bufio.Scanner` with a hand-rolled `strings.Split(mixed, "\n")` and print each result
   with `%q` — spot the leftover `\r` on the Windows-originated line that `Scanner` was hiding
   from you.
3. Look up your own text editor or IDE's line-ending setting (VS Code shows it in the status bar)
   and confirm which convention it's currently using for files you edit.

## ⚠️ Common Mistakes

- Manually splitting on `"\n"` alone and being surprised when a comparison against a Windows-
  originated line silently fails — the invisible trailing `\r` is almost always the cause.
- Assuming all text on a Unix-like system is guaranteed to be LF-only — files can (and often do)
  arrive from Windows-authored sources regardless of what OS you're currently running on.
