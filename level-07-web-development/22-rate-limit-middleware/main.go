package main

import "fmt"

func summarizeRateLimitMiddleware() (string, int) {
	topic := "Rate Limit Middleware"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRateLimitMiddleware()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
