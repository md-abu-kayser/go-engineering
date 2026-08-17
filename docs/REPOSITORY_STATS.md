# Repository Statistics

> This file records the current shape of the repository and explains what the numbers mean.

## Current snapshot

| Metric                    | Current value |
| ------------------------- | ------------: |
| Go source files           |     **1,235** |
| Curriculum levels         |        **20** |
| Lesson directories        |     **1,160** |
| README files              |     **1,191** |
| Go test files             |        **63** |
| Integrated projects       |        **10** |
| Minimum required Go files |     **1,111** |
| Reserve above minimum     |       **124** |

The headline target is intentionally exceeded. The repository should not sit exactly on the minimum because normal maintenance can temporarily remove or merge a lesson.

---

## What counts as a Go file?

The official count is based on:

```bash
find . -type f -name '*.go'
```

The count includes curriculum source files, tests, projects, and repository tooling. It excludes generated build artifacts and files outside the repository tree.

Run:

```bash
./scripts/count-go.sh
```

to reproduce the count locally.

---

## Curriculum distribution

The repository is organized around 20 levels rather than one flat examples directory. This gives the project a navigable learning spine while allowing each level to grow independently.

### Level sizing philosophy

The number of lessons in a level is allowed to vary. A topic such as concurrency naturally deserves more depth than a short toolchain orientation.

A healthy level should contain:

- foundational concepts;
- edge cases;
- practice examples;
- engineering guidance;
- tests where useful;
- links toward the next conceptual step.

---

## Files are not the unit of mastery

The 1,235-file count is an architectural signal, not a score.

A learner who understands 100 carefully designed examples is ahead of someone who has skimmed 2,000 files without understanding them.

The repository therefore measures its quality through several dimensions:

```text
Breadth
  +
Depth
  +
Tests
  +
Documentation
  +
Real projects
  +
Engineering trade-offs
  =
Useful knowledge base
```

---

## Production-project inventory

The repository includes 10 integrated project tracks. These serve as bridges between isolated lessons and system-level thinking.

They are intentionally positioned outside the level directories because they combine concepts from multiple levels.

---

## Reproducibility

The statistics file should be updated when repository structure changes materially.

Recommended verification sequence:

```bash
./scripts/count-go.sh
./scripts/verify.sh
```

CI also checks the minimum file-count invariant.

---

## Growth policy

The project can scale far beyond 1,111 Go files. New files should be added when they provide a distinct learning outcome or engineering insight.

Good growth:

- another concurrency failure mode;
- a benchmark showing a real trade-off;
- a secure and insecure implementation comparison;
- a new standard-library edge case;
- a distributed-systems failure simulation;
- a new project slice with independent architectural value.

Bad growth:

- renaming the same example repeatedly;
- splitting one trivial statement into many files;
- copying examples without a new concept;
- creating files only to increase the count.

---

## Update policy

When adding a substantial body of work, update all relevant evidence together:

1. lesson README;
2. `docs/LESSON_INDEX.json`;
3. this statistics document;
4. roadmap references when the learning sequence changes;
5. tests and CI expectations when quality gates change.

The repository should always describe what actually exists.
