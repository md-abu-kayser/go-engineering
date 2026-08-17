package main

import "fmt"

func summarizeNetworkBufferSizing() (string, int) {
	topic := "Network Buffer Sizing"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeNetworkBufferSizing()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
