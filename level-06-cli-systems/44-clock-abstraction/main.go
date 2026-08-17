package main

import "fmt"

func summarizeClockAbstraction() (string, int) {
	topic := "Clock Abstraction"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeClockAbstraction()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
