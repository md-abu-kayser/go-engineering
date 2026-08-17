package main

import "fmt"

func summarizeCapacityPlanning() (string, int) {
	topic := "Capacity Planning"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCapacityPlanning()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
