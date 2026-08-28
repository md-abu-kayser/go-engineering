# 57 — Checksum Database Overview

## 🎯 Learning Objectives

- Explain what `go.sum` actually records, and what specific risk it protects against.
- Understand the role of the public checksum database (`sum.golang.org`).
- Know when and why you'd configure `GOPRIVATE` or `GONOSUMCHECK`.

## 📖 Concept

[Lesson 17](../17-go-mod-tidy) mentioned `go.sum` briefly as "a checksum ledger, never hand-
edited." This lesson goes deeper into **why** it exists and how it's independently verified.

### What `go.sum` actually protects against

Every line in `go.sum` looks roughly like:

```
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
```

That hash is a cryptographic checksum of that **exact** module version's content. When you build
your project, Go re-computes the hash of whatever it actually downloaded and compares it against
this recorded value. If they don't match, the build **fails immediately** — this is what stops a
compromised or tampered module (whether from a hacked package registry, a man-in-the-middle
network attacker, or a malicious mirror) from silently substituting different code for a version
you thought you trusted.

### The independent check: `sum.golang.org`

`go.sum` alone only proves "this matches what *I* recorded before." The **checksum database**
(`sum.golang.org`, a public, append-only, cryptographically-verifiable log run by Google) adds a
second, independent layer: when you first fetch a new module version, Go checks its hash against
this public database too — so even a compromised `go.sum` file **committed to your own repo**
(say, via a compromised contributor's account) wouldn't be enough on its own; the module's hash
also has to match what the wider world has independently observed for that exact version.

```bash
go env GOSUMDB
# -> sum.golang.org (the default)
```

### When you'd turn this off: `GOPRIVATE` and `GONOSUMCHECK`

Private, internal modules that will never be published publicly obviously can't be verified
against a public database — nobody outside your organization has ever fetched them to compare
against. For these:

```bash
go env -w GOPRIVATE=github.com/mycompany/*
```

`GOPRIVATE` tells Go: modules matching this pattern are private — skip the public checksum
database check for them (and also skip routing their downloads through the public module proxy,
since a private repository's code shouldn't transit a public proxy either).

## 🔍 Code Walkthrough (`main.go`)

```go
out, err := exec.Command("go", "env", "GOSUMDB", "GOPRIVATE", "GONOSUMCHECK").Output()
```

Same principle as [lesson 56](../56-module-cache-overview): ask the toolchain directly for its
own effective configuration, rather than guessing defaults.

## ▶️ How to Run

```bash
cd level-00-getting-started/57-checksum-database-overview
go run main.go
```

## ✅ Expected Output (shape)

```
=== Checksum Database Overview ===
----------------------------------
GOSUMDB       : sum.golang.org
GOPRIVATE     :
GONOSUMCHECK  :

This repository has no go.sum yet, since it has zero external
dependencies. See the README for what go.sum records once it does,
and how GOSUMDB independently verifies it.
```

(`GOPRIVATE` and `GONOSUMCHECK` are commonly empty by default — both are opt-in, and
`GONOSUMCHECK` in particular is a legacy variable predating `GOSUMDB`/`GONOSUMDB` that most
modern setups leave unset entirely.)

## 🧠 Key Takeaways

- Each `go.sum` line is a cryptographic checksum for one exact module version.
- Go re-verifies this checksum on every build — a mismatch fails the build immediately.
- `sum.golang.org` is a second, independent, public verification layer beyond your own `go.sum`.
- `GOPRIVATE` opts specific module path patterns (typically your own org's) out of the public
  checksum database, appropriately, for genuinely private code.

## 🛠️ Try It Yourself

1. Run `go env GOSUMDB` yourself and confirm the default (`sum.golang.org`) on your machine.
2. In a scratch module with a real dependency, open its `go.sum` and identify the module name,
   version, and checksum on one line.
3. Read Go's own documentation on the checksum database (search "Go checksum database") for how
   it's structured as a cryptographically verifiable transparency log, similar in spirit to
   Certificate Transparency for TLS certificates.

## ⚠️ Common Mistakes

- Setting `GOPRIVATE` overly broadly (e.g. accidentally matching public modules too), which
  disables real security verification for things that didn't need it disabled.
- Manually editing a `go.sum` entry to "fix" a checksum mismatch instead of investigating **why**
  it mismatched — a mismatch is either a genuine security concern or a sign something else is
  misconfigured; papering over it defeats the entire mechanism's purpose.
