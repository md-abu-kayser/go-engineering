package main

import "fmt"

func summarizeCliExitCodes() (string, int) {
	topic := "Cli Exit Codes"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCliExitCodes()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
