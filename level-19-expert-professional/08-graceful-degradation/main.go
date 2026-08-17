package main

import "fmt"

func summarizeGracefulDegradation() (string, int) {
	topic := "Graceful Degradation"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGracefulDegradation()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
