package main

import "fmt"

func summarizeGossipAwareness() (string, int) {
	topic := "Gossip Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGossipAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
