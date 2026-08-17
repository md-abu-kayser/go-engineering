# Learning Roadmap — 0% → 100%

> The roadmap is a competency map, not a race. Finish the concepts, build something with them, then move forward.

## How to read this roadmap

Each level contains focused lessons. The percentages below are **curriculum milestones**, not a claim that mastery can be measured precisely by file count.

A learner should move forward when they can:

1. explain the concept without copying the example;
2. modify the example safely;
3. build a small variation from memory;
4. identify common failure modes;
5. explain where the technique belongs in a production system.

---

## Phase I — Language Fluency

### 0–5% · Level 00 — Getting Started

Learn the toolchain and establish a healthy workflow.

**Core skills**

- Install and verify Go
- `go run`, `go build`, `go test`
- modules and `go.mod`
- packages and imports
- source layout
- documentation comments
- debugging basics
- environment and command-line behavior

**Milestone project:** a small command-line utility that builds as a single binary.

---

### 5–12% · Level 01 — Fundamentals

Build complete fluency with Go's basic syntax and semantics.

**Core skills**

- variables and constants
- strings, runes, bytes, numerics
- operators
- control flow
- functions
- scope
- closures
- packages
- basic error handling

**Milestone project:** a deterministic CLI calculator/data processor.

---

### 12–20% · Level 02 — Core Go

Move from syntax to composition.

**Core skills**

- arrays and slices
- maps
- structs
- pointers
- methods
- interfaces
- composition
- custom errors
- functional-style utilities where appropriate

**Milestone project:** a small in-memory domain application with validation and tests.

---

### 20–25% · Level 03 — Intermediate Go

Understand the language features that distinguish modern Go code.

**Core skills**

- generics
- reflection
- embedding
- advanced interfaces
- package design
- modules
- versioning
- dependency management

**Milestone:** design a reusable package without over-abstracting it.

---

## Phase II — Production Programming

### 25–33% · Level 04 — Concurrency

Concurrency becomes a design tool rather than a syntax trick.

**Core skills**

- goroutines
- channels
- `select`
- worker pools
- pipelines
- backpressure
- mutexes
- atomics
- `sync.Once`
- cancellation
- context propagation
- race conditions
- deadlocks

**Milestone project:** a cancellable worker pipeline with bounded concurrency.

---

### 33–38% · Level 05 — Standard Library Mastery

Learn to recognize when the standard library already solves the problem.

**Coverage**

`fmt`, `strings`, `bytes`, `io`, `os`, `path`, `filepath`, `time`, `regexp`, `encoding/json`, `encoding/csv`, `net/http`, `net/url`, `sync`, `context`, `sort`, `slices`, `maps`, and more.

**Milestone:** solve a real utility problem using the standard library before reaching for a dependency.

---

### 38–43% · Level 06 — CLI & Systems Programming

Build programs that interact directly with the operating environment.

**Coverage**

- CLI design
- configuration
- environment variables
- files and directories
- processes
- signals
- subprocesses
- exit codes
- logging
- cross-compilation

**Milestone project:** a production-style CLI with configuration, validation, logging, and graceful shutdown.

---

## Phase III — Backend Engineering

### 43–50% · Level 07 — Web Development

Build reliable HTTP services.

**Coverage**

- HTTP fundamentals
- routing
- middleware
- REST semantics
- validation
- cookies
- sessions
- authentication
- authorization
- JSON APIs
- error responses
- API testing

**Milestone project:** a production-oriented REST API.

---

### 50–56% · Level 08 — Databases & Data Access

Learn to make persistence a deliberate boundary.

**Coverage**

- relational modeling
- SQL
- PostgreSQL
- transactions
- isolation
- connection pooling
- repositories
- migrations
- caching
- Redis concepts
- NoSQL trade-offs

**Milestone project:** a transactional API with repository boundaries and integration tests.

---

### 56–61% · Level 09 — Testing & Quality

Testing becomes part of design rather than an afterthought.

**Coverage**

- table-driven tests
- subtests
- examples
- integration tests
- benchmarks
- fuzzing
- coverage
- mocks/stubs/fakes
- test fixtures
- deterministic tests

**Milestone:** design a package whose behavior can be tested without a real network or database.

---

## Phase IV — Software Architecture

### 61–66% · Level 10 — Go Architecture

Learn to make boundaries explicit.

**Coverage**

- layered architecture
- dependency inversion
- clean architecture
- hexagonal architecture
- domain-oriented design
- ports and adapters
- application services
- infrastructure boundaries

**Milestone project:** refactor an API into independently testable domain/application/infrastructure layers.

---

### 66–70% · Level 11 — Design Patterns

Patterns are learned as trade-offs, not templates.

**Coverage**

- factory
- builder
- strategy
- adapter
- observer
- repository
- dependency injection
- functional options
- middleware chains
- Go-specific patterns

**Milestone:** choose a pattern only when it improves a real design constraint.

---

### 70–75% · Level 12 — Production Engineering

Make systems observable and operable.

**Coverage**

- configuration management
- structured logging
- metrics
- tracing
- health checks
- readiness/liveness
- graceful shutdown
- operational endpoints
- incident-oriented diagnostics

**Milestone project:** an observable service that explains its health and behavior to operators.

---

## Phase V — Security, Performance, and Systems

### 75–80% · Level 13 — Security

Treat security as architecture.

**Coverage**

- authentication
- authorization
- password hashing
- JWT considerations
- OAuth concepts
- input validation
- secure HTTP behavior
- common web vulnerabilities
- secrets management
- threat modeling basics

**Milestone:** audit a deliberately vulnerable example and harden it systematically.

---

### 80–84% · Level 14 — Performance & Optimization

Measure before optimizing.

**Coverage**

- benchmarks
- CPU profiling
- memory profiling
- allocations
- escape analysis
- garbage collection
- contention
- throughput
- latency
- concurrency tuning

**Milestone:** use profiling evidence to improve a measurable bottleneck.

---

### 84–89% · Level 15 — Networking & Distributed Systems

Understand communication beyond a single process.

**Coverage**

- TCP
- UDP
- sockets
- WebSockets
- RPC
- gRPC concepts
- service-to-service communication
- queues
- events
- retries
- timeouts
- idempotency

**Milestone project:** a small distributed workflow with failure-aware communication.

---

### 89–93% · Level 16 — Microservices & Cloud Native Go

Learn what changes when a process becomes one component of a larger system.

**Coverage**

- service decomposition
- Docker
- Kubernetes concepts
- service discovery
- configuration
- resilience
- deployment concerns
- rolling upgrades
- infrastructure boundaries

**Milestone:** design a multi-service system and explain its operational failure modes.

---

### 93–97% · Level 17 — Advanced Go Engineering

Study the runtime and the language at a deeper level.

**Coverage**

- memory model
- scheduler concepts
- stack growth
- garbage collection behavior
- runtime internals
- advanced concurrency
- compiler/runtime interactions
- performance engineering

**Milestone:** explain a performance or concurrency behavior from runtime principles rather than guesswork.

---

## Phase VI — Professional Mastery

### 97–99% · Level 18 — Real-World Projects

Integrate multiple levels into production-style systems.

Project families include:

- production API
- authentication service
- job platform
- file-processing pipeline
- real-time system
- URL shortener
- notification platform
- distributed demo
- multi-tenant SaaS
- expert capstone

The goal is not project count. The goal is coherent systems thinking.

---

### 99–100% · Level 19 — Expert / Professional Go

At this stage the question changes from:

> “How do I write this Go code?”

to:

> “What system should exist, what trade-offs does it make, and how will we operate it safely at scale?”

**Focus**

- large-scale architecture
- reliability engineering
- fault tolerance
- scalability
- observability
- cost awareness
- maintainability
- incident response
- technical leadership
- engineering decision records
- system design

---

## Recommended learning loop

```text
READ
  ↓
RUN
  ↓
MODIFY
  ↓
BREAK IT
  ↓
DEBUG
  ↓
TEST
  ↓
EXPLAIN
  ↓
BUILD WITHOUT THE EXAMPLE
  ↓
CONNECT IT TO A REAL SYSTEM
```

That loop matters more than moving quickly through the file tree.

---

## Portfolio milestones

A learner can turn the repository into a public engineering portfolio by publishing progress around these milestones:

| Milestone | Evidence to publish                          |
| --------- | -------------------------------------------- |
| 20%       | Strong command of Go fundamentals            |
| 40%       | CLI + concurrency examples                   |
| 50%       | Production-style REST API                    |
| 60%       | Database-backed service + tests              |
| 70%       | Clean architecture implementation            |
| 80%       | Security + profiling work                    |
| 90%       | Distributed/cloud-native project             |
| 100%      | Capstone architecture + engineering write-up |

The repository is most impressive when it shows **reasoning, trade-offs, tests, measurements, and production thinking**, not merely a large file count.
