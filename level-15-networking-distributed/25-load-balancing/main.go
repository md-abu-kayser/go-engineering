package main

import "fmt"

func summarizeLoadBalancing() (string, int) {
	topic := "Load Balancing"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeLoadBalancing()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
