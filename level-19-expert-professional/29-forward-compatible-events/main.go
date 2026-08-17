package main

import "fmt"

func summarizeForwardCompatibleEvents() (string, int) {
	topic := "Forward Compatible Events"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeForwardCompatibleEvents()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
