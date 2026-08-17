package main

import "fmt"

func summarizeReflectionDynamicCall() (string, int) {
	topic := "Reflection Dynamic Call"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeReflectionDynamicCall()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
