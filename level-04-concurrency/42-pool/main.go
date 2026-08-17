package main

import "fmt"

func summarizePool() (string, int) {
	topic := "Pool"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePool()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
