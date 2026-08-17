# Repository Statistics

> This file records the current shape of the repository, explains what the numbers mean, defines how they are measured, and establishes governance for keeping the reported state trustworthy.

---

# 1. Current Snapshot

| Metric                    | Current Value |
| ------------------------- | ------------: |
| Go source files           |     **1,235** |
| Curriculum levels         |        **20** |
| Lesson directories        |     **1,160** |
| README files              |     **1,191** |
| Go test files             |        **63** |
| Integrated projects       |        **10** |
| Minimum required Go files |     **1,111** |
| Reserve above minimum     |       **124** |
| Curriculum range          | **0% → 100%** |
| Primary language          |        **Go** |

The headline target is intentionally exceeded.

The repository should not sit exactly on the minimum because normal maintenance may temporarily remove, merge, rename, or replace lessons.

The statistics document therefore describes the repository's current state, while the verification tooling defines what counts as acceptable structure.

---

# 2. What These Numbers Mean

Statistics are useful only when they describe something meaningful.

The repository does **not** define quality as:

```text
more files = better repository
```

Instead:

```text
Breadth
  +
Depth
  +
Correctness
  +
Tests
  +
Documentation
  +
Projects
  +
Engineering trade-offs
  +
Maintainability
  =
Useful engineering knowledge base
```

The 1,235-file figure is therefore an architectural signal, not a learning score.

A smaller lesson can be more valuable than many duplicated examples.

---

# 3. Official Go File Count

The official Go source-file count is based on:

```bash
find . -type f -name '*.go'
```

Run:

```bash
./scripts/count-go.sh
```

to reproduce the count locally.

The count includes:

- curriculum source files;
- curriculum tests;
- integrated projects;
- repository tooling written in Go.

It excludes:

- generated build artifacts;
- binaries;
- files outside the repository tree.

The repository should avoid silently changing the counting definition.

If the counting method changes, this document and the verification script must be updated together.

---

# 4. Count Integrity

The same definition should be used by:

```text
local development
        ↓
scripts/count-go.sh
        ↓
scripts/verify.sh
        ↓
CI
        ↓
documentation
```

The goal is deterministic reporting.

Running the same counting command against the same repository state should produce the same result.

---

# 5. Curriculum Distribution

The repository is organized around 20 levels:

```text
00  Getting Started
01  Fundamentals
02  Core Go
03  Intermediate Go
04  Concurrency
05  Standard Library
06  CLI & Systems
07  Web Development
08  Databases & Data Access
09  Testing & Quality
10  Go Architecture
11  Design Patterns
12  Production Engineering
13  Security
14  Performance & Optimization
15  Networking & Distributed Systems
16  Microservices & Cloud Native
17  Advanced Go Engineering
18  Real-World Projects
19  Expert / Professional Go
```

These levels provide the repository's learning spine.

The levels are allowed to differ significantly in file count because difficulty and scope are not uniform.

---

# 6. Level Sizing Philosophy

A large level is not automatically better.

A healthy level should contain the material needed to understand its engineering domain, including where appropriate:

- foundational concepts;
- mental models;
- practical examples;
- edge cases;
- common mistakes;
- debugging guidance;
- exercises;
- tests;
- engineering trade-offs;
- production connections;
- links to subsequent concepts.

A topic such as concurrency may legitimately require substantially more material than initial toolchain setup.

The repository should therefore optimize for **conceptual completeness**, not equal file counts per level.

---

# 7. Lesson Directory Statistics

The repository currently targets:

```text
1,160 lesson directories
```

Lesson count should not be treated as an achievement by itself.

A lesson directory is useful when it contains a meaningful learning unit.

Good lesson boundaries usually correspond to:

```text
one coherent concept
        OR
one closely related cluster of concepts
        OR
one engineering experiment
        OR
one meaningful production scenario
```

Avoid creating lessons merely because a directory name looks impressive.

---

# 8. README Statistics

The current target is:

```text
1,191 README files
```

README files serve different scopes:

```text
Root README
    ↓
Level README
    ↓
Lesson README
    ↓
Project README
```

Documentation should be distributed according to responsibility.

A README should explain the scope of the directory it belongs to rather than repeating repository-wide information.

---

# 9. Test Statistics

The current target is:

```text
63 Go test files
```

This number is not a universal target for every future repository revision.

Tests should exist where they add meaningful confidence.

The repository should prioritize:

```text
important behavior
+
edge cases
+
regression protection
+
deterministic execution
```

over raw test-file count.

A lesson may legitimately have no test when testing would add no educational or engineering value.

A production project should have materially stronger testing expectations.

---

# 10. Integrated Project Inventory

The repository includes:

```text
10 integrated projects
```

These projects are intentionally separate from the lesson tree because they combine concepts from multiple levels.

A simplified model is:

```text
levels/
    teaches individual capabilities

projects/
    integrates capabilities into systems
```

Projects should therefore demonstrate synthesis rather than repeat isolated exercises.

---

# 11. Project Maturity

A substantial project should progressively demonstrate:

```text
requirements
   ↓
architecture
   ↓
implementation
   ↓
testing
   ↓
failure handling
   ↓
security
   ↓
observability
   ↓
performance
   ↓
deployment
   ↓
operational thinking
```

Not every project needs every stage.

The expected maturity should be proportional to project scope.

---

# 12. Minimum File Invariant

The repository currently defines:

```text
Minimum required Go files: 1,111
Current target:            1,235
Reserve:                     124
```

The minimum is a structural invariant, not a quality guarantee.

Its purpose is to detect accidental large-scale removal of curriculum material.

It must not become an incentive to generate low-value files.

---

# 13. Reserve Policy

The current reserve above minimum is:

```text
1,235 - 1,111 = 124
```

The reserve provides space for normal repository maintenance.

The reserve should not be treated as a requirement to add 124 new files.

A repository can improve while the file count stays flat or decreases if lessons are merged into clearer, better-designed units.

---

# 14. Growth Policy

The project may grow well beyond 1,111 Go files.

New files are justified when they provide a distinct:

- concept;
- implementation technique;
- edge case;
- failure mode;
- testing strategy;
- engineering pattern;
- performance investigation;
- security scenario;
- production scenario;
- project capability.

### Good growth

```text
new concurrency failure mode
new benchmark with measurable trade-off
secure vs insecure implementation
standard-library edge case
distributed-system failure simulation
new independent project capability
new testing technique
```

### Bad growth

```text
same example with a different filename
one trivial statement split into many files
copied code without a new learning outcome
files created only to increase statistics
duplicate lessons with no meaningful distinction
```

The governing question is:

> **Would an experienced engineer defend the existence of this file?**

---

# 15. Files Are Not the Unit of Mastery

The repository deliberately rejects file-count learning.

A learner who deeply understands 100 carefully designed implementations is ahead of someone who has skimmed 2,000 examples without understanding them.

Therefore:

```text
file count
    ≠
mastery
```

Instead:

```text
understanding
+
practice
+
debugging
+
testing
+
integration
+
engineering judgment
```

represent more meaningful evidence of progress.

---

# 16. Documentation Health

Useful documentation should make important knowledge discoverable.

Track at least:

```text
README coverage
lesson index coverage
project documentation
architecture documentation
roadmap coverage
broken-link status
```

A high file count with poor navigation is not considered healthy repository growth.

---

# 17. Lesson Quality Health

A lesson should ideally provide appropriate evidence for its scope.

Possible quality dimensions:

| Dimension   | Question                                |
| ----------- | --------------------------------------- |
| Explanation | Is the concept understandable?          |
| Example     | Can the learner see it working?         |
| Experiment  | Can the learner modify behavior?        |
| Exercise    | Can the learner practice independently? |
| Testing     | Is important behavior verified?         |
| Debugging   | Are failure modes teachable?            |
| Production  | Is real-world relevance explained?      |
| Trade-offs  | Are design choices justified?           |
| Navigation  | Can the learner find the next step?     |

A lesson does not need every dimension at maximum depth.

---

# 18. Curriculum Health

A healthy curriculum should have:

```text
prerequisites
      ↓
concept
      ↓
practice
      ↓
verification
      ↓
integration
      ↓
next concept
```

The repository should periodically check for:

- orphan lessons;
- broken links;
- missing prerequisites;
- inconsistent level numbering;
- duplicate lesson concepts;
- stale references;
- unindexed lessons.

---

# 19. Lesson Index Statistics

`docs/LESSON_INDEX.json` acts as machine-readable metadata for the curriculum.

Useful fields include:

```text
level
level_number
lesson_number
title
path
category
difficulty
status
estimated_minutes
prerequisites
skills
production_relevance
testing_required
exercise_available
challenge_available
project_connections
```

This enables:

```text
search
automation
analytics
curriculum dashboards
AI-assisted navigation
future web interfaces
```

The index should be generated deterministically.

---

# 20. Repository Verification

Use:

```bash
./scripts/count-go.sh
./scripts/verify.sh
```

The verification workflow should check, where applicable:

```text
Go file count
required directories
required files
README presence
formatting
module validity
lesson index generation
tests
documentation structure
link integrity
repository metadata
```

The repository should fail loudly when an important structural assumption is violated.

---

# 21. CI Integrity

Continuous integration should verify that normal contributions preserve repository health.

Conceptually:

```text
push / pull request
       ↓
repository integrity
       ↓
formatting
       ↓
module validation
       ↓
static analysis
       ↓
tests
       ↓
race detection
       ↓
linting
       ↓
documentation checks
       ↓
PASS
```

CI should be strict about real correctness problems without becoming unnecessarily fragile for legitimate educational contributions.

---

# 22. Reproducibility

Repository statistics should be reproducible.

Recommended sequence:

```bash
./scripts/count-go.sh
./scripts/verify.sh
```

When practical, generated statistics should come from scripts rather than manual counting.

A number copied manually into documentation can become stale.

Automation is preferred.

---

# 23. Update Policy

Update this file when repository structure changes materially.

When adding a substantial body of work, review:

1. lesson README;
2. `docs/LESSON_INDEX.json`;
3. `docs/ROADMAP.md`;
4. this statistics document;
5. verification scripts;
6. CI expectations;
7. project inventory when relevant.

The repository should always describe what actually exists.

---

# 24. Versioned Statistics

A statistics snapshot describes a point in repository history.

For significant releases, it can be useful to record:

```text
date
commit
Go version
Go file count
lesson count
test count
project count
verification result
```

Example:

```text
Repository snapshot
Date: YYYY-MM-DD
Commit: <short-sha>
Go files: 1,235
Lessons: 1,160
Tests: 63
Projects: 10
Verification: PASS
```

This prevents historical growth claims from becoming ambiguous.

---

# 25. Statistics Must Have Definitions

Every important metric should have a clear definition.

For example:

```text
Go source files
= repository-tree .go files included by the official counting command

Lesson directory
= directory recognized by repository curriculum structure

Project
= independently documented integrated project track

README
= Markdown README file recognized by repository documentation conventions
```

The exact definition should remain stable unless intentionally changed.

---

# 26. Avoiding Metric Gaming

The repository must never optimize for statistics alone.

Do not:

```text
split one concept into artificial files
duplicate documentation
duplicate tests
create shallow exercises
inflate project counts
```

The quality rule is:

> **A statistic is useful only when it corresponds to real engineering or educational value.**

---

# 27. Recommended Health Metrics

Beyond raw file counts, future versions may track:

```text
lesson coverage
test coverage
broken-link count
indexed-lesson percentage
documented-lesson percentage
exercise coverage
challenge coverage
project integration count
stale lesson count
deprecated lesson count
CI pass rate
average lesson complexity
```

These are more informative than file count alone.

---

# 28. Repository Quality Dashboard

A future generated dashboard could present:

```text
Repository
├── 20 curriculum levels
├── 1,160 lessons
├── 1,235 Go files
├── 63 test files
├── 10 integrated projects
├── documentation health
├── CI health
├── lesson-index health
├── broken-link health
└── curriculum coverage
```

The purpose is visibility, not vanity.

---

# 29. Long-Term Growth Model

The repository is expected to evolve.

Healthy evolution looks like:

```text
more knowledge
      ↓
better organization
      ↓
better examples
      ↓
better testing
      ↓
better explanations
      ↓
better projects
      ↓
better engineering judgment
```

Not:

```text
more files
      ↓
more directories
      ↓
more numbers
```

---

# 30. Repository Maintenance Rules

Maintainability requires periodic review.

Review for:

- stale Go versions;
- outdated APIs;
- obsolete libraries;
- duplicated lessons;
- dead examples;
- broken links;
- obsolete project instructions;
- inaccurate statistics;
- inconsistent naming;
- outdated screenshots or diagrams;
- deprecated techniques.

A professional knowledge repository must evolve with the language and ecosystem.

---

# 31. Deprecation Policy

Lessons may eventually become deprecated.

When a lesson is no longer preferred:

```text
active
  ↓
deprecated
  ↓
replacement identified
  ↓
migration guidance
  ↓
eventual removal
```

Do not silently delete historically important lessons.

Explain why the recommendation changed.

---

# 32. Consolidation Policy

Two lessons may be merged when:

```text
their learning outcomes are effectively identical
```

A reduction in file count after a well-designed consolidation is considered a positive result when it improves clarity.

The repository values:

```text
quality
over
quantity
```

---

# 33. Project Independence

Integrated projects should remain independently understandable.

A project should avoid hidden dependencies on:

- unfinished lessons;
- another project's private code;
- developer-specific paths;
- local-only credentials;
- untracked generated files;
- undocumented external services.

A project may conceptually depend on lessons, but it should not require the learner to have a particular private workspace state.

---

# 34. Statistics and Roadmap Relationship

The roadmap explains:

```text
what should be learned
```

The statistics file explains:

```text
what currently exists
```

These are related but not identical.

The roadmap may describe a target state before all material exists.

The statistics file must describe actual repository state.

Therefore:

```text
ROADMAP.md
    = intended curriculum

REPOSITORY_STATS.md
    = observed repository state
```

---

# 35. Statistics and Architecture Relationship

`docs/ARCHITECTURE.md` explains:

```text
how the repository is designed
```

`docs/ROADMAP.md` explains:

```text
how a learner progresses
```

`docs/REPOSITORY_STATS.md` explains:

```text
what currently exists
```

The three documents should remain aligned:

```text
ARCHITECTURE
      ↓
defines structure
      ↓
ROADMAP
      ↓
defines learning progression
      ↓
REPOSITORY_STATS
      ↓
reports actual implementation
```

---

# 36. Definition of a Healthy Repository

A healthy repository should demonstrate:

```text
structured curriculum
+
clear navigation
+
independent lessons
+
reproducible tooling
+
meaningful tests
+
integrated projects
+
quality documentation
+
deterministic metadata
+
CI enforcement
+
maintainable growth
```

The project is healthy when a new contributor can understand:

```text
what exists
why it exists
where it belongs
how to verify it
how to extend it
```

---

# 37. Final Principle

The repository should never confuse scale with quality.

The numbers are useful because they show that the project is substantial.

The real achievement is that those numbers represent:

```text
purposeful lessons
+
coherent progression
+
working Go code
+
tests
+
experiments
+
real systems
+
engineering decisions
+
documented trade-offs
```

The guiding rule remains:

> **Every statistic should correspond to something a learner or engineer can actually use.**
