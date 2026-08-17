package main

import "fmt"

func summarizeNamedResults() (string, int) {
	topic := "Named Results"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeNamedResults()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
