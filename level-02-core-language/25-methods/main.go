package main

import "fmt"

func summarizeMethods() (string, int) {
	topic := "Methods"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMethods()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
