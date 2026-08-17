package main

import "fmt"

func summarizeZeroValues() (string, int) {
	topic := "Zero Values"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeZeroValues()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
