package main

import "fmt"

func summarizeSafeRollout() (string, int) {
	topic := "Safe Rollout"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSafeRollout()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
