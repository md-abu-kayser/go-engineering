package main

import "fmt"

func summarizeCorrelationId() (string, int) {
	topic := "Correlation Id"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCorrelationId()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
