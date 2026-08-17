package main

import "fmt"

func summarizeLeasePattern() (string, int) {
	topic := "Lease Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeLeasePattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
