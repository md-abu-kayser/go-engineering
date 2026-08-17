package main

import "fmt"

func summarizeCliDeduplicator() (string, int) {
	topic := "Cli Deduplicator"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCliDeduplicator()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
