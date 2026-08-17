package main

import "fmt"

func summarizeCopyAvoidance() (string, int) {
	topic := "Copy Avoidance"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCopyAvoidance()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
