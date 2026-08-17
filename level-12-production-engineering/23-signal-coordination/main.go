package main

import "fmt"

func summarizeSignalCoordination() (string, int) {
	topic := "Signal Coordination"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSignalCoordination()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
