package main

import "fmt"

func summarizeScope() (string, int) {
	topic := "Scope"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeScope()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
