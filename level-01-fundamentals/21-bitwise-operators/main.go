package main

import "fmt"

func summarizeBitwiseOperators() (string, int) {
	topic := "Bitwise Operators"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBitwiseOperators()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
