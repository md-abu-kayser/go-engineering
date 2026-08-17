package main

import "fmt"

func summarizeRbacAwareness() (string, int) {
	topic := "Rbac Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRbacAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
