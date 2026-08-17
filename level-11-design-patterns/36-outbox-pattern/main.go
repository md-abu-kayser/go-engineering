package main

import "fmt"

func summarizeOutboxPattern() (string, int) {
	topic := "Outbox Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeOutboxPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
