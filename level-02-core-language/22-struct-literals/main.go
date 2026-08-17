package main

import "fmt"

func summarizeStructLiterals() (string, int) {
	topic := "Struct Literals"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStructLiterals()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
