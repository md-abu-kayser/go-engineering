package main

import "fmt"

func summarizeFanOut() (string, int) {
	topic := "Fan Out"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFanOut()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
