package main

import "fmt"

func summarizeInterfaceCost() (string, int) {
	topic := "Interface Cost"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeInterfaceCost()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
