package main

import "fmt"

func summarizeNewFunction() (string, int) {
	topic := "New Function"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeNewFunction()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
