package main

import "fmt"

func summarizeDatabaseObservability() (string, int) {
	topic := "Database Observability"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDatabaseObservability()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
