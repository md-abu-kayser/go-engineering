package main

import "fmt"

func summarizeGcSweepPhase() (string, int) {
	topic := "Gc Sweep Phase"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGcSweepPhase()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
