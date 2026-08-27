# 10 — go test

## 🎯 Learning Objectives

- Write a table-driven test — the idiomatic Go way to cover many cases with little code.
- Use `t.Run` to create readable, individually-reportable subtests.
- Use the most common `go test` flags: `-v`, `-run`, `-cover`, `-bench`.
- Write and run a basic benchmark.

## 📖 Concept

Go's testing philosophy is deliberately minimal: no assertion libraries, no test framework
magic — just the standard `testing` package and a naming convention:

- A test file is named `<name>_test.go` and lives **next to** the code it tests.
- A test function is named `TestXxx(t *testing.T)` — it must start with `Test`, take exactly one
  parameter of type `*testing.T`, and return nothing.
- You report failure by calling `t.Error` / `t.Errorf` (test continues) or `t.Fatal` /
  `t.Fatalf` (test stops immediately).

### Table-driven tests

Instead of writing a separate `TestXxx` function per case, Go convention is to define a slice of
"cases" and loop over it:

```go
cases := []struct {
    name string
    in   string
    want bool
}{
    {"simple palindrome", "level", true},
    {"not a palindrome", "gopher", false},
}

for _, tc := range cases {
    t.Run(tc.name, func(t *testing.T) {
        got := isPalindrome(tc.in)
        if got != tc.want {
            t.Errorf("isPalindrome(%q) = %t, want %t", tc.in, got, tc.want)
        }
    })
}
```

`t.Run(name, func(t *testing.T) {...})` creates a **subtest** — each case gets its own named
result in `-v` output (`TestIsPalindrome/simple_palindrome`), and a failure in one case doesn't
stop the others from running.

### Useful `go test` flags

| Flag | Purpose |
|---|---|
| `-v` | Verbose — show every test (and subtest) name and its result. |
| `-run <regex>` | Only run tests whose name matches the pattern, e.g. `-run TestIsPalindrome`. |
| `-cover` | Report the percentage of code statements executed by the tests. |
| `-bench <regex>` | Run benchmarks matching the pattern (use `-run=^$` to skip regular tests). |
| `-race` | Run tests with the race detector enabled. |

## ▶️ How to Run

```bash
cd level-00-getting-started/10-go-test-command
go run main.go
go test -v ./...
go test -cover ./...
go test -bench=. -run=^$ ./...
```

## ✅ Expected Output (test run)

```
=== RUN   TestIsPalindrome
=== RUN   TestIsPalindrome/empty_string
=== RUN   TestIsPalindrome/single_char
=== RUN   TestIsPalindrome/simple_palindrome
=== RUN   TestIsPalindrome/racecar
=== RUN   TestIsPalindrome/not_a_palindrome
=== RUN   TestIsPalindrome/two_different_chars
--- PASS: TestIsPalindrome (0.00s)
    --- PASS: TestIsPalindrome/empty_string (0.00s)
    --- PASS: TestIsPalindrome/single_char (0.00s)
    --- PASS: TestIsPalindrome/simple_palindrome (0.00s)
    --- PASS: TestIsPalindrome/racecar (0.00s)
    --- PASS: TestIsPalindrome/not_a_palindrome (0.00s)
    --- PASS: TestIsPalindrome/two_different_chars (0.00s)
PASS
```

## 🧠 Key Takeaways

- Go's test framework is just a naming convention plus the `testing` package — nothing exotic.
- Table-driven tests + `t.Run` are the standard way to cover many cases cleanly.
- `-cover` tells you how much of your code your tests actually exercise.
- Benchmarks (`func BenchmarkXxx(b *testing.B)`) use the same file and toolchain as tests.

## 🛠️ Try It Yourself

1. Add a new case to the table for a palindrome with mixed case, e.g. `"Level"` — notice it
   fails, since `isPalindrome` is case-sensitive. Fix `isPalindrome` to ignore case (hint:
   `strings.ToLower`) and confirm the test passes.
2. Run `go test -cover ./...` and note the coverage percentage.
3. Deliberately break `isPalindrome` (e.g. change `!=` to `==`) and watch the test output pinpoint
   exactly which subtests fail.

## ⚠️ Common Mistakes

- Naming a test function `Testxxx` (lowercase `x`) instead of `TestXxx` — Go won't recognize it
  as a test at all, and it will silently not run.
- Forgetting `-run=^$` when only benchmarking — without it, `go test -bench=.` also runs your
  regular tests first.
