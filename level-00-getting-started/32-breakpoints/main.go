package main

import "fmt"

func summarizeBreakpoints() (string, int) {
	topic := "Breakpoints"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBreakpoints()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
