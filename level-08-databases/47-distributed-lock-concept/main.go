package main

import "fmt"

func summarizeDistributedLockConcept() (string, int) {
	topic := "Distributed Lock Concept"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDistributedLockConcept()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
