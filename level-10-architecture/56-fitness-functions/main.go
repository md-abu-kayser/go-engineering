package main

import "fmt"

func summarizeFitnessFunctions() (string, int) {
	topic := "Fitness Functions"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFitnessFunctions()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
