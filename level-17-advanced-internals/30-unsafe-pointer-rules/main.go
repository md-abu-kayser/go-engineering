package main

import "fmt"

func summarizeUnsafePointerRules() (string, int) {
	topic := "Unsafe Pointer Rules"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeUnsafePointerRules()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
