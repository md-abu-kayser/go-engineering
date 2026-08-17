package main

import "fmt"

func summarizeApplicationBootstrap() (string, int) {
	topic := "Application Bootstrap"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeApplicationBootstrap()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
