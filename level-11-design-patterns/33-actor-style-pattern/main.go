package main

import "fmt"

func summarizeActorStylePattern() (string, int) {
	topic := "Actor Style Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeActorStylePattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
