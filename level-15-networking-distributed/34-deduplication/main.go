package main

import "fmt"

func summarizeDeduplication() (string, int) {
	topic := "Deduplication"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDeduplication()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
