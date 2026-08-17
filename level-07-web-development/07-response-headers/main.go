package main

import "fmt"

func summarizeResponseHeaders() (string, int) {
	topic := "Response Headers"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeResponseHeaders()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
