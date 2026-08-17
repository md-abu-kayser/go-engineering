package main

import "fmt"

func summarizeIota() (string, int) {
	topic := "Iota"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeIota()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
