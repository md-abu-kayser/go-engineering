package main

import "fmt"

func summarizeCacheKeyDesign() (string, int) {
	topic := "Cache Key Design"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCacheKeyDesign()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
