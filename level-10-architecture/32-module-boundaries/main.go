package main

import "fmt"

func summarizeModuleBoundaries() (string, int) {
	topic := "Module Boundaries"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeModuleBoundaries()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
