package main

import "fmt"

func summarizeIoReadall() (string, int) {
	topic := "Io Readall"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeIoReadall()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
