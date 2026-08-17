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
	b, err := json.Marshal(lesson{Topic: "Cli Json Tool", Level: 18})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
