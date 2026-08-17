package main

import "fmt"

func summarizeClockPort() (string, int) {
	topic := "Clock Port"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeClockPort()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
