package main

import "fmt"

func summarizeContentType() (string, int) {
	topic := "Content Type"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeContentType()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
