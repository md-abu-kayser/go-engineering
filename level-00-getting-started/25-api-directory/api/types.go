// Package api holds the shared data contracts for this lesson's example
// service — the request/response shapes that any client and server code
// both need to agree on. This is one common use of a top-level api/
// directory: a single source of truth for the "shape" of an API,
// independent of any one handler's implementation.
package api

// GreetRequest is what a client sends to ask for a greeting.
type GreetRequest struct {
	Name string `json:"name"`
}

// GreetResponse is what the server sends back.
type GreetResponse struct {
	Message string `json:"message"`
}
