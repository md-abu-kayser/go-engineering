package main

import "fmt"

func summarizeServiceRepositoryBoundary() (string, int) {
	topic := "Service Repository Boundary"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeServiceRepositoryBoundary()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
