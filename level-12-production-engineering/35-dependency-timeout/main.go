package main

import "fmt"

func summarizeDependencyTimeout() (string, int) {
	topic := "Dependency Timeout"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDependencyTimeout()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
