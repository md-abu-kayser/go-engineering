package main

import "fmt"

func summarizeGoBuild() (string, int) {
	topic := "Go Build"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGoBuild()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
