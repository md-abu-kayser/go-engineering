package main

import "fmt"

func summarizeAllocProfile() (string, int) {
	topic := "Alloc Profile"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAllocProfile()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
