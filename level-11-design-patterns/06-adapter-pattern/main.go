package main

import "fmt"

func summarizeAdapterPattern() (string, int) {
	topic := "Adapter Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAdapterPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
