package main

import "fmt"

func summarizeSignalShutdown() (string, int) {
	topic := "Signal Shutdown"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSignalShutdown()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
