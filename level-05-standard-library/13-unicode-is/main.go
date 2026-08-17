package main

import "fmt"

func summarizeUnicodeIs() (string, int) {
	topic := "Unicode Is"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeUnicodeIs()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
