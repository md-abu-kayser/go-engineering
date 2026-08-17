package main

import "fmt"

func summarizeMathRandBasics() (string, int) {
	topic := "Math Rand Basics"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMathRandBasics()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
