package main

import "fmt"

func summarizeMultipartUpload() (string, int) {
	topic := "Multipart Upload"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMultipartUpload()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
