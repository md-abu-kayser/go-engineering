package main

import "fmt"

func summarizeRestResourceDesign() (string, int) {
	topic := "Rest Resource Design"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRestResourceDesign()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
