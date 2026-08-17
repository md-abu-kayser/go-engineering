package main

import "fmt"

func summarizeNetUrl() (string, int) {
	topic := "Net Url"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeNetUrl()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
