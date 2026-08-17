package main

import "fmt"

func summarizeStringsCut() (string, int) {
	topic := "Strings Cut"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStringsCut()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
