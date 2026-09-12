# 47 — Benchmark Baseline

## 🎯 Learning Objectives

By the end of this lesson you will be able to:

- Explain the engineering purpose of **Benchmark Baseline** in the context of testing quality.
- Read and modify a small, deterministic Go example without relying on external services.
- Identify the boundary or language feature demonstrated by the example and its trade-off.

## 📖 Concept

**Benchmark Baseline** is most useful when it makes a program's behavior clearer, safer, or easier to change.
This lesson deliberately keeps the example local and observable: it demonstrates a pure function with an error path that is ready for a table-driven test or benchmark. The
same idea can grow into a production concern, but a small executable example makes the contract
visible before configuration, networking, or infrastructure obscure it.

## 🔍 Code Walkthrough (main.go)

The program starts with a lesson-specific constant so the output identifies the concept being
run. Its supporting types and functions hold the behavior outside main(). That separation is
intentional: compute or validate in a focused function, then let main() assemble dependencies
and display the result. If a real application later needs a test, that focused behavior is the
natural seam to test first.

Read the example in this order:

1. Find the small type or interface that states the program's contract.
2. Follow the helper function that applies the contract to a concrete value.
3. Return to main() to see the deterministic input and output.

## ▶️ How to Run

`ash
cd level-09-testing-quality/47-benchmark-baseline
go run .
`

## ✅ Expected Output

The first line names this lesson. The following line prints the result of the example. Exact
formatting is intentionally simple so you can change an input and immediately predict what will
change in the output.

## 🧠 Key Takeaways

- Small examples are easier to verify, change, and reuse than a monolithic demo.
- Keep I/O and presentation at the edge; place the lesson's decision in a named function or type.
- Prefer explicit contracts and deterministic inputs while learning a new engineering concept.

## 🛠️ Try It Yourself

1. Change one input in main() and predict the new output before running the program.
2. Extract one assertion you would write if this example had a _test.go file.
3. Add a second valid case, then add one invalid case and decide where the error should be handled.

## ⚠️ Common Mistakes

- Treating the demonstration as a complete production implementation instead of a focused mental model.
- Hiding validation, I/O, or errors inside main() where they become difficult to test.
- Adding a dependency before the standard-library example and its trade-offs are understood.