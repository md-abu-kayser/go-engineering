package main

import "fmt"

func summarizeGmpModelAwareness() (string, int) {
	topic := "Gmp Model Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGmpModelAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
