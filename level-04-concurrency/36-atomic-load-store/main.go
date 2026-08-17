package main

import "fmt"

func summarizeAtomicLoadStore() (string, int) {
	topic := "Atomic Load Store"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAtomicLoadStore()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
