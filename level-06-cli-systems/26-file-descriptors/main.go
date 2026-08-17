package main

import "fmt"

func summarizeFileDescriptors() (string, int) {
	topic := "File Descriptors"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFileDescriptors()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
