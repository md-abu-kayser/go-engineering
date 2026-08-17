package main

import "fmt"

func summarizeHedgingPattern() (string, int) {
	topic := "Hedging Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeHedgingPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
