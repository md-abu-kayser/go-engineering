package main

import "fmt"

func summarizeStringsReplace() (string, int) {
	topic := "Strings Replace"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStringsReplace()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
