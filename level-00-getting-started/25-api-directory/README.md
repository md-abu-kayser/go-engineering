# 25 — api/ Directory

## 🎯 Learning Objectives

- Use an `api/` directory to hold a shared request/response contract.
- Understand why separating "the shape of the data" from "the handler logic" is useful.
- Recognize that `api/` often holds **non-Go** files too (OpenAPI, protobuf).

## 📖 Concept

A top-level `api/` directory is a common convention for holding the **contract** of a service —
the request and response shapes any client needs to know about — separately from the handler
code that implements the behavior.

```
25-api-directory/
├── main.go              # (stand-in) handler logic
└── api/
    ├── types.go           # Go structs: the request/response contract
    └── openapi.yaml        # the same contract, described for non-Go consumers
```

### Why separate the contract from the handler?

- **Multiple handlers can share one contract.** An HTTP handler, a gRPC handler, and a test all
  reference the same `api.GreetRequest` / `api.GreetResponse` types — no duplication, no drift.
- **The contract can be versioned and reviewed independently** — changing a field name is a
  conscious edit to `api/types.go`, not something buried inside handler logic.
- **Non-Go consumers need a contract too.** `api/openapi.yaml` in this lesson expresses the same
  shape in a language-agnostic format — the kind of file a frontend team, another language's
  client generator, or API documentation tooling would consume directly.

### Real-world tools this pairs with

In production Go services, `api/` directories commonly hold `.proto` files (for gRPC, compiled
into Go code with `protoc`) or OpenAPI specs (compiled into Go server/client stubs with tools
like `oapi-codegen`) — the directory becomes the **single source of truth**, with generated Go
code living elsewhere, often gitignored and regenerated from the spec.

## 🔍 Code Walkthrough (`main.go`)

```go
func handleGreet(req api.GreetRequest) api.GreetResponse {
    return api.GreetResponse{Message: "Hello, " + req.Name + "!"}
}
```

Notice `handleGreet` only knows about `api.GreetRequest` / `api.GreetResponse` — it has no idea
whether it will eventually be wired up to `net/http`, gRPC, or called directly in a test. That
decoupling is the entire point of pulling the contract into its own package.

```go
reqJSON, _ := json.Marshal(req)
```

The `json:"name"` struct tags in `api/types.go` control exactly how these types serialize —
worth comparing against `api/openapi.yaml`'s `name`/`message` fields, which describe the same
shape for non-Go consumers.

## ▶️ How to Run

```bash
cd level-00-getting-started/25-api-directory
go run main.go
```

## ✅ Expected Output

```
=== api/ directory ===
----------------------------------
Request  (api.GreetRequest)  : {"name":"Gopher"}
Response (api.GreetResponse) : {"message":"Hello, Gopher!"}

See api/openapi.yaml for the same contract described for non-Go clients.
```

## 🧠 Key Takeaways

- `api/` centralizes the request/response contract, separate from handler implementation.
- Struct tags (`` `json:"name"` ``) control JSON serialization for Go API types.
- `api/` directories often hold non-Go contract files (OpenAPI, protobuf) alongside or instead of
  Go types, especially where code generation is involved.
- Keeping the contract in one place avoids duplicated, drifting type definitions across handlers.

## 🛠️ Try It Yourself

1. Add a `Language string` field (with a `json:"language,omitempty"` tag) to `GreetRequest`, and
   have `handleGreet` use it (reuse the logic from [lesson 13](../13-go-doc-command)'s
   `Greeter.Greet`).
2. Add the equivalent field to `api/openapi.yaml`'s request schema, and confirm both descriptions
   of the contract stay in sync.
3. Look up `oapi-codegen` or `protoc-gen-go` to see how real projects generate Go types directly
   from `api/` contract files instead of hand-writing them.

## ⚠️ Common Mistakes

- Defining the same request/response shape separately in multiple handler files, letting them
  quietly drift apart over time — centralizing in `api/` is specifically meant to prevent this.
- Forgetting JSON struct tags entirely — without them, Go's default field-name-based JSON keys
  (`Name`, `Message`) may not match what an OpenAPI spec or external client expects.
