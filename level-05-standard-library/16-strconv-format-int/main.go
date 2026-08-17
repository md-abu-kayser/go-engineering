package main

import "fmt"

func summarizeStrconvFormatInt() (string, int) {
	topic := "Strconv Format Int"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeStrconvFormatInt()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
