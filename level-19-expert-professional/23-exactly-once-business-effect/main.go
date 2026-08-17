package main

import "fmt"

func summarizeExactlyOnceBusinessEffect() (string, int) {
	topic := "Exactly Once Business Effect"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeExactlyOnceBusinessEffect()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
