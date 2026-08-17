package main

import "fmt"

func summarizeOperationalExcellence() (string, int) {
	topic := "Operational Excellence"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeOperationalExcellence()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
