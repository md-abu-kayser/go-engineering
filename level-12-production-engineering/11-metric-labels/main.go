package main

import "fmt"

func summarizeMetricLabels() (string, int) {
	topic := "Metric Labels"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMetricLabels()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
