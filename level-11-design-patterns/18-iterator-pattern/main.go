package main

import "fmt"

func summarizeIteratorPattern() (string, int) {
	topic := "Iterator Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeIteratorPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
