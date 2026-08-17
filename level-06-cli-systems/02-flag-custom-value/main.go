package main

import "fmt"

func summarizeFlagCustomValue() (string, int) {
	topic := "Flag Custom Value"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFlagCustomValue()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
