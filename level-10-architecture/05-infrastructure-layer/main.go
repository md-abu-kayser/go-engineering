package main

import "fmt"

func summarizeInfrastructureLayer() (string, int) {
	topic := "Infrastructure Layer"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeInfrastructureLayer()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
