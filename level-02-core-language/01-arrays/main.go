package main

import "fmt"

func summarizeArrays() (string, int) {
	topic := "Arrays"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeArrays()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
