package main

import "fmt"

func summarizeRequestIdMiddleware() (string, int) {
	topic := "Request Id Middleware"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRequestIdMiddleware()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
