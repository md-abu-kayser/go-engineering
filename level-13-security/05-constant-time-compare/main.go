package main

import "fmt"

func summarizeConstantTimeCompare() (string, int) {
	topic := "Constant Time Compare"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeConstantTimeCompare()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
