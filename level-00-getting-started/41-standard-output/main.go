package main

import "fmt"

func summarizeStandardOutput() (string, int) {
	topic := "Standard Output"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStandardOutput()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
