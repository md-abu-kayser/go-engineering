package main

import "fmt"

func summarizeReflectionTags() (string, int) {
	topic := "Reflection Tags"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeReflectionTags()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
