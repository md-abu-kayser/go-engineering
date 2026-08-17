package main

import "fmt"

func summarizeInitFunctions() (string, int) {
	topic := "Init Functions"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeInitFunctions()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
