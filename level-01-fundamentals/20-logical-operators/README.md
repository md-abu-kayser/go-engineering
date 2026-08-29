# 20 — Logical Operators

## 🎯 Learning Objectives

- Combine multiple comparison expressions into a single, correct compound condition.
- Correctly account for `&&`/`||` precedence in a multi-condition expression.
- Apply De Morgan's laws to simplify or verify a negated compound condition.

## 📖 Concept

[Lesson 07](../07-booleans) introduced `bool` and short-circuit evaluation. This lesson focuses
on the next step: **combining multiple comparisons** into a single, correct real-world condition.

### A realistic compound condition

```go
func canVote(age int, isCitizen bool) bool {
    return age >= 18 && isCitizen
}
```

This reads naturally as "age is at least 18, **and** is a citizen" — exactly matching the logical
requirement in plain English, which is the goal for any compound boolean expression: it should
read as closely as possible to how you'd state the rule out loud.

### Three-condition expressions need precedence awareness

```go
canDrive := age >= 18 || (hasLicense && hasPermit)
```

As [lesson 17](../17-operator-precedence) covered, `&&` binds tighter than `||` — so the
parentheses here are technically optional, but included anyway for a human skimming the
condition; the intended grouping ("either you're 18+, or you have both a license and a permit")
is instantly clear with them, and only clear-after-a-moment's-thought without them.

### De Morgan's laws

When you need to **negate** a compound condition, De Morgan's laws tell you exactly how the `!`
distributes:

```go
!(a && b)  ==  !a || !b
!(a || b)  ==  !a && !b
```

In words: "NOT (both)" is the same as "NOT the first, OR NOT the second" — and "NOT (either)" is
the same as "NOT the first, AND NOT the second." This is genuinely useful for simplifying a
confusing negated condition in real code — e.g. rewriting
`if !(user.IsActive() && user.HasPermission())` into the (arguably clearer)
`if !user.IsActive() || !user.HasPermission()`.

## 🔍 Code Walkthrough (`main.go`)

```go
notBoth := !(a && b)
deMorgan1 := !a || !b
fmt.Printf("!(a && b)      = %t\n", notBoth)
fmt.Printf("!a || !b       = %t (equivalent, by De Morgan's law)\n", deMorgan1)
```

Both expressions are computed **independently** and printed side by side — proof, not just
assertion, that they always produce the same result for the same inputs.

## ▶️ How to Run

```bash
cd level-01-fundamentals/20-logical-operators
go run main.go
```

## ✅ Expected Output

```
=== Logical Operators ===
----------------------------------
canVote(20, true)  : true
canVote(20, false) : false
canVote(15, true)  : false

canDrive (16, no license, has permit) : false

--- De Morgan's laws ---
!(a && b)      = true
!a || !b       = true (equivalent, by De Morgan's law)
!(a || b)      = false
!a && !b       = false (equivalent, by De Morgan's law)
```

## 🧠 Key Takeaways

- Write compound boolean expressions so they read as closely as possible to the plain-English rule.
- `&&` binds tighter than `||` in a mixed expression — use parentheses for human clarity even when
  the compiler doesn't strictly need them.
- De Morgan's laws: `!(a && b) == (!a || !b)`, and `!(a || b) == (!a && !b)`.
- These laws are genuinely useful for simplifying a confusing negated real-world condition.

## 🛠️ Try It Yourself

1. Change `a`/`b` to different combinations of `true`/`false` and confirm both De Morgan
   equivalences hold in every case (there are only four combinations to check).
2. Rewrite `canDrive`'s condition using De Morgan's law to express its **negation**
   ("cannot drive") instead, and confirm it's the logical opposite.
3. Write a three-condition De Morgan expansion (`!(a && b && c)`) and verify it against
   `!a || !b || !c`.

## ⚠️ Common Mistakes

- Negating a compound condition by just slapping `!` in front without redistributing it — 
  `!(a && b)` and `!a && !b` are **not** the same thing (the second is `!a && !b`, but De Morgan's
  law says the correct equivalent is `!a || !b`, with `||` not `&&`).
- Skipping parentheses in a three-or-more-condition expression mixing `&&` and `||`, producing
  code that's *technically* correct but genuinely hard for a reviewer to verify at a glance.
