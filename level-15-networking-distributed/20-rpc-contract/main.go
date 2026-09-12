// Lesson 20: Rpc Contract
//
// Goal: Put a message in an explicit envelope so a receiver can correlate,
// deduplicate, and evolve data independently of transport details.
package main

import (
	"encoding/json"
	"fmt"
)

const lesson = "Rpc Contract"

type envelope struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Version int    `json:"version"`
	Payload string `json:"payload"`
}

func main() {
	message := envelope{ID: "evt-42", Type: "order.created", Version: 1, Payload: "order-7"}
	encoded, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Println(string(encoded))
}