package main

import "fmt"

func summarizeTimeTimer() (string, int) {
	topic := "Time Timer"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeTimeTimer()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
