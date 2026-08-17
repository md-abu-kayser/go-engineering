package main

import "fmt"

func summarizeGoVersionSelection() (string, int) {
	topic := "Go Version Selection"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGoVersionSelection()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
