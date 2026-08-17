package main

import "fmt"

func summarizeAtomicFileWrite() (string, int) {
	topic := "Atomic File Write"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAtomicFileWrite()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
