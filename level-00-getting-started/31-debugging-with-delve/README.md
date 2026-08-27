# 31 — Debugging with Delve

## 🎯 Learning Objectives

- Install and run [Delve](https://github.com/go-delve/delve) (`dlv`), Go's standard debugger.
- Set a breakpoint, step through code, and inspect variables from the command line.
- Understand that VS Code's built-in debugger ([lesson 29](../29-editor-workflow)) is a UI on
  top of this exact same tool.

## 📖 Concept

[Lesson 29](../29-editor-workflow) showed debugging through VS Code's UI — click the gutter,
press `F5`. Underneath, that entire experience is powered by **Delve**, a standalone command-line
debugger built specifically for Go (general-purpose debuggers like `gdb` don't understand Go's
goroutines, channels, or runtime well; Delve does). This lesson uses Delve **directly**, which is
useful for:

- Debugging on a remote machine or in a container with no editor attached.
- Understanding exactly what's happening when VS Code's debug UI feels like a "black box".
- Using any other editor (Vim, Emacs, JetBrains) that also drives Delve underneath.

### Installing Delve

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

(If you're using the VS Code workflow from lesson 29, this was likely already installed for you.)

### Starting a debug session

```bash
cd level-00-getting-started/31-debugging-with-delve
dlv debug .
```

This compiles the program with debugging information and drops you into an interactive
`(dlv)` prompt — the program is loaded but **not yet running**.

### Core Delve commands

| Command | Effect |
|---|---|
| `break main.applyDiscount` (or `b main.go:14`) | Set a breakpoint at a function or file:line |
| `continue` (`c`) | Run until the next breakpoint |
| `next` (`n`) | Step over the current line |
| `step` (`s`) | Step into a function call |
| `print <expr>` (`p`) | Print a variable or expression's current value |
| `locals` | Print every local variable in the current scope |
| `args` | Print the current function's arguments |
| `continue` again | Resume until the next breakpoint or program exit |
| `quit` (`q`) | Exit Delve |

## 🔍 A worked debugging session

This lesson's `applyDiscount` has a subtle bug: for prices that don't divide evenly, integer
division silently truncates the discount downward, which is arguably fine — but let's *prove*
what's happening at each step, rather than guessing, using Delve itself:

```bash
dlv debug .
```

```
(dlv) break main.applyDiscount
Breakpoint 1 set at ... for main.applyDiscount() ./main.go:14
(dlv) continue
> main.applyDiscount() ./main.go:14 (hits goroutine(1):1 total)
(dlv) args
priceCents = 1999
percentOff = 10
(dlv) next
(dlv) print discount
199
(dlv) next
(dlv) print priceCents - discount
1800
(dlv) continue
```

Notice you didn't have to add a single `fmt.Println` to see `discount`'s exact value at that
exact moment — `dlv` let you pause execution and ask directly.

## ▶️ How to Run

```bash
cd level-00-getting-started/31-debugging-with-delve
go run main.go              # normal run, no debugger
dlv debug .                  # interactive debugging session (see above)
```

## ✅ Expected Output (normal run)

```
price: 1999 -> after 10% off: 1800
price: 500 -> after 10% off: 450
price: 1250 -> after 10% off: 1125
Total: 3375 cents

See the README for a step-by-step dlv session on this exact program.
```

## 🧠 Key Takeaways

- Delve (`dlv`) is Go's purpose-built debugger — VS Code's debug UI drives it for you.
- `break`, `continue`, `next`, `step`, and `print` cover the vast majority of real debugging needs.
- `dlv debug .` compiles with debug info and pauses before running, ready for breakpoints.
- Command-line Delve works anywhere — remote servers, containers, any editor — not just VS Code.

## 🛠️ Try It Yourself

1. Run through the worked session above yourself, line by line, on this lesson's `main.go`.
2. Set a breakpoint inside the `for` loop in `main` instead, and use `continue` repeatedly to
   watch `total` accumulate across each of the three cart items.
3. Try `dlv exec ./31-debugging-with-delve` after a `go build` instead of `dlv debug .`, and
   compare the experience of debugging an already-built binary versus building fresh each time.

## ⚠️ Common Mistakes

- Forgetting `continue` after setting a breakpoint — `dlv debug .` starts paused, before your
  program has even begun; nothing runs until you tell it to.
- Confusing `next` (step over — stays in the current function) with `step` (step into — follows
  the call into whatever function is being called) — mixing them up is the most common source of
  "wait, why did it jump in there?" confusion for first-time Delve users.
