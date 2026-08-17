package main

import "fmt"

func summarizeProxyPattern() (string, int) {
	topic := "Proxy Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeProxyPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
