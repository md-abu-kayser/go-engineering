package main

import "fmt"

func summarizeAtLeastOnceDelivery() (string, int) {
	topic := "At Least Once Delivery"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAtLeastOnceDelivery()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
