package main

import "fmt"

func summarizeVersionEndpoint() (string, int) {
	topic := "Version Endpoint"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeVersionEndpoint()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
