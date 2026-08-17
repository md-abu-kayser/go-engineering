package main

import "fmt"

func summarizeRepositoryPattern() (string, int) {
	topic := "Repository Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRepositoryPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
