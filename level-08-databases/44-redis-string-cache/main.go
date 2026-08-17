package main

import "fmt"

func summarizeRedisStringCache() (string, int) {
	topic := "Redis String Cache"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRedisStringCache()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
