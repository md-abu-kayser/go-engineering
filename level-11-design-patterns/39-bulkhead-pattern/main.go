package main

import "fmt"

func summarizeBulkheadPattern() (string, int) {
	topic := "Bulkhead Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBulkheadPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
