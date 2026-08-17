# Go Engineering 1111+

> **A world-class Go learning ecosystem: from absolute beginner to production engineering and expert systems design.**

This repository is intentionally built as more than a tutorial. It is a structured **Go University + Engineering Handbook + Practical Code Laboratory + Production Reference + Portfolio**.

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white) ![Lessons](https://img.shields.io/badge/Go%20files-1%2C000%2B-111111) ![Levels](https://img.shields.io/badge/levels-20-6f42c1) ![Focus](https://img.shields.io/badge/focus-production%20engineering-0f766e)

## Why this exists

A serious Go learner needs more than syntax examples. They need progressively harder problems, package design, concurrency, networking, APIs, databases, tests, architecture, performance, security, distributed systems, and repeated exposure to production constraints. This repository turns those concerns into a navigable curriculum.

## The 0% → 100% journey

1. **Foundation** — install Go, understand the toolchain, packages, modules, and execution model.
2. **Language fluency** — types, functions, control flow, slices, maps, structs, methods, interfaces, and errors.
3. **Engineering depth** — generics, reflection, package design, modules, and dependency boundaries.
4. **Concurrency** — goroutines, channels, synchronization, cancellation, worker pools, pipelines, and backpressure.
5. **Standard library mastery** — I/O, files, time, encoding, crypto primitives, networking, templates, and process control.
6. **Systems + web** — CLIs, HTTP, APIs, middleware, validation, authentication, uploads, sessions, and graceful shutdown.
7. **Data** — SQL, transactions, PostgreSQL concepts, repositories, caching, idempotency, and migrations.
8. **Quality** — unit, integration, benchmark, fuzz, race, contract, and architecture-oriented testing.
9. **Architecture** — clean architecture, ports/adapters, DDD concepts, dependency inversion, modular monoliths, and CQRS awareness.
10. **Production** — configuration, structured logging, metrics, tracing, health, resilience, rollout, and incident readiness.
11. **Security** — secure authentication, authorization, secret handling, input validation, web defenses, TLS, and auditability.
12. **Performance** — allocation behavior, profiling, GC pressure, contention, hot paths, batching, and performance budgets.
13. **Distributed systems** — TCP/UDP, RPC, queues, service discovery, event delivery, consistency, and fault handling.
14. **Cloud-native** — containers, Kubernetes concepts, CI/CD, release artifacts, configuration, scaling, and operational boundaries.
15. **Internals** — runtime, memory model, scheduler concepts, compiler behavior, profiling, tracing, unsafe rules, and cgo boundaries.
16. **Real systems** — APIs, auth, queues, payments-style ledgers, realtime services, file pipelines, notifications, gateways, and multi-tenancy.
17. **Expert practice** — capacity, failure domains, tail latency, migrations, incident response, disaster recovery, architecture review, and capstone engineering.

## Repository architecture

Every level has 55 focused lesson directories. Every lesson has its own `README.md` and executable `.go` example so that concepts remain isolated and navigable.

| Level | Domain | Lessons | Path |
|---:|---|---:|---|
| 00 | Getting Started | 58 | `level-00-getting-started/` |
| 01 | Go Fundamentals | 58 | `level-01-fundamentals/` |
| 02 | Core Go | 59 | `level-02-core-language/` |
| 03 | Intermediate Go | 57 | `level-03-intermediate-go/` |
| 04 | Concurrency | 58 | `level-04-concurrency/` |
| 05 | Standard Library Mastery | 58 | `level-05-standard-library/` |
| 06 | CLI & Systems Programming | 57 | `level-06-cli-systems/` |
| 07 | Web Development | 58 | `level-07-web-development/` |
| 08 | Databases & Data Access | 59 | `level-08-databases/` |
| 09 | Testing & Quality | 59 | `level-09-testing-quality/` |
| 10 | Go Architecture | 58 | `level-10-architecture/` |
| 11 | Design Patterns | 57 | `level-11-design-patterns/` |
| 12 | Production Engineering | 59 | `level-12-production-engineering/` |
| 13 | Security | 58 | `level-13-security/` |
| 14 | Performance & Optimization | 58 | `level-14-performance/` |
| 15 | Networking & Distributed Systems | 58 | `level-15-networking-distributed/` |
| 16 | Microservices & Cloud-Native Go | 58 | `level-16-cloud-native/` |
| 17 | Advanced Go Engineering | 57 | `level-17-advanced-internals/` |
| 18 | Real-World Projects | 56 | `level-18-real-world-projects/` |
| 19 | Expert / Professional Go | 60 | `level-19-expert-professional/` |

## The 1,111+ `.go` philosophy

The file count is not achieved by splitting meaningless snippets into microscopic pieces. The repository uses **micro-lessons** deliberately: one concept, one mental model, one concrete implementation. Examples are small enough to inspect quickly but substantive enough to modify. Selected lessons also carry test files, and the repository includes project-level Go programs for capstones and tools.

The generated curriculum contains **1,233 Go source files** across the lesson corpus, supporting tests, tools, and project exercises. Use `scripts/count-go.sh` to verify the exact count after checkout.

## How to study

```bash
# 1. Clone
git clone https://github.com/md-abu-kayser/go-engineering-1111.git
cd go-engineering-1111

# 2. Verify your toolchain
go version

# 3. Run any lesson
cd level-01-go-fundamentals/05-runes
go run .

# 4. Return to the root and validate the repo
cd ../..
./scripts/verify.sh
```

### The recommended loop

**Read → Run → Predict → Modify → Break → Fix → Test → Explain → Refactor.**

Do not rush through lesson counts. A file is complete when you can explain its behavior, change it safely, and predict the consequences.

## Progress tracking

Create a personal checklist in your fork or use the lesson index at `docs/LESSON_INDEX.json`. A practical milestone system is:

- [ ] 0–10% — fundamentals + core language
- [ ] 10–25% — intermediate Go + standard library
- [ ] 25–40% — concurrency + CLI/systems
- [ ] 40–55% — web + databases + testing
- [ ] 55–70% — architecture + patterns + production engineering
- [ ] 70–82% — security + performance
- [ ] 82–92% — networking + distributed + cloud-native
- [ ] 92–97% — runtime/internals
- [ ] 97–100% — real-world projects + expert capstone

## Technology coverage

The curriculum emphasizes Go's standard library first, then production concerns around HTTP, SQL/PostgreSQL, caching, distributed messaging, Docker/Kubernetes concepts, observability, profiling, security, and service architecture. Third-party integrations are intentionally treated as **boundaries** rather than prerequisites so the engineering concepts remain portable.

## Engineering standards

- idiomatic Go and `gofmt`
- explicit error handling
- small functions and cohesive packages
- context-aware I/O for request-scoped work
- interfaces at dependency consumers
- deterministic tests
- race detection for concurrent code
- benchmarks for performance claims
- secure defaults and least privilege
- structured logs and operational signals
- clear package boundaries
- documentation that explains intent and trade-offs

## Quality gates

The repository includes CI configuration for formatting, vetting, and testing. Local verification is intentionally simple:

```bash
./scripts/count-go.sh
gofmt -l .
go vet ./...
go test ./...
```

For performance work, add benchmarks and compare before/after behavior rather than relying on intuition. For security work, assume inputs are hostile and secrets are ephemeral.

## Projects

The `projects/` directory converts isolated lessons into coherent systems: production API, authentication service, job platform, file-processing pipeline, realtime system, URL shortener, notification platform, distributed-systems lab, multi-tenant SaaS API, and an expert production capstone.

## Contribution guidelines

New lessons should have a single learning objective, idiomatic implementation, clear naming, deterministic execution when practical, and a local README. Prefer adding depth over increasing raw file count. See `CONTRIBUTING.md`.

## Learning philosophy

**The goal is not to finish 1,111 files. The goal is to become the engineer who could build the systems represented by those files.**

The repository therefore moves from syntax to constraints: correctness, APIs, data, concurrency, failure, security, observability, performance, distributed behavior, and operational responsibility.

## Suggested milestones

Use the 20 levels as a formal curriculum, but revisit earlier levels whenever later topics expose a weak mental model. Production mastery is iterative.

---

Built as a long-term open-source learning system with a bias toward small examples, clear boundaries, and professional engineering habits.
