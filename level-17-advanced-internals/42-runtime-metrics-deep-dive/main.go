package main

import "fmt"

func summarizeRuntimeMetricsDeepDive() (string, int) {
	topic := "Runtime Metrics Deep Dive"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRuntimeMetricsDeepDive()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
