package main

import "fmt"

func summarizeStderrLogging() (string, int) {
	topic := "Stderr Logging"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStderrLogging()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
