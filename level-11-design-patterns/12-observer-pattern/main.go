package main

import "fmt"

func summarizeObserverPattern() (string, int) {
	topic := "Observer Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeObserverPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
