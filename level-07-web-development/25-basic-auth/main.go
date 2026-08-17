package main

import "fmt"

func summarizeBasicAuth() (string, int) {
	topic := "Basic Auth"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBasicAuth()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
