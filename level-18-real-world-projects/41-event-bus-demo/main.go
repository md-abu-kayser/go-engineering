package main

import "fmt"

func summarizeEventBusDemo() (string, int) {
	topic := "Event Bus Demo"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeEventBusDemo()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
