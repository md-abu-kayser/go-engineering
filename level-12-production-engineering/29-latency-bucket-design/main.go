package main

import "fmt"

func summarizeLatencyBucketDesign() (string, int) {
	topic := "Latency Bucket Design"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeLatencyBucketDesign()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
