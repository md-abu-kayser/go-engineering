// Lesson 25: api/ Directory
//
// Goal: Use an api/ directory to hold the shared contract (request and
// response shapes) between a client and a server, independent of any
// one handler's implementation.
package main

import (
	"encoding/json"
	"fmt"

	"go-engineering/level-00-getting-started/25-api-directory/api"
)

// handleGreet is a stand-in for a real HTTP handler: it accepts an
// api.GreetRequest and returns an api.GreetResponse. Both types come
// from the shared api/ package, not from this file.
func handleGreet(req api.GreetRequest) api.GreetResponse {
	return api.GreetResponse{Message: "Hello, " + req.Name + "!"}
}

func main() {
	fmt.Println("=== api/ directory ===")
	fmt.Println("----------------------------------")

	req := api.GreetRequest{Name: "Gopher"}
	resp := handleGreet(req)

	reqJSON, _ := json.Marshal(req)
	respJSON, _ := json.Marshal(resp)

	fmt.Printf("Request  (api.GreetRequest)  : %s\n", reqJSON)
	fmt.Printf("Response (api.GreetResponse) : %s\n", respJSON)
	fmt.Println("\nSee api/openapi.yaml for the same contract described for non-Go clients.")
}
