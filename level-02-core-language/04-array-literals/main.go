package main

import "fmt"

func summarizeArrayLiterals() (string, int) {
	topic := "Array Literals"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeArrayLiterals()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
