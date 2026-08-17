package main

import "fmt"

func summarizeIncidentMitigation() (string, int) {
	topic := "Incident Mitigation"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeIncidentMitigation()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
