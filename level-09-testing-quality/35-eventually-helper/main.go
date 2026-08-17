package main

import "fmt"

func summarizeEventuallyHelper() (string, int) {
	topic := "Eventually Helper"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeEventuallyHelper()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
