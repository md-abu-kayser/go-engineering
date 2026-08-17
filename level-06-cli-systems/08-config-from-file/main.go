package main

import "fmt"

func summarizeConfigFromFile() (string, int) {
	topic := "Config From File"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeConfigFromFile()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
