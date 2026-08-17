package main

import "fmt"

func summarizeInterfaceComposition() (string, int) {
	topic := "Interface Composition"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeInterfaceComposition()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
