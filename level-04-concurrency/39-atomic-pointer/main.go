package main

import "fmt"

func summarizeAtomicPointer() (string, int) {
	topic := "Atomic Pointer"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAtomicPointer()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
