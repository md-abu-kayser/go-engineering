package main

import "fmt"

func summarizeGcPercent() (string, int) {
	topic := "Gc Percent"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGcPercent()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
