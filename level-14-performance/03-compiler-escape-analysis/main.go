package main

import "fmt"

func summarizeCompilerEscapeAnalysis() (string, int) {
	topic := "Compiler Escape Analysis"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCompilerEscapeAnalysis()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
