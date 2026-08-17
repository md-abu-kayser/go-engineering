package main

import "fmt"

func summarizeFlyweightAwareness() (string, int) {
	topic := "Flyweight Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFlyweightAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
