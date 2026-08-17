package main

import "fmt"

func summarizeStackCopyingAwareness() (string, int) {
	topic := "Stack Copying Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStackCopyingAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
