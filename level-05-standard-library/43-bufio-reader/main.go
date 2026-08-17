package main

import "fmt"

func summarizeBufioReader() (string, int) {
	topic := "Bufio Reader"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeBufioReader()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
