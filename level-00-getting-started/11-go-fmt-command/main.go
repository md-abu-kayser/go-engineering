package main

import "fmt"

func summarizeGoFmtCommand() (string, int) {
	topic := "Go Fmt Command"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGoFmtCommand()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
