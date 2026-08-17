package main

import "fmt"

func summarizeConstructorFunctions() (string, int) {
	topic := "Constructor Functions"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeConstructorFunctions()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
