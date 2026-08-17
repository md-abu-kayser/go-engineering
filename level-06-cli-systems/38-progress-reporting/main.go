package main

import "fmt"

func summarizeProgressReporting() (string, int) {
	topic := "Progress Reporting"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeProgressReporting()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
