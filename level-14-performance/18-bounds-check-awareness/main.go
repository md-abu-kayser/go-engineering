package main

import "fmt"

func summarizeBoundsCheckAwareness() (string, int) {
	topic := "Bounds Check Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBoundsCheckAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
