package main

import "fmt"

func summarizeReadinessProbe() (string, int) {
	topic := "Readiness Probe"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeReadinessProbe()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
