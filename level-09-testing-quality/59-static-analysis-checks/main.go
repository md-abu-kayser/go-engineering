package main

import "fmt"

func summarizeStaticAnalysisChecks() (string, int) {
	topic := "Static Analysis Checks"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStaticAnalysisChecks()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
