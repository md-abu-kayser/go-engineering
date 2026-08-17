package main

import "fmt"

func summarizeGenericTypeParameters() (string, int) {
	topic := "Generic Type Parameters"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGenericTypeParameters()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
