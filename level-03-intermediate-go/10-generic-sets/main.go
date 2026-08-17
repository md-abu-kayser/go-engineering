package main

import "fmt"

func summarizeGenericSets() (string, int) {
	topic := "Generic Sets"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGenericSets()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
