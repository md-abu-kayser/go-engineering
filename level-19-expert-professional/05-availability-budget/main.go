package main

import "fmt"

func summarizeAvailabilityBudget() (string, int) {
	topic := "Availability Budget"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAvailabilityBudget()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
