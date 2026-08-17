package main

import "fmt"

func summarizeForLoop() (string, int) {
	topic := "For Loop"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeForLoop()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
