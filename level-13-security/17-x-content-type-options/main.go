package main

import "fmt"

func summarizeXContentTypeOptions() (string, int) {
	topic := "X Content Type Options"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeXContentTypeOptions()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
