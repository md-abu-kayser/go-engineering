package main

import "fmt"

func summarizeStagingParity() (string, int) {
	topic := "Staging Parity"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStagingParity()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
