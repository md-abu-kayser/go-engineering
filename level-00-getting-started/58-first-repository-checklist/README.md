# 58 — First Repository Checklist

## 🎯 Learning Objectives

- Pull together everything from this level into one practical, repeatable checklist.
- Know exactly what to check before calling a new Go repository "done" for a first release.
- Recognize which checks can be automated (and should run in CI) versus which need a human.

## 📖 Concept

This is a **capstone lesson** — no new Go syntax, just a checklist tying together lessons 01–57
into the sequence you'd actually run through when starting (or reviewing) a real Go repository.

### The checklist

**1. Module & structure**

- [ ] `go mod init <appropriate-module-path>` has been run ([lesson 16](../16-go-mod-init)) —
  and the path is right for how this will actually be fetched, not just a placeholder.
- [ ] Source layout matches the project's actual size — flat is fine for something small
  ([lesson 20](../20-source-layout)); `internal/`/`cmd/`/`pkg/` only where they earn their keep
  ([lessons 22–24](../22-internal-packages)).
- [ ] No accidental nested `go.mod` unless a real multi-module workspace is intended
  ([lesson 21](../21-workspace-layout)).

**2. Hygiene: formatting, vetting, testing**

- [ ] `gofmt -l .` reports nothing ([lesson 11](../11-go-fmt-command)).
- [ ] `go vet ./...` reports nothing ([lesson 12](../12-go-vet-command)).
- [ ] `go build ./...` succeeds cleanly.
- [ ] `go test ./...` passes, with meaningful table-driven tests where it counts
  ([lesson 10](../10-go-test-command)).
- [ ] `go mod tidy` has been run, and `go.mod`/`go.sum` are committed and up to date
  ([lesson 17](../17-go-mod-tidy)).

**3. Documentation**

- [ ] A root `README.md` explains what the project is, how to install it, and how to use it.
- [ ] Every exported symbol has a proper doc comment
  ([lessons 13](../13-go-doc-command), [27](../27-documentation-comments),
  [28](../28-godoc-style)).
- [ ] `go doc -all ./...` produces something you'd be comfortable handing to a new contributor.

**4. Files every repository needs**

- [ ] A `LICENSE` file, with a license you've actually chosen deliberately.
- [ ] A `.gitignore` that excludes build artifacts, `go.work` (unless intentionally shared —
  see [lesson 21](../21-workspace-layout)), and anything OS/editor-specific.
- [ ] `.gitattributes`, if the project has contributors on multiple operating systems
  ([lesson 52](../52-windows-line-endings)).

**5. Build & release readiness**

- [ ] The project builds cleanly for every platform it needs to support
  ([lesson 48](../48-cross-compilation-overview)).
- [ ] Release binaries (if any) follow the `<name>-<os>-<arch>` naming convention
  ([lesson 54](../54-binary-naming)).
- [ ] Version/commit metadata is wired up via `-ldflags -X`, if the project is a CLI tool
  users will run `--version` against ([lesson 55](../55-build-metadata)).
- [ ] Builds use `-trimpath` for reproducibility, where that matters
  ([lesson 53](../53-reproducible-build-basics)).

**6. Error handling & exit behavior**

- [ ] Errors are wrapped with useful context (`%w`) and checked with `errors.Is`/`errors.As`
  ([lesson 38](../38-reading-error-messages)).
- [ ] Panics are only recovered at deliberate boundaries, not swallowed everywhere
  ([lesson 37](../37-runtime-panics)).
- [ ] The program uses meaningful exit codes and writes diagnostics to stderr, not stdout
  ([lessons 39](../39-exit-status), [41](../41-standard-output), [42](../42-standard-error)).

### What's automatable vs. what needs a human

Everything in section 2 is trivially scriptable — this is exactly what a CI pipeline's first job
should run, on every single commit:

```bash
gofmt -l . && \
go vet ./... && \
go build ./... && \
go test ./...
```

Sections 3–6 mostly need human judgment — no tool can tell you whether your README is actually
*clear*, or whether your chosen module path is the right one for your project's future.

## 🔍 Code Walkthrough (`main.go`)

This lesson's `main.go` mechanically checks just two items from the list (that `README.md` and
`main.go` exist in the current folder) — a small, honest illustration that **some** of a
checklist can be automated, while making clear that most of it genuinely can't be, from inside a
single program.

## ▶️ How to Run

```bash
cd level-00-getting-started/58-first-repository-checklist
go run main.go
echo "exit status: $?"
```

## ✅ Expected Output

```
=== First Repository Checklist (partial, automated check) ===
----------------------------------
✅ README.md exists in this folder
✅ main.go exists in this folder

The REST of the checklist (LICENSE, .gitignore, gofmt/vet/test cleanliness,
a clear README structure) isn't mechanically checkable from inside one lesson
folder — see the README for the full list to apply at the repository root.
exit status: 0
```

## 🧠 Key Takeaways

- Formatting, vetting, building, and testing are fully automatable — wire them into CI on day one.
- Documentation, licensing, and module-path choices need human judgment, not just tooling.
- This repository (`GO-ENGINEERING` itself) is a working example of applying this exact checklist
  — its own root `README.md`, `LICENSE`, `.gitignore`, and per-lesson structure follow it.
- A good checklist isn't about perfection on day one — it's about not forgetting something
  important under deadline pressure.

## 🛠️ Try It Yourself

1. Run the automatable command block above (`gofmt`/`go vet`/`go build`/`go test`) against the
   **entire** `GO-ENGINEERING` repository from its root, and confirm it's still clean after 58
   lessons.
2. Pick any one item from sections 3–6 above that this repository is missing or could improve,
   and actually do it.
3. Sketch (even just as a comment or a scratch file) what a minimal GitHub Actions CI workflow
   running section 2's checklist on every push would look like.

## ⚠️ Common Mistakes

- Treating a checklist like this as a one-time setup task instead of something to re-run —
  `gofmt`/`go vet`/`go test` cleanliness should be continuously enforced (via CI), not just
  checked once at project creation.
- Skipping the "human judgment" sections because they're harder to verify mechanically — a
  perfectly `gofmt`-clean repository with a confusing README or no license is still not genuinely
  ready to share.
