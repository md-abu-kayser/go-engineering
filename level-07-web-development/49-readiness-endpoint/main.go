package main

import "fmt"

func summarizeReadinessEndpoint() (string, int) {
	topic := "Readiness Endpoint"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeReadinessEndpoint()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
