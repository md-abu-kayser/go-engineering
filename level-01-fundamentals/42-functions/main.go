package main

import "fmt"

func summarizeFunctions() (string, int) {
	topic := "Functions"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFunctions()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
