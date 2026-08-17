package main

import "fmt"

func summarizeExamplesDirectory() (string, int) {
	topic := "Examples Directory"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeExamplesDirectory()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
