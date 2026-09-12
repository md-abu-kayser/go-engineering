// Lesson 30: Time Zones
//
// Goal: Combine focused standard-library packages to transform a value
// without introducing an application dependency.
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const lesson = "Time Zones"

type report struct {
	Words []string `json:"words"`
	Count int      `json:"count"`
}

func summarize(input string) (report, error) {
	words := strings.Fields(strings.ToLower(input))
	return report{Words: words, Count: len(words)}, nil
}

func main() {
	value, err := summarize("Go keeps standard-library tools close at hand")
	if err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Println(string(encoded))
}