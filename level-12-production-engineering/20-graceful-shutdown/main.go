package main

import "fmt"

func summarizeGracefulShutdown() (string, int) {
	topic := "Graceful Shutdown"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGracefulShutdown()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
