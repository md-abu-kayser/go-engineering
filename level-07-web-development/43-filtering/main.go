package main

import "fmt"

func summarizeFiltering() (string, int) {
	topic := "Filtering"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFiltering()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
