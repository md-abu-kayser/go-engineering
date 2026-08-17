package main

import "fmt"

func summarizeBuildCacheInternals() (string, int) {
	topic := "Build Cache Internals"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBuildCacheInternals()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
