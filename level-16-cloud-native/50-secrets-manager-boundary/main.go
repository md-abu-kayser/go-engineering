package main

import "fmt"

func summarizeSecretsManagerBoundary() (string, int) {
	topic := "Secrets Manager Boundary"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSecretsManagerBoundary()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
