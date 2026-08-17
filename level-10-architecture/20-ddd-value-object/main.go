package main

import "fmt"

func summarizeDddValueObject() (string, int) {
	topic := "Ddd Value Object"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDddValueObject()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
