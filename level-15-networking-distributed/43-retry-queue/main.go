package main

import "fmt"

func summarizeRetryQueue() (string, int) {
	topic := "Retry Queue"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRetryQueue()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
