package main

import "fmt"

func summarizeSecurityLogging() (string, int) {
	topic := "Security Logging"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSecurityLogging()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
