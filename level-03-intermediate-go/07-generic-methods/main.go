package main

import "fmt"

func summarizeGenericMethods() (string, int) {
	topic := "Generic Methods"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGenericMethods()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
