package main

import "fmt"

func summarizeOsCreate() (string, int) {
	topic := "Os Create"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeOsCreate()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
