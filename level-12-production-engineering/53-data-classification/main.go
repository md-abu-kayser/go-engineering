package main

import "fmt"

func summarizeDataClassification() (string, int) {
	topic := "Data Classification"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDataClassification()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
