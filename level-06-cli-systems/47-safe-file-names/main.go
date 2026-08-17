package main

import "fmt"

func summarizeSafeFileNames() (string, int) {
	topic := "Safe File Names"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSafeFileNames()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
