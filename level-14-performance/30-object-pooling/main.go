package main

import "fmt"

func summarizeObjectPooling() (string, int) {
	topic := "Object Pooling"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeObjectPooling()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
