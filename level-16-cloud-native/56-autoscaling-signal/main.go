package main

import "fmt"

func summarizeAutoscalingSignal() (string, int) {
	topic := "Autoscaling Signal"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAutoscalingSignal()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
