package main

import "fmt"

func summarizePackageCycles() (string, int) {
	topic := "Package Cycles"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePackageCycles()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
