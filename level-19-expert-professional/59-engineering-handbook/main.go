package main

import "fmt"

func summarizeEngineeringHandbook() (string, int) {
	topic := "Engineering Handbook"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeEngineeringHandbook()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
