package main

import "fmt"

func summarizeStringerPattern() (string, int) {
	topic := "Stringer Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStringerPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
