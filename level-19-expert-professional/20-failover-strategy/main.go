package main

import "fmt"

func summarizeFailoverStrategy() (string, int) {
	topic := "Failover Strategy"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFailoverStrategy()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
