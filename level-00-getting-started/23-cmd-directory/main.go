package main

import "fmt"

func summarizeCmdDirectory() (string, int) {
	topic := "Cmd Directory"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCmdDirectory()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
