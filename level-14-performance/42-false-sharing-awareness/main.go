package main

import "fmt"

func summarizeFalseSharingAwareness() (string, int) {
	topic := "False Sharing Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFalseSharingAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
