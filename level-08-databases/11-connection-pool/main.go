package main

import "fmt"

func summarizeConnectionPool() (string, int) {
	topic := "Connection Pool"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeConnectionPool()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
