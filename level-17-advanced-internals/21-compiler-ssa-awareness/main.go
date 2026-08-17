package main

import "fmt"

func summarizeCompilerSsaAwareness() (string, int) {
	topic := "Compiler Ssa Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCompilerSsaAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
