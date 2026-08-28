# 50 — File Permissions

## 🎯 Learning Objectives

- Read the Unix permission model: owner / group / other, each with read / write / execute.
- Create a file with specific permission bits via `os.WriteFile`.
- Inspect a file's mode with `os.Stat`, and change it with `os.Chmod`.

## 📖 Concept

On Unix-like systems (Linux, macOS), every file has a set of **permission bits** controlling who
can read, write, or execute it — expressed as three groups of three bits: **owner**, **group**,
and **other**, each with **r**ead, **w**rite, and e**x**ecute.

### Octal notation

```
0644
│└┴┴─ owner=6 (rw-), group=4 (r--), other=4 (r--)
└──── leading 0 marks this as an octal literal in Go source
```

Each digit is a sum: read=4, write=2, execute=1. So `6` = read+write (4+2), `4` = read-only,
`7` = read+write+execute. `0644` is the extremely common "owner can edit, everyone else can only
read" pattern for ordinary files; `0755` is the equivalent for executables/directories (everyone
gets execute/traverse permission too).

### Creating a file with specific permissions

```go
os.WriteFile(path, []byte("hello"), 0644)
```

The third argument is an `os.FileMode` — Go represents it as a plain integer type, which is why
octal literals (`0644`) are the natural way to write one.

### Reading a file's mode back

```go
info, err := os.Stat(path)
info.Mode()       // the full mode, e.g. "-rw-r--r--"
info.Mode().Perm() // just the permission bits, e.g. 0644
```

`Mode()` includes extra bits beyond plain permissions (like whether it's a directory), which is
why `.Perm()` exists specifically to isolate the permission portion.

### Changing permissions: `os.Chmod`

```go
os.Chmod(path, 0600)
```

Same octal-literal style as creation — `0600` means "owner can read+write, nobody else can do
anything," a common tightening for files containing secrets (credentials, private keys).

## 🔍 Code Walkthrough (`main.go`)

```go
defer os.Remove(path)
```

Cleans up the demo file when `main` exits, regardless of how it exits — a small, real-world
example of `defer` for cleanup, distinct from the panic-recovery use in
[lesson 37](../37-runtime-panics).

```go
fmt.Printf("Created with mode : %s (%#o)\n", info.Mode(), info.Mode().Perm())
```

`%s` on an `os.FileMode` prints the familiar `-rw-r--r--` style string; `%#o` prints it as a
`0`-prefixed octal number — both represent the exact same bits, just formatted differently.

## ▶️ How to Run

```bash
cd level-00-getting-started/50-file-permissions
go run main.go
```

## ✅ Expected Output

```
=== File Permissions ===
----------------------------------
Created with mode : -rw-r--r-- (0644)
After os.Chmod    : -rw------- (0600)
Is a directory?   : false
```

## 🧠 Key Takeaways

- Unix permissions are three groups (owner/group/other) of three bits (read/write/execute) each.
- Octal literals (`0644`, `0600`, `0755`) are the idiomatic way to express `os.FileMode` values.
- `os.Stat(path).Mode()` reads the current mode; `.Perm()` isolates just the permission bits.
- `os.Chmod` changes permissions on an existing file.

## 🛠️ Try It Yourself

1. Change the creation mode to `0600` directly (instead of `0644` then `Chmod`) and confirm the
   printed mode matches immediately.
2. Try `os.Chmod(path, 0000)` (no permissions for anyone) and then attempt to read the file with
   `os.ReadFile` — observe the permission-denied error, then restore permissions before deleting.
3. Create a file with `0755` instead, and compare its printed mode string
   (`-rwxr-xr-x`) to `0644`'s (`-rw-r--r--`) — note the added `x` bits.

## ⚠️ Common Mistakes

- Forgetting the leading `0` on an octal literal (`644` instead of `0644`) — in Go, `644` without
  the `0` is the **decimal** number 644, an entirely different (and invalid) permission value.
- Assuming this permission model applies identically on Windows — Windows has a fundamentally
  different, ACL-based permission system; Go's `os.FileMode` bits are a best-effort approximation
  there, not a literal match.
