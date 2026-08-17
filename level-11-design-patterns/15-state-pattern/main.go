package main

import "fmt"

func summarizeStatePattern() (string, int) {
	topic := "State Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStatePattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
