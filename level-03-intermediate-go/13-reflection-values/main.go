package main

import "fmt"

func summarizeReflectionValues() (string, int) {
	topic := "Reflection Values"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeReflectionValues()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
