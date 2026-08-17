package main

import "fmt"

func summarizeSourceLayout() (string, int) {
	topic := "Source Layout"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSourceLayout()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
