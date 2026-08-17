package main

import "fmt"

func summarizeHappensBefore() (string, int) {
	topic := "Happens Before"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeHappensBefore()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
