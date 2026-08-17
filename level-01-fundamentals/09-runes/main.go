package main

import "fmt"

func summarizeRunes() (string, int) {
	topic := "Runes"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRunes()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
