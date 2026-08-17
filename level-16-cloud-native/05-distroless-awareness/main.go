package main

import "fmt"

func summarizeDistrolessAwareness() (string, int) {
	topic := "Distroless Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDistrolessAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
