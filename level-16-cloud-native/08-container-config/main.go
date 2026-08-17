package main

import "fmt"

func summarizeContainerConfig() (string, int) {
	topic := "Container Config"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeContainerConfig()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
