package main

import "fmt"

func summarizePathParameters() (string, int) {
	topic := "Path Parameters"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePathParameters()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
