package main

import "fmt"

func summarizeRequestBodyHash() (string, int) {
	topic := "Request Body Hash"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRequestBodyHash()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
