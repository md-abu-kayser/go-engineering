package main

import "fmt"

func summarizePresignedUrlConcept() (string, int) {
	topic := "Presigned Url Concept"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePresignedUrlConcept()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
