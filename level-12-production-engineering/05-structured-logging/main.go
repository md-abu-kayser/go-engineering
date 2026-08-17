package main

import "fmt"

func summarizeStructuredLogging() (string, int) {
	topic := "Structured Logging"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStructuredLogging()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
