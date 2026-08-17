package main

import "fmt"

func summarizeRangeLoop() (string, int) {
	topic := "Range Loop"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRangeLoop()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
