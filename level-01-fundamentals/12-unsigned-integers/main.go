package main

import "fmt"

func summarizeUnsignedIntegers() (string, int) {
	topic := "Unsigned Integers"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeUnsignedIntegers()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
