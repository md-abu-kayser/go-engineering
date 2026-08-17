package main

import "fmt"

func summarizeLockFiles() (string, int) {
	topic := "Lock Files"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeLockFiles()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
