# Architecture Guide

## Repository layers

- `level-00` through `level-19`: the curriculum
- `projects/`: integrated systems
- `docs/`: navigation and design guidance
- `scripts/`: repository quality gates
- `.github/`: continuous integration
- `tools/`: tiny repository utilities

## Lesson contract

Each lesson directory is a self-contained unit with a README and a focused Go source file. The unit should be understandable independently, runnable locally, and easy to evolve without changing unrelated lessons.
