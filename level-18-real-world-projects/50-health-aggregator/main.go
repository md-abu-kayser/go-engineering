package main

import "fmt"

func summarizeHealthAggregator() (string, int) {
	topic := "Health Aggregator"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeHealthAggregator()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
