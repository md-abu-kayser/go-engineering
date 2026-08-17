package main

import "fmt"

func summarizeContextDeadline() (string, int) {
	topic := "Context Deadline"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeContextDeadline()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
