package main

import "fmt"

func summarizeEscapeAnalysisExample() (string, int) {
	topic := "Escape Analysis Example"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeEscapeAnalysisExample()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
