package main

import "fmt"

func summarizeArenaAwareness() (string, int) {
	topic := "Arena Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeArenaAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
