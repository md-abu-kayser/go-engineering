package main

import "fmt"

func summarizeSubprocessSupervision() (string, int) {
	topic := "Subprocess Supervision"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSubprocessSupervision()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
