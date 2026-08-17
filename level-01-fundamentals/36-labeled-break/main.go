package main

import "fmt"

func summarizeLabeledBreak() (string, int) {
	topic := "Labeled Break"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeLabeledBreak()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
