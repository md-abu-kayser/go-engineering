package main

import "fmt"

func summarizeDtoBoundary() (string, int) {
	topic := "Dto Boundary"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDtoBoundary()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
