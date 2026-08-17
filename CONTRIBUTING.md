# Contributing to Go Engineering

Thank you for helping build this repository. The project is intentionally larger and more structured than a normal tutorial, so contributions should protect its learning quality as it grows.

## What belongs here?

A strong contribution usually does at least one of these:

- teaches a distinct Go concept;
- demonstrates an important edge case;
- documents a production trade-off;
- adds a useful test or benchmark;
- improves an existing example's correctness;
- adds a meaningful integration project;
- improves repository tooling or documentation.

Adding files only to increase the repository count is not considered a useful contribution.

---

## Lesson contribution standard

A new lesson should normally look like:

```text
NN-topic-name/
├── README.md
├── main.go
└── optional_test.go
```

### README expectations

Explain:

1. what the learner should understand;
2. why the topic matters;
3. prerequisites;
4. how to run it;
5. what to observe;
6. common mistakes;
7. a practical exercise.

### Go code expectations

Prefer:

- idiomatic naming;
- small functions;
- explicit error handling;
- deterministic output;
- simple dependencies;
- useful comments about decisions;
- code that can be modified safely by a learner.

Avoid cleverness that obscures the lesson.

---

## Production-project standard

For `projects/`, contributions should demonstrate multiple concepts working together.

A project is stronger when it has clear boundaries such as:

```text
cmd/
internal/domain/
internal/service/
internal/repository/
internal/transport/
```

The exact structure may change according to the problem. Architecture is a means, not a ritual.

---

## Development workflow

Before opening a pull request:

```bash
gofmt -w .
go mod verify
go vet ./...
go test ./...
./scripts/verify.sh
```

When available, also run:

```bash
golangci-lint run
```

Do not commit generated binaries, editor metadata, secrets, local environment files, or dependency caches.

---

## Commit conventions

Use Conventional Commits.

Examples:

```text
feat: add cancellation-aware worker pool lesson
fix: correct slice aliasing example
refactor: simplify repository service boundary
docs: expand concurrency roadmap
perf: benchmark buffered worker pipeline
test: add table-driven URL validation cases
chore: update CI toolchain configuration
```

Keep the first line concise and focused on the actual change.

---

## Pull requests

A useful pull request should explain:

- what changed;
- why it changed;
- which level/project it belongs to;
- how it was verified;
- any trade-offs a reviewer should know about.

For teaching material, include the learner outcome, not only the implementation details.

---

## Updating generated metadata

When lesson structure changes, regenerate or update:

- `docs/LESSON_INDEX.json`
- `docs/REPOSITORY_STATS.md`
- relevant level README content
- roadmap references when sequencing changes

CI intentionally treats the lesson index as reproducible metadata.

---

## Code review philosophy

Reviewers should prioritize:

1. correctness;
2. learning value;
3. clarity;
4. idiomatic Go;
5. testability;
6. maintainability;
7. security and failure behavior.

A technically correct example can still be rejected if it teaches a poor production habit.

---

## Educational integrity

Please do not copy large portions of proprietary code or introduce examples whose licensing is unclear. Prefer code that the project can safely redistribute and learners can confidently study.
