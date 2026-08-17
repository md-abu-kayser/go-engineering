package main

import "fmt"

func summarizeUnsafeOffsetof() (string, int) {
	topic := "Unsafe Offsetof"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeUnsafeOffsetof()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
