package main

import "fmt"

func summarizePodDraining() (string, int) {
	topic := "Pod Draining"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePodDraining()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
