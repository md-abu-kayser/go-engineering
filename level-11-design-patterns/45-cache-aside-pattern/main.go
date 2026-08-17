package main

import "fmt"

func summarizeCacheAsidePattern() (string, int) {
	topic := "Cache Aside Pattern"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCacheAsidePattern()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
