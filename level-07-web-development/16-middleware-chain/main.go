package main

import "fmt"

func summarizeMiddlewareChain() (string, int) {
	topic := "Middleware Chain"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMiddlewareChain()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
