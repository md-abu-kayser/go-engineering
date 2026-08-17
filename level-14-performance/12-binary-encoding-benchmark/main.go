package main

import (
	"encoding/json"
	"fmt"
)

type lesson struct {
	Topic string `json:"topic"`
	Level int    `json:"level"`
}

func main() {
	b, err := json.Marshal(lesson{Topic: "Binary Encoding Benchmark", Level: 14})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
