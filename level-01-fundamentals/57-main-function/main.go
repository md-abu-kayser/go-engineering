package main

import "fmt"

func summarizeMainFunction() (string, int) {
	topic := "Main Function"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMainFunction()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
