package main

import "fmt"

func summarizeTHelper() (string, int) {
	topic := "T Helper"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeTHelper()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
