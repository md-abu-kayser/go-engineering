package main

import "fmt"

func summarizeMessagePublisherPort() (string, int) {
	topic := "Message Publisher Port"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMessagePublisherPort()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
