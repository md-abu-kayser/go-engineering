package main

import "fmt"

func summarizeCmpPackage() (string, int) {
	topic := "Cmp Package"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCmpPackage()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
