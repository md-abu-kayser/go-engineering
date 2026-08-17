package main

import "fmt"

func summarizeStackTraces() (string, int) {
	topic := "Stack Traces"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStackTraces()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
