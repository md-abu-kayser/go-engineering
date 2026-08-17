package main

import "fmt"

func summarizeHotPathDesign() (string, int) {
	topic := "Hot Path Design"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeHotPathDesign()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
