package main

import "fmt"

func summarizeAnonymousFunctions() (string, int) {
	topic := "Anonymous Functions"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAnonymousFunctions()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
