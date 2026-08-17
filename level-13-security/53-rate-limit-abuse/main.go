package main

import "fmt"

func summarizeRateLimitAbuse() (string, int) {
	topic := "Rate Limit Abuse"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRateLimitAbuse()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
