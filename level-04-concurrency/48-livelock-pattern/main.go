package main

import "fmt"

func summarizeLivelockPattern() (string, int) {
	topic := "Livelock Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeLivelockPattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
