package main

import "fmt"

func summarizeGoEnvCommand() (string, int) {
	topic := "Go Env Command"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGoEnvCommand()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
