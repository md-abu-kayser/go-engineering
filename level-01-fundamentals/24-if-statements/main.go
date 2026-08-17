package main

import "fmt"

func summarizeIfStatements() (string, int) {
	topic := "If Statements"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeIfStatements()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
