package main

import "fmt"

func summarizeTraceContext() (string, int) {
	topic := "Trace Context"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeTraceContext()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
