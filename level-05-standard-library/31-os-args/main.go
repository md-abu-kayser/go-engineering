package main

import "fmt"

func summarizeOsArgs() (string, int) {
	topic := "Os Args"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeOsArgs()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
