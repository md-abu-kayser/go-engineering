package main

import "fmt"

func summarizeNumericConversion() (string, int) {
	topic := "Numeric Conversion"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeNumericConversion()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
