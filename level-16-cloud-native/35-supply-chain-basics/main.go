package main

import "fmt"

func summarizeSupplyChainBasics() (string, int) {
	topic := "Supply Chain Basics"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSupplyChainBasics()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
