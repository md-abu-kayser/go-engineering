package main

import "fmt"

func summarizeLeaderElectionConcept() (string, int) {
	topic := "Leader Election Concept"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeLeaderElectionConcept()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
