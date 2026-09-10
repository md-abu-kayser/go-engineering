# Level 01 — Fundamentals

With the environment, tooling, and project-structure groundwork from
[Level 00](../level-00-getting-started) in place, this level moves into Go's actual **language
fundamentals**: how values are declared, the built-in types you'll use in virtually every
program, the operators that act on them, every control-flow construct Go offers, functions as a
first-class concept, and how scope and visibility work across a whole package.

By the end of this level you will be able to:

- Declare variables every way Go supports (`var`, `:=`, grouped, package-level), and know exactly
  when each is appropriate.
- State the zero value for any type from memory, and know which zero values are directly usable.
- Declare and compute constants, including flexible untyped constants and `iota`-based
  enumerations — including the classic bit-flag pattern.
- Use `bool` and Go's short-circuit logical operators correctly, including De Morgan's laws.
- Work with text correctly at every level: `string` (bytes), `rune` (Unicode code points), and
  `byte` (raw data) — and know exactly when each is the right tool.
- Choose the right integer type, understand truncating division, and recognize both signed
  overflow and unsigned underflow before they become bugs.
- Use `float64`/`float32` correctly, including comparing floats safely and a Go-specific subtlety
  in how untyped constant arithmetic can mask (or reveal) classic floating-point imprecision.
- Use Go's built-in complex number types.
- Convert explicitly between numeric types and between numbers and strings — including the
  classic `string(anInt)` gotcha.
- Apply Go's full operator set correctly: precedence, arithmetic, comparison, logical, bitwise
  (including Go's own `&^` bit-clear operator), assignment, and increment/decrement — and know
  the Go-specific rule that `++`/`--` are statements, never expressions.
- Control program flow with `if` (including init statements and mandatory braces), `switch`
  (value-based, tagless, and a preview of type switches), and every shape of Go's single
  looping keyword, `for` — counted, while-style, and infinite.
- Iterate any collection with `range`, control loops precisely with `break`/`continue` (including
  their nearest-enclosing-construct scoping rules) and their labeled forms for nested loops.
- Use `defer` correctly — including its argument-evaluation timing and LIFO execution order — and
  use `panic`/`recover` as language-level control-flow mechanisms, including recover's exact
  placement rule.
- Declare and call functions in every shape Go supports: multiple/named return values, variadic
  parameters, function values, anonymous functions, closures, and recursion.
- Reason precisely about scope (block, function, and package) and shadowing, and use exported vs.
  unexported identifiers as Go's entire cross-package visibility mechanism — including where
  `main()` and package initialization fit into the picture.

## Lessons

**Part 1 — Declarations & values**

| #   | Folder                  | Topic              |
| --- | ----------------------- | ------------------ |
| 01  | `01-var-declarations`   | var Declarations   |
| 02  | `02-short-declarations` | Short Declarations |
| 03  | `03-zero-values`        | Zero Values        |
| 04  | `04-constants`          | Constants          |
| 05  | `05-typed-constants`    | Typed Constants    |
| 06  | `06-iota`               | iota               |

**Part 2 — Basic types**

| #   | Folder                  | Topic              |
| --- | ----------------------- | ------------------ |
| 07  | `07-booleans`           | Booleans           |
| 08  | `08-strings`            | Strings            |
| 09  | `09-runes`              | Runes              |
| 10  | `10-bytes`              | Bytes              |
| 11  | `11-integers`           | Integers           |
| 12  | `12-unsigned-integers`  | Unsigned Integers  |
| 13  | `13-floating-point`     | Floating-Point     |
| 14  | `14-complex-numbers`    | Complex Numbers    |
| 15  | `15-numeric-conversion` | Numeric Conversion |
| 16  | `16-string-conversion`  | String Conversion  |

**Part 3 — Operators**

| #   | Folder                    | Topic                 |
| --- | ------------------------- | --------------------- |
| 17  | `17-operator-precedence`  | Operator Precedence   |
| 18  | `18-arithmetic-operators` | Arithmetic Operators  |
| 19  | `19-comparison-operators` | Comparison Operators  |
| 20  | `20-logical-operators`    | Logical Operators     |
| 21  | `21-bitwise-operators`    | Bitwise Operators     |
| 22  | `22-assignment-operators` | Assignment Operators  |
| 23  | `23-increment-decrement`  | Increment & Decrement |

**Part 4 — Control flow: conditionals**

| #   | Folder                     | Topic                   |
| --- | -------------------------- | ----------------------- |
| 24  | `24-if-statements`         | if Statements           |
| 25  | `25-if-with-init`          | if with Init            |
| 26  | `26-else-branches`         | else Branches           |
| 27  | `27-switch-statements`     | switch Statements       |
| 28  | `28-expression-switches`   | Expression Switches     |
| 29  | `29-type-switches-preview` | Type Switches (Preview) |

**Part 5 — Control flow: loops**

| #   | Folder                | Topic            |
| --- | --------------------- | ---------------- |
| 30  | `30-for-loop`         | for Loop         |
| 31  | `31-while-style-for`  | While-Style for  |
| 32  | `32-infinite-loop`    | Infinite Loop    |
| 33  | `33-range-loop`       | range Loop       |
| 34  | `34-break`            | break            |
| 35  | `35-continue`         | continue         |
| 36  | `36-labeled-break`    | Labeled break    |
| 37  | `37-labeled-continue` | Labeled continue |

**Part 6 — defer, panic & recover**

| #   | Folder                | Topic            |
| --- | --------------------- | ---------------- |
| 38  | `38-defer-basics`     | defer Basics     |
| 39  | `39-defer-lifo-order` | defer LIFO Order |
| 40  | `40-panic-basics`     | panic Basics     |
| 41  | `41-recover-basics`   | recover Basics   |

**Part 7 — Functions**

| #   | Folder                      | Topic                  |
| --- | --------------------------- | ---------------------- |
| 42  | `42-functions`              | Functions              |
| 43  | `43-parameters`             | Parameters             |
| 44  | `44-multiple-return-values` | Multiple Return Values |
| 45  | `45-named-results`          | Named Results          |
| 46  | `46-variadic-functions`     | Variadic Functions     |
| 47  | `47-function-values`        | Function Values        |
| 48  | `48-anonymous-functions`    | Anonymous Functions    |
| 49  | `49-closures`               | Closures               |
| 50  | `50-recursion`              | Recursion              |

**Part 8 — Scope & program structure**

| #   | Folder                      | Topic                  |
| --- | --------------------------- | ---------------------- |
| 51  | `51-scope`                  | Scope                  |
| 52  | `52-shadowing`              | Shadowing              |
| 53  | `53-package-scope`          | Package Scope          |
| 54  | `54-exported-identifiers`   | Exported Identifiers   |
| 55  | `55-unexported-identifiers` | Unexported Identifiers |
| 56  | `56-doc-comments`           | Doc Comments           |
| 57  | `57-main-function`          | The main Function      |
| 58  | `58-package-initialization` | Package Initialization |

## How to work through this level

Go through the lessons **in order** — each one assumes you understand the ones before it. For
every lesson:

1. Read the lesson's `README.md`.
2. Open `main.go` (and any accompanying files) and read the comments before running anything.
3. Run the program (`go run main.go`) and compare the output to what the README predicts.
4. Try the "Try It Yourself" exercise at the bottom of the README before moving on.

Once you've completed all fifty-eight lessons, you're ready for the next module of Level 01
(data structures: arrays, slices, and maps in depth), continuing the same structure.
