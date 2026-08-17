package main

import "fmt"

func summarizeDistributedTimeouts() (string, int) {
	topic := "Distributed Timeouts"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDistributedTimeouts()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
