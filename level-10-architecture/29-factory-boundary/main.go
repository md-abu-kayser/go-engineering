package main

import "fmt"

func summarizeFactoryBoundary() (string, int) {
	topic := "Factory Boundary"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFactoryBoundary()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
