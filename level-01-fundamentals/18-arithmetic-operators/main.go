package main

import "fmt"

func summarizeArithmeticOperators() (string, int) {
	topic := "Arithmetic Operators"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeArithmeticOperators()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
