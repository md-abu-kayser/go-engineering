package main

import "fmt"

func summarizeRtoRpo() (string, int) {
	topic := "Rto Rpo"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRtoRpo()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
