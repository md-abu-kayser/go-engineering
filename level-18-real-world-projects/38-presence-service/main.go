package main

import "fmt"

func summarizePresenceService() (string, int) {
	topic := "Presence Service"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePresenceService()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
