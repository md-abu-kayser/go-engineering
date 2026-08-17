package main

import "fmt"

func summarizeModuleCacheOverview() (string, int) {
	topic := "Module Cache Overview"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeModuleCacheOverview()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
