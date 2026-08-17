package main

import "fmt"

func summarizeGrpcConcept() (string, int) {
	topic := "Grpc Concept"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGrpcConcept()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
