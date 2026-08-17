package main

import "fmt"

func summarizeRequestMethod() (string, int) {
	topic := "Request Method"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRequestMethod()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
