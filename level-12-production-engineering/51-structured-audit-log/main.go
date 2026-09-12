// Lesson 51: Structured Audit Log
//
// Goal: Validate configuration before work starts and emit a structured,
// deterministic operational record that can be inspected by people or tools.
package main

import (
	"encoding/json"
	"fmt"
)

const lesson = "Structured Audit Log"

type config struct {
	Service string
	Port    int
}

func (c config) Validate() error {
	if c.Service == "" || c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("invalid service configuration")
	}
	return nil
}

func main() {
	value := config{Service: "catalog", Port: 8080}
	if err := value.Validate(); err != nil {
		panic(err)
	}
	record, err := json.Marshal(map[string]any{"event": "service.ready", "service": value.Service, "port": value.Port})
	if err != nil {
		panic(err)
	}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Println(string(record))
}