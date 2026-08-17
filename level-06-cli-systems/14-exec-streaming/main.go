package main

import "fmt"

func summarizeExecStreaming() (string, int) {
	topic := "Exec Streaming"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeExecStreaming()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
