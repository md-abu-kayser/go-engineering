package main

import "fmt"

func summarizeFileWalk() (string, int) {
	topic := "File Walk"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeFileWalk()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
