package main

import "fmt"

func summarizeEventReplay() (string, int) {
	topic := "Event Replay"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeEventReplay()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
