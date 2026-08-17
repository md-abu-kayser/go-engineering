package main

import "fmt"

func summarizeSidecarAwareness() (string, int) {
	topic := "Sidecar Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSidecarAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
