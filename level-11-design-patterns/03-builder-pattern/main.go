package main

import "fmt"

func summarizeBuilderPattern() (string, int) {
	topic := "Builder Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBuilderPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
