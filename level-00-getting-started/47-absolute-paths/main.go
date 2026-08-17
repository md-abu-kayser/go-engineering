package main

import "fmt"

func summarizeAbsolutePaths() (string, int) {
	topic := "Absolute Paths"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAbsolutePaths()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
