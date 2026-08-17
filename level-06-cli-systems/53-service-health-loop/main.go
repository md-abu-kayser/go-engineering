package main

import "fmt"

func summarizeServiceHealthLoop() (string, int) {
	topic := "Service Health Loop"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeServiceHealthLoop()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
