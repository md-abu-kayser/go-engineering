# Learning Roadmap — 0% → 100%

> The roadmap is a competency map, not a race. Finish the concepts, build something with them, prove that you understand them, then move forward.

## Roadmap Philosophy

This roadmap describes the intended progression from absolute beginner to professional Go engineering.

The percentage markers are **curriculum milestones**, not measurements of intelligence, seniority, or guaranteed mastery.

The goal is not:

```text
read more files
```

The goal is:

```text
understand
   ↓
experiment
   ↓
practice
   ↓
debug
   ↓
test
   ↓
explain
   ↓
build
   ↓
integrate
   ↓
operate
```

The repository therefore treats mastery as demonstrated capability rather than file completion.

---

# How to Read This Roadmap

A learner should normally move forward when they can:

1. explain the concept without copying the example;
2. predict the behavior of a small program before running it;
3. modify the example safely;
4. rebuild a small variation from memory;
5. identify common failure modes;
6. write or understand tests for the behavior;
7. debug a broken implementation;
8. explain the trade-offs of the approach;
9. connect the concept to a larger Go system;
10. explain where the technique belongs—and where it does not belong—in production.

A learner does **not** need perfect mastery before continuing.

The correct pattern is:

```text
understand enough
      ↓
practice
      ↓
move forward
      ↓
revisit when complexity exposes gaps
```

---

# Competency Evidence Model

Every major milestone should produce evidence.

```text
Level 1 — Recognition
Can identify the concept.

Level 2 — Explanation
Can explain the concept in simple language.

Level 3 — Modification
Can safely modify an existing implementation.

Level 4 — Construction
Can build a small implementation independently.

Level 5 — Debugging
Can diagnose common failures.

Level 6 — Testing
Can prove important behavior with tests.

Level 7 — Integration
Can use the concept inside a larger system.

Level 8 — Judgment
Can explain trade-offs and alternative designs.

Level 9 — Production Thinking
Can reason about security, failure, observability, and operations.
```

A learner does not need every lesson to reach Level 9. The maturity required depends on the topic.

---

# Learning Loop

Use this loop throughout the curriculum:

```text
READ
  ↓
UNDERSTAND
  ↓
RUN
  ↓
OBSERVE
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
CONNECT TO A REAL SYSTEM
  ↓
REVIEW TRADE-OFFS
```

The most important step is the one immediately after understanding:

> **Do something with the concept instead of only reading about it.**

---

# Phase I — Language Fluency

## 0–5% · Level 00 — Getting Started

### Objective

Remove environmental friction and establish a healthy Go development workflow.

### Core skills

- Install and verify Go
- Understand `go version`
- Understand the Go toolchain
- `go run`
- `go build`
- `go test`
- `go fmt`
- `go vet`
- `go mod`
- `go.mod`
- packages
- imports
- `package main`
- `func main()`
- source layout
- documentation comments
- terminal basics
- editor integration
- basic debugging
- environment variables
- exit codes

### Mental model

```text
source code
    ↓
package
    ↓
module
    ↓
Go toolchain
    ↓
binary / tests
```

### Evidence of readiness

A learner should be able to create a tiny Go program from an empty directory, explain its module structure, run it, build it, format it, test it, and diagnose a basic compiler error.

### Milestone project

A small command-line utility that builds as a single binary.

### Recommended proof

```text
source
+ README
+ test
+ build command
+ execution example
```

---

## 5–12% · Level 01 — Fundamentals

### Objective

Build strong fluency with Go's basic syntax and semantics.

### Core skills

- variables
- constants
- zero values
- strings
- runes
- bytes
- numeric types
- booleans
- operators
- conditionals
- `switch`
- `for`
- functions
- parameters
- return values
- multiple returns
- scope
- shadowing
- closures
- packages
- imports
- basic error handling
- documentation comments

### Learning progression

```text
expression
   ↓
statement
   ↓
function
   ↓
multiple functions
   ↓
multiple files
   ↓
package
```

### Evidence of readiness

The learner can solve small deterministic problems without copying syntax from the lesson.

### Milestone project

A deterministic CLI calculator or data processor.

---

## 12–20% · Level 02 — Core Go

### Objective

Move from basic syntax into Go's core composition model.

### Core skills

- arrays
- slices
- maps
- strings
- structs
- pointers
- methods
- interfaces
- embedding
- composition
- custom errors
- error wrapping
- type assertions
- type switches
- package design
- zero-value design

### Key principle

> Prefer simple composition over unnecessary abstraction.

### Required reasoning

For each major feature, understand:

```text
What is it?
Why does it exist?
How does it behave?
What can go wrong?
When should I use it?
When should I avoid it?
```

### Milestone project

A small in-memory domain application with:

- validation;
- business rules;
- errors;
- tests;
- package boundaries.

---

## 20–25% · Level 03 — Intermediate Go

### Objective

Understand modern language features and reusable package design.

### Core skills

- generics
- type constraints
- advanced interfaces
- reflection
- embedding
- package architecture
- modules
- dependency management
- semantic versioning concepts
- exported vs unexported identifiers
- API design
- zero-value-friendly design
- functional options
- error wrapping
- sentinel errors
- custom error types

### Engineering exercise

Design a reusable package and justify every public API.

### Readiness test

You should be able to explain:

```text
Why is this abstraction necessary?
Why is it not smaller?
Why is it not larger?
Why is this API exported?
Why should the consumer depend on it?
```

### Milestone

A reusable Go package with documentation and tests.

---

# Phase II — Production Programming

## 25–33% · Level 04 — Concurrency

### Objective

Treat concurrency as a correctness and lifecycle problem—not merely a syntax feature.

### Core skills

- goroutines
- channels
- buffered channels
- unbuffered channels
- `select`
- timeouts
- tickers
- worker pools
- fan-out
- fan-in
- pipelines
- mutexes
- `RWMutex`
- atomics
- `WaitGroup`
- `Once`
- condition synchronization
- semaphores
- bounded concurrency
- backpressure
- context cancellation
- race detection
- deadlock prevention
- goroutine lifecycle
- concurrent caches

### Learning progression

```text
one goroutine
      ↓
multiple goroutines
      ↓
synchronization
      ↓
channels
      ↓
select
      ↓
cancellation
      ↓
worker pools
      ↓
backpressure
      ↓
graceful shutdown
      ↓
production concurrency
```

### Failure labs

Practice deliberately reproducing:

- data race;
- deadlock;
- goroutine leak;
- blocked sender;
- blocked receiver;
- unbounded queue;
- shutdown race;
- cancellation failure.

### Milestone project

A cancellable worker pipeline with bounded concurrency, metrics, graceful shutdown, and tests.

---

## 33–38% · Level 05 — Standard Library Mastery

### Objective

Know what Go already provides before introducing dependencies.

### Coverage

- `fmt`
- `strings`
- `bytes`
- `unicode`
- `unicode/utf8`
- `strconv`
- `sort`
- `slices`
- `maps`
- `cmp`
- `math`
- `math/big`
- `math/rand`
- `crypto/*`
- `encoding/*`
- `io`
- `io/fs`
- `os`
- `path`
- `path/filepath`
- `bufio`
- `archive/*`
- `compress/*`
- `time`
- `regexp`
- `net`
- `net/url`
- `net/http`
- `net/http/httptest`
- `sync`
- `sync/atomic`
- `context`
- `errors`
- `log`
- `log/slog`
- `flag`
- `database/sql`
- `embed`
- `reflect`
- `runtime`
- `debug`
- `testing`

### Milestone

Solve a meaningful utility problem using the standard library before reaching for an external dependency.

---

## 38–43% · Level 06 — CLI & Systems Programming

### Objective

Build programs that interact directly with the operating environment.

### Coverage

- CLI design
- argument parsing
- flags
- environment variables
- configuration
- files
- directories
- permissions
- processes
- subprocesses
- signals
- shutdown
- stdin/stdout/stderr
- exit codes
- structured logging
- JSON configuration
- temporary files
- lock files
- Unix concepts
- cross-platform behavior
- cross-compilation

### Failure labs

- invalid configuration;
- missing file;
- permission error;
- interrupted process;
- child process failure;
- partial write;
- invalid environment variable.

### Milestone project

A production-style CLI with configuration, validation, logging, structured output, and graceful shutdown.

---

# Phase III — Backend Engineering

## 43–50% · Level 07 — Web Development

### Objective

Build reliable HTTP services using Go's native HTTP model.

### Coverage

- HTTP fundamentals
- request lifecycle
- routing
- handlers
- middleware
- REST semantics
- JSON APIs
- validation
- pagination
- filtering
- sorting
- cookies
- sessions
- authentication
- authorization
- CORS
- CSRF concepts
- file uploads
- downloads
- streaming
- timeouts
- graceful shutdown
- request IDs
- structured logging
- API errors
- API versioning
- API testing

### Architecture progression

```text
http.Handler
    ↓
handler composition
    ↓
router
    ↓
middleware
    ↓
application service
    ↓
domain logic
    ↓
repository
    ↓
database / external service
```

### Milestone project

A production-oriented REST API.

Required evidence:

- API documentation;
- validation;
- tests;
- error model;
- graceful shutdown;
- configuration;
- structured logging.

---

## 50–56% · Level 08 — Databases & Data Access

### Objective

Treat persistence as a deliberate architectural boundary.

### Coverage

- relational modeling
- SQL
- PostgreSQL
- `database/sql`
- connection pools
- prepared statements
- transactions
- isolation
- locking
- indexes
- query design
- migrations
- repository design
- unit-of-work concepts
- caching
- Redis concepts
- NoSQL trade-offs
- pagination
- optimistic locking
- database observability

### Architecture

```text
HTTP
 ↓
Application Service
 ↓
Domain
 ↓
Repository Interface
 ↓
PostgreSQL Implementation
```

### Failure labs

- transaction rollback;
- lock contention;
- timeout;
- duplicate operation;
- stale data;
- connection exhaustion;
- migration failure.

### Milestone project

A transactional API with repository boundaries and integration tests.

---

## 56–61% · Level 09 — Testing & Quality

### Objective

Treat testing as part of design.

### Coverage

- `testing.T`
- subtests
- table-driven tests
- examples
- test helpers
- fixtures
- isolation
- integration tests
- HTTP tests
- database tests
- benchmarks
- fuzz tests
- race detection
- coverage
- golden files
- mocks
- fakes
- stubs
- contract tests
- property-oriented testing
- deterministic tests

### Testing progression

```text
function
 ↓
package
 ↓
integration
 ↓
system
 ↓
performance
 ↓
production verification
```

### Milestone

Design a package whose important behavior can be tested without a real network or database.

---

# Phase IV — Software Architecture

## 61–66% · Level 10 — Go Architecture

### Objective

Make system boundaries intentional.

### Coverage

- layered architecture
- Clean Architecture
- Hexagonal Architecture
- Ports and Adapters
- dependency inversion
- dependency injection
- domain boundaries
- application services
- repositories
- use cases
- DTOs
- mappers
- domain models
- package boundaries
- `internal`
- public APIs
- modular monoliths
- architecture tests

### Architecture reasoning model

```text
problem
   ↓
constraints
   ↓
candidate designs
   ↓
trade-offs
   ↓
decision
   ↓
implementation
   ↓
verification
```

### Milestone project

Refactor an API into independently testable domain, application, and infrastructure layers.

---

## 66–70% · Level 11 — Design Patterns

### Objective

Recognize recurring design problems and select the simplest appropriate solution.

### Coverage

- Factory
- Builder
- Strategy
- Adapter
- Decorator
- Observer
- Command
- State
- Proxy
- Facade
- Repository
- Unit of Work
- Dependency Injection
- Functional Options
- Middleware
- Pipeline
- Worker Pool
- Event Bus
- Retry
- Circuit Breaker
- Bulkhead

### Rule

Never introduce a pattern merely because it exists.

A pattern should answer:

```text
What problem does this solve?
Why is the simple solution insufficient?
What complexity does it introduce?
```

### Milestone

Choose and defend a design pattern only when a real constraint justifies it.

---

## 70–75% · Level 12 — Production Engineering

### Objective

Make software observable, operable, and resilient.

### Coverage

- configuration
- structured logging
- log levels
- request IDs
- correlation IDs
- metrics
- tracing
- health checks
- readiness
- liveness
- graceful shutdown
- timeouts
- retries
- backoff
- circuit breakers
- rate limiting
- load shedding
- resource limits
- dependency health
- secrets handling
- operational diagnostics

### Operational questions

```text
How do we know it is healthy?
How do we know it is failing?
How do we observe it?
How do we stop it safely?
How do we recover it?
How do we limit its damage?
```

### Milestone project

An observable service with operational health, structured logs, metrics, tracing, and graceful shutdown.

---

# Phase V — Security, Performance, and Systems

## 75–80% · Level 13 — Security

### Objective

Treat security as a cross-cutting engineering responsibility.

### Coverage

- authentication
- authorization
- password hashing
- session security
- JWT concepts
- OAuth concepts
- OIDC concepts
- API keys
- input validation
- output encoding
- CSRF
- CORS
- SQL injection prevention
- command injection prevention
- SSRF concepts
- path traversal
- secure file handling
- secret management
- TLS concepts
- certificate validation
- secure headers
- rate limiting
- audit logging
- security testing
- dependency security
- threat modeling basics

### Security learning model

```text
identify asset
   ↓
identify threat
   ↓
understand vulnerability
   ↓
reproduce safely
   ↓
mitigate
   ↓
test mitigation
   ↓
prevent regression
```

### Milestone

Audit a deliberately vulnerable example and harden it systematically.

---

## 80–84% · Level 14 — Performance & Optimization

### Objective

Optimize only after establishing a measurable reason.

### Coverage

- benchmarks
- CPU profiling
- memory profiling
- allocations
- escape analysis
- garbage collection
- heap behavior
- stack behavior
- contention
- parallelism
- I/O performance
- network performance
- database performance
- cache behavior
- profiling-driven optimization

### Performance workflow

```text
correctness
   ↓
measure
   ↓
profile
   ↓
identify bottleneck
   ↓
change one thing
   ↓
benchmark
   ↓
compare
   ↓
document trade-off
   ↓
protect with regression test
```

### Milestone

Use profiling evidence to improve a measurable bottleneck and prove the improvement.

---

## 84–89% · Level 15 — Networking & Distributed Systems

### Objective

Understand communication, partial failure, and coordination across processes.

### Coverage

- TCP
- UDP
- sockets
- DNS concepts
- HTTP
- TLS concepts
- WebSockets
- streaming
- RPC
- gRPC
- connection management
- timeouts
- retries
- idempotency
- message queues
- publish/subscribe
- events
- serialization
- distributed locking
- consensus concepts
- leader election concepts
- failure detection
- network partitions
- eventual consistency
- distributed tracing

### Central principle

> Distributed systems are primarily failure-management systems.

### Milestone project

A small distributed workflow with failure-aware communication and explicit timeout/retry behavior.

---

## 89–93% · Level 16 — Microservices & Cloud Native Go

### Objective

Understand what changes when a process becomes one component of a larger system.

### Coverage

- service boundaries
- modular monoliths
- microservices
- service discovery
- configuration
- containers
- Docker
- Kubernetes concepts
- health probes
- deployments
- service accounts
- secrets
- horizontal scaling
- load balancing
- resilience
- zero-downtime deployment
- rolling deployments
- blue/green concepts
- autoscaling concepts
- cloud-native observability

### Important decision

The learner must also understand:

> When should a system **not** be split into microservices?

### Milestone

Design a multi-service system and explain its operational, security, and failure modes.

---

## 93–97% · Level 17 — Advanced Go Engineering

### Objective

Study Go and its runtime with deeper engineering mental models.

### Coverage

- Go memory model
- goroutine scheduling
- runtime behavior
- garbage collector concepts
- scheduler concepts
- stack growth
- escape analysis
- interfaces internally
- reflection internals
- runtime diagnostics
- `pprof`
- execution tracing
- advanced synchronization
- atomic memory-ordering concepts
- lock-free concepts
- runtime metrics
- compiler behavior
- assembly awareness
- performance engineering

### Milestone

Explain a surprising performance or concurrency behavior from runtime principles rather than guesswork.

---

# Phase VI — Professional Mastery

## 97–99% · Level 18 — Real-World Projects

### Objective

Integrate multiple levels into coherent systems.

### Project families

- production API
- authentication service
- URL shortener
- job queue
- notification service
- file-processing system
- real-time service
- microservices platform
- distributed system
- expert capstone

### Project maturity requirements

Each substantial project should answer:

```text
What problem are we solving?
Who are the users?
What are the requirements?
What are the constraints?
How is the system designed?
How is data stored?
How are failures handled?
How is security implemented?
How is it tested?
How is it observed?
How is it deployed?
How is it maintained?
```

### Milestone

Complete one system with an architecture document, tests, operational considerations, and engineering trade-offs.

---

## 99–100% · Level 19 — Expert / Professional Go

### Objective

Develop engineering judgment.

At this stage the question changes from:

> “How do I write this Go code?”

to:

> “What system should exist, what trade-offs does it make, and how will we operate it safely?”

### Focus

- large-scale architecture
- system design
- capacity planning
- reliability engineering
- availability
- fault tolerance
- scalability
- backpressure
- performance budgets
- observability architecture
- incident response
- operational maturity
- API evolution
- backward compatibility
- dependency strategy
- technical debt
- maintainability
- code ownership
- architecture governance
- production debugging
- failure analysis
- long-term system evolution
- technical leadership
- architecture decision records

### Final milestone

Produce a capstone system and engineering write-up that explains:

```text
problem
architecture
constraints
trade-offs
failure modes
security
testing
performance
observability
operations
future evolution
```

---

# Cross-Level Mastery Model

A strong learner should repeatedly move through:

```text
Level N
  ↓
Understand
  ↓
Practice
  ↓
Test
  ↓
Debug
  ↓
Build
  ↓
Integrate into Level N+1
  ↓
Revisit Level N with new context
```

Learning is therefore intentionally non-linear.

A learner may return to:

```text
fundamentals
concurrency
errors
interfaces
testing
```

many times as later systems expose deeper questions.

---

# Project Integration Map

Projects exist to integrate learning.

```text
Level 02
   ↓
data structures + interfaces
   ↓
Level 04
   ↓
concurrency
   ↓
Level 05
   ↓
standard library
   ↓
Level 07
   ↓
HTTP
   ↓
Level 08
   ↓
database
   ↓
Level 09
   ↓
testing
   ↓
Level 10
   ↓
architecture
   ↓
Level 12
   ↓
observability
   ↓
Level 13
   ↓
security
   ↓
Level 14
   ↓
performance
   ↓
Level 15
   ↓
distributed behavior
```

The resulting project becomes evidence that multiple capabilities can be integrated into one system.

---

# Portfolio Milestones

| Milestone | Expected evidence                            |
| --------- | -------------------------------------------- |
| 20%       | Strong command of Go fundamentals            |
| 40%       | CLI + concurrency examples                   |
| 50%       | Production-style REST API                    |
| 60%       | Database-backed service + tests              |
| 70%       | Explicit architecture implementation         |
| 80%       | Security + profiling work                    |
| 90%       | Distributed/cloud-native project             |
| 100%      | Capstone architecture + engineering write-up |

The strongest public evidence includes:

```text
code
+
tests
+
benchmarks
+
architecture diagrams
+
failure experiments
+
trade-offs
+
documentation
```

Not merely file count.

---

# Interview & Career Readiness

The roadmap should gradually produce the ability to answer:

### Beginner

- What is a package?
- What is a slice?
- What is a map?
- What is an interface?
- What is an error?

### Intermediate

- Why are slices different from arrays?
- How does interface composition work?
- How do you design errors?
- How do you make code testable?
- How do you avoid goroutine leaks?

### Advanced

- How would you diagnose a data race?
- How would you design bounded concurrency?
- How would you handle dependency timeouts?
- How would you structure a larger service?
- When would you introduce an abstraction?

### Expert

- What are the system's failure modes?
- Where is the bottleneck?
- What are the scaling constraints?
- What is the operational model?
- What trade-offs did you intentionally accept?
- What would you change at 10× traffic?
- What would you change at 100× traffic?
- What would you refuse to optimize?

---

# Competency Gate Before Moving Forward

Before leaving a major level, ask:

```text
[ ] Can I explain the main concepts?
[ ] Can I modify the examples?
[ ] Can I build a small example without copying?
[ ] Can I identify common mistakes?
[ ] Can I write or understand tests?
[ ] Can I debug a broken implementation?
[ ] Can I explain trade-offs?
[ ] Can I connect the concept to production?
[ ] Can I explain when NOT to use it?
```

Do not require every checkbox for every lesson. Use the gate proportionally to the topic.

---

# Recommended Revision Cycle

Learning is not finished when the first implementation works.

Use:

```text
Version 1 — make it work
      ↓
Version 2 — make it clear
      ↓
Version 3 — test it
      ↓
Version 4 — inspect failure modes
      ↓
Version 5 — improve design
      ↓
Version 6 — measure if performance matters
      ↓
Version 7 — document the trade-offs
```

This is especially valuable for:

- concurrency;
- APIs;
- database access;
- architecture;
- security;
- performance;
- distributed systems.

---

# Final Definition of 100%

100% does **not** mean:

```text
I know every Go feature.
```

100% means the learner has developed enough capability to:

```text
understand unfamiliar Go code
        ↓
reason about behavior
        ↓
design a solution
        ↓
implement it
        ↓
test it
        ↓
debug it
        ↓
secure it
        ↓
measure it
        ↓
operate it
        ↓
explain the trade-offs
```

The ultimate goal is not syntax mastery.

> **The ultimate goal is professional engineering judgment expressed through Go.**
