package main

import "fmt"

func summarizeIncidentContext() (string, int) {
	topic := "Incident Context"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeIncidentContext()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
