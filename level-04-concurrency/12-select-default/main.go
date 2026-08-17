package main

import "fmt"

func summarizeSelectDefault() (string, int) {
	topic := "Select Default"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSelectDefault()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
