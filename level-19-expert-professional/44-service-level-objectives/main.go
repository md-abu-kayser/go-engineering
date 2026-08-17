package main

import "fmt"

func summarizeServiceLevelObjectives() (string, int) {
	topic := "Service Level Objectives"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeServiceLevelObjectives()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
