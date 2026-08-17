package main

import "fmt"

func summarizeConsistencyTradeoffs() (string, int) {
	topic := "Consistency Tradeoffs"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeConsistencyTradeoffs()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
