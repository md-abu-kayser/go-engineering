package main

import "fmt"

func summarizeCommandArguments() (string, int) {
	topic := "Command Arguments"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCommandArguments()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
