package main

import "fmt"

func summarizeInterfaceHeaderAwareness() (string, int) {
	topic := "Interface Header Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeInterfaceHeaderAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
