package main

import "fmt"

func summarizeShutdownOrder() (string, int) {
	topic := "Shutdown Order"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeShutdownOrder()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
