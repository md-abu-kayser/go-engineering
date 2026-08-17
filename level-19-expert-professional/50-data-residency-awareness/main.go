package main

import "fmt"

func summarizeDataResidencyAwareness() (string, int) {
	topic := "Data Residency Awareness"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDataResidencyAwareness()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
