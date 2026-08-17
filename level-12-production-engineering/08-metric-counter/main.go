package main

import "fmt"

func summarizeMetricCounter() (string, int) {
	topic := "Metric Counter"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMetricCounter()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
