package main

import "fmt"

func summarizeBytesReader() (string, int) {
	topic := "Bytes Reader"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBytesReader()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
