package main

import "fmt"

func summarizeSoftDelete() (string, int) {
	topic := "Soft Delete"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSoftDelete()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
