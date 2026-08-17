package main

import "fmt"

func summarizeDebugEndpointSecurity() (string, int) {
	topic := "Debug Endpoint Security"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDebugEndpointSecurity()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
