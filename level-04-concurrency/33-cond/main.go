package main

import "fmt"

func summarizeCond() (string, int) {
	topic := "Cond"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCond()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
