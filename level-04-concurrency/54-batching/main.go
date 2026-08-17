package main

import "fmt"

func summarizeBatching() (string, int) {
	topic := "Batching"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBatching()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
