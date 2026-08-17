package main

import "fmt"

func summarizeCleanRepositoryPort() (string, int) {
	topic := "Clean Repository Port"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCleanRepositoryPort()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
