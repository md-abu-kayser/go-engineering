package main

import "fmt"

func summarizeLazyInitialization() (string, int) {
	topic := "Lazy Initialization"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeLazyInitialization()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
