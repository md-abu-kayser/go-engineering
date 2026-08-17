package main

import "fmt"

func summarizeMaps() (string, int) {
	topic := "Maps"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMaps()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
