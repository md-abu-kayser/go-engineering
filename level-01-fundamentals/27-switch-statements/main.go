package main

import "fmt"

func summarizeSwitchStatements() (string, int) {
	topic := "Switch Statements"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSwitchStatements()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
