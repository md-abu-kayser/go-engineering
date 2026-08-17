package main

import "fmt"

func summarizePointers() (string, int) {
	topic := "Pointers"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePointers()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
