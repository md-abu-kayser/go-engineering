package main

import "fmt"

func summarizeRaceFreeDesign() (string, int) {
	topic := "Race Free Design"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRaceFreeDesign()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
