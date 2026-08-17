package main

import "fmt"

func summarizeServiceTemplate() (string, int) {
	topic := "Service Template"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeServiceTemplate()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
