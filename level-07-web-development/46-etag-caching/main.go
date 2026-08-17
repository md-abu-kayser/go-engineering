package main

import "fmt"

func summarizeEtagCaching() (string, int) {
	topic := "Etag Caching"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeEtagCaching()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
