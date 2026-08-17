package main

import "fmt"

func summarizeCiBuildCache() (string, int) {
	topic := "Ci Build Cache"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCiBuildCache()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
