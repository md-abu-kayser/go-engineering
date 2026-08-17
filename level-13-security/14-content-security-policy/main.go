package main

import "fmt"

func summarizeContentSecurityPolicy() (string, int) {
	topic := "Content Security Policy"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeContentSecurityPolicy()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
