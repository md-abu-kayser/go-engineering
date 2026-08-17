package main

import "fmt"

func summarizeDistributedLockService() (string, int) {
	topic := "Distributed Lock Service"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDistributedLockService()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
