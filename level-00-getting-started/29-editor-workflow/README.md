# 29 — Editor Workflow (VS Code)

## 🎯 Learning Objectives

- Set up VS Code for productive, everyday Go development.
- Use format-on-save, go-to-definition, and inline diagnostics without thinking about them.
- Set a breakpoint and step through real code with the built-in debugger.

## 📖 Concept

### One-time setup

1. Install the official **Go extension** (`golang.go`) from the VS Code marketplace.
2. Open this repository's folder in VS Code. The extension will offer to install a set of
   supporting tools (`gopls`, `dlv`, `staticcheck`, and others) — accept it. This is a one-time
   step per machine.
3. Confirm format-on-save is enabled: **Settings → search "format on save"** → check the box (the
   Go extension wires this to `gofmt`/`goimports` automatically once installed).

### The everyday loop

| Action | How |
|---|---|
| Run the current file | Click the ▷ "Run" CodeLens above `func main()`, or right-click → "Run Package" |
| Run tests in this file | Click "run test" CodeLens above any `func TestXxx`, from [lesson 10](../10-go-test-command) |
| Format on save | Just save (`Ctrl/Cmd+S`) — no manual `gofmt` needed |
| Organize imports on save | Happens automatically alongside formatting, once configured |
| Jump to a definition | `F12` (or `Cmd+Click` / `Ctrl+Click`) on any identifier |
| Find all references | `Shift+F12` on any identifier |
| Rename a symbol everywhere | `F2` on any identifier — safely renames across the whole module |
| See inline errors/warnings | Just type — `gopls` reports them live, no manual `go build` needed |

### Debugging with a breakpoint

1. Open `main.go` in this folder.
2. Click in the gutter to the left of the `return amountCents / people` line inside
   `splitBill` — a red dot (breakpoint) appears.
3. Press `F5` (or the Run/Debug panel's ▷ button) and choose "Go" if prompted.
4. Execution pauses at your breakpoint. Hover over `amountCents` and `people` to see their live
   values, or check the "Variables" panel in the sidebar.
5. Press `F10` to step to the next line, or `F5` to continue running.

This uses [Delve](https://github.com/go-delve/delve), the standard Go debugger, wired up
automatically by the Go extension — no manual configuration needed for a simple case like this.

## 🔍 Code Walkthrough (`main.go`)

`splitBill` uses plain integer division on purpose — it's simple enough to predict by hand, which
makes it easy to confirm the debugger is showing you exactly what you expect at each step.

## ▶️ How to Run

```bash
cd level-00-getting-started/29-editor-workflow
go run main.go
```

Then work through the "Debugging with a breakpoint" steps above inside VS Code.

## ✅ Expected Output

```
Each person owes: $6.49

See the README for a checklist of editor features to try on this file.
```

## 🧠 Key Takeaways

- The Go extension's one-time tool install (`gopls`, `dlv`, etc.) unlocks the entire workflow
  below it.
- Format-on-save + organize-imports-on-save means you almost never run `gofmt` manually.
- `F12` / `Shift+F12` / `F2` (definition / references / rename) work across the whole module, not
  just the open file.
- `F5` + a gutter-click breakpoint gives you a full debugger with zero configuration for simple
  programs.

## 🛠️ Try It Yourself

1. Set a breakpoint inside `splitBill`, run the debugger, and step through one call.
2. Rename `splitBill` to `divideBill` using `F2` and confirm every usage updates automatically.
3. Deliberately introduce a type error (e.g. pass a string where `splitBill` expects an `int`)
   and watch the inline diagnostic appear **before** you even try to run the program.

## ⚠️ Common Mistakes

- Skipping the "install supporting tools" prompt the first time you open a Go file — without
  `gopls`, most of this workflow silently doesn't work.
- Editing `.go` files in a plain text editor out of habit, missing out on the live diagnostics
  and one-keystroke refactors this workflow provides for free.
