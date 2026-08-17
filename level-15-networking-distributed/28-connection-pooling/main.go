package main

import "fmt"

func summarizeConnectionPooling() (string, int) {
	topic := "Connection Pooling"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeConnectionPooling()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
