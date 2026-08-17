package main

import "fmt"

func summarizeBackpressureDesign() (string, int) {
	topic := "Backpressure Design"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBackpressureDesign()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
