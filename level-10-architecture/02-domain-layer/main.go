package main

import "fmt"

func summarizeDomainLayer() (string, int) {
	topic := "Domain Layer"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDomainLayer()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
