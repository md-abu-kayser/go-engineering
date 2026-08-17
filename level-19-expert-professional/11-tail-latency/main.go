package main

import "fmt"

func summarizeTailLatency() (string, int) {
	topic := "Tail Latency"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeTailLatency()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
