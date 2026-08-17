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
	b, err := json.Marshal(lesson{Topic: "Json Unmarshal Preview", Level: 2})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
